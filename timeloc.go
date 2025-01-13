// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package timeloc

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	astinspector "golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "timeloc",
	Doc: `Detects time.Time methods used without setting location and prevents time.Local usage. 
Many time.Time methods depend on the location (timezone) setting. If not explicitly set, these methods will use the system default timezone, 
which can lead to inconsistencies across different environments. This analyzer enforces explicit location setting
through time.In(), time.ParseInLocation(), or functions that internally set location. It also prevents usage of time.Local
which relies on system timezone.`,
	Run: run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var targetMethods = map[string]bool{
	"AppendFormat": true,
	"Clock":        true,
	"Date":         true,
	"Day":          true,
	"Format":       true,
	"Hour":         true,
	"ISOWeek":      true,
	"Minute":       true,
	"Month":        true,
	"Second":       true,
	"Weekday":      true,
	"Year":         true,
	"YearDay":      true,
}

type state struct {
	pass          *analysis.Pass
	timeVars      map[*types.Var]bool  // true if location is set
	unsafeFuncs   map[*types.Func]bool // true if function uses time methods on parameters
	analyzedFuncs map[*types.Func]bool // track which functions we've analyzed
}

func run(pass *analysis.Pass) (interface{}, error) {
	s := &state{
		pass:          pass,
		timeVars:      make(map[*types.Var]bool),
		unsafeFuncs:   make(map[*types.Func]bool),
		analyzedFuncs: make(map[*types.Func]bool),
	}

	inspector := pass.ResultOf[inspect.Analyzer].(*astinspector.Inspector)

	// analyze functions
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				s.analyzeFuncDecl(fn)
			}
		}
	}

	// track assignments and check calls
	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.AssignStmt:
			s.handleAssign(node)
		case *ast.CallExpr:
			s.handleCall(node)
		}
	})

	return nil, nil
}

func (s *state) analyzeFuncDecl(fn *ast.FuncDecl) *types.Func {
	obj := s.pass.TypesInfo.ObjectOf(fn.Name).(*types.Func)
	if s.analyzedFuncs[obj] {
		return obj
	}
	s.analyzedFuncs[obj] = true

	paramNames := make(map[string]bool)
	if fn.Type.Params != nil {
		for _, field := range fn.Type.Params.List {
			if s.isTimeType(s.pass.TypesInfo.TypeOf(field.Type)) {
				for _, name := range field.Names {
					paramNames[name.Name] = true
				}
			}
		}
	}

	var usesUnsafeMethods bool
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			switch fun := call.Fun.(type) {
			case *ast.SelectorExpr:
				if id, ok := fun.X.(*ast.Ident); ok {
					if paramNames[id.Name] && targetMethods[fun.Sel.Name] {
						usesUnsafeMethods = true
						return false
					}
				}
			case *ast.Ident:
				if calledFn := s.pass.TypesInfo.ObjectOf(fun); calledFn != nil {
					if f, ok := calledFn.(*types.Func); ok {
						if s.unsafeFuncs[f] {
							for _, arg := range call.Args {
								if id, ok := arg.(*ast.Ident); ok {
									if paramNames[id.Name] {
										usesUnsafeMethods = true
										return false
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	s.unsafeFuncs[obj] = usesUnsafeMethods
	return obj
}

func (s *state) handleAssign(assign *ast.AssignStmt) {
	for i, rhs := range assign.Rhs {
		if i >= len(assign.Lhs) {
			continue
		}

		if id, ok := assign.Lhs[i].(*ast.Ident); ok {
			if obj, ok := s.pass.TypesInfo.ObjectOf(id).(*types.Var); ok {
				if call, ok := rhs.(*ast.CallExpr); ok {
					switch {
					case s.isTimeNow(call) || s.isTimeParse(call):
						s.timeVars[obj] = false
					case s.isTimeParseInLocation(call) || s.isTimeLocSetter(call):
						s.timeVars[obj] = true
					case s.isInCall(call):
						if rec, ok := s.getReceiverVar(call); ok {
							s.timeVars[rec] = true
						}
						s.timeVars[obj] = true
					}
				} else if assign.Tok == token.ASSIGN { // handle simple assignment
					if rhsId, ok := rhs.(*ast.Ident); ok {
						if rhsObj, ok := s.pass.TypesInfo.ObjectOf(rhsId).(*types.Var); ok {
							if s.isTimeType(rhsObj.Type()) {
								s.timeVars[obj] = s.timeVars[rhsObj]
							}
						}
					}
				}
			}
		}
	}
}

func (s *state) handleCall(call *ast.CallExpr) {
	// check for time.Local usage in various contexts
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			if x.Name == "time" {
				// check time.ParseInLocation with time.Local
				if sel.Sel.Name == "ParseInLocation" && len(call.Args) >= 3 {
					if s.isTimeLocal(call.Args[2]) {
						s.pass.Reportf(call.Pos(), "time.Local usage is not allowed as it relies on system timezone")
					}
				}
				// check time.Date with time.Local
				if sel.Sel.Name == "Date" && len(call.Args) >= 8 {
					if s.isTimeLocal(call.Args[7]) {
						s.pass.Reportf(call.Pos(), "time.Local usage is not allowed as it relies on system timezone")
					}
				}
			}
		}
		// check t.In(time.Local)
		if sel.Sel.Name == "In" && len(call.Args) > 0 {
			if s.isTimeLocal(call.Args[0]) {
				s.pass.Reportf(call.Pos(), "time.Local usage is not allowed as it relies on system timezone")
			}
		}
	}

	// continue with existing time location checks
	if s.isInCall(call) {
		if rec, ok := s.getReceiverVar(call); ok {
			s.timeVars[rec] = true
		}
		return
	}

	// check for time-dependent method calls
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if id, ok := sel.X.(*ast.Ident); ok {
			if obj, ok := s.pass.TypesInfo.ObjectOf(id).(*types.Var); ok {
				if s.isTimeType(obj.Type()) && targetMethods[sel.Sel.Name] {
					if hasLocation, exists := s.timeVars[obj]; exists && !hasLocation {
						s.pass.Reportf(call.Pos(), "(t time.Time).%s called on %s before setting time.Location", sel.Sel.Name, id.Name)
					}
				}
			}
		}
	}

	// check function calls with time parameters
	if fn := s.getFuncDecl(call); fn != nil && s.unsafeFuncs[fn] {
		for _, arg := range call.Args {
			if id, ok := arg.(*ast.Ident); ok {
				if obj, ok := s.pass.TypesInfo.ObjectOf(id).(*types.Var); ok {
					if s.isTimeType(obj.Type()) {
						if hasLocation, exists := s.timeVars[obj]; exists && !hasLocation {
							s.pass.Reportf(call.Pos(), "passing time.Time value without location set to function that may use location-dependent methods")
						}
					}
				}
			}
		}
	}
}

func (s *state) getReceiverVar(call *ast.CallExpr) (*types.Var, bool) {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if id, ok := sel.X.(*ast.Ident); ok {
			if obj, ok := s.pass.TypesInfo.ObjectOf(id).(*types.Var); ok {
				return obj, true
			}
		}
	}
	return nil, false
}

func (s *state) isTimeLocal(expr ast.Expr) bool {
	if sel, ok := expr.(*ast.SelectorExpr); ok {
		if id, ok := sel.X.(*ast.Ident); ok {
			return id.Name == "time" && sel.Sel.Name == "Local"
		}
	}
	return false
}

func (s *state) isTimeType(t types.Type) bool {
	if named, ok := t.(*types.Named); ok {
		pkg := named.Obj().Pkg()
		// ignore built-in interface. for instance, error type
		if pkg == nil {
			return false
		}
		return pkg.Path() == "time" && named.Obj().Name() == "Time"
	}
	return false
}

func (s *state) isTimeNow(call *ast.CallExpr) bool {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			return x.Name == "time" && sel.Sel.Name == "Now"
		}
	}
	return false
}

func (s *state) isTimeParse(call *ast.CallExpr) bool {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			return x.Name == "time" && sel.Sel.Name == "Parse"
		}
	}
	return false
}

func (s *state) isTimeParseInLocation(call *ast.CallExpr) bool {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			return x.Name == "time" && sel.Sel.Name == "ParseInLocation"
		}
	}
	return false
}

func (s *state) isInCall(call *ast.CallExpr) bool {
	if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
		if id, ok := sel.X.(*ast.Ident); ok {
			if obj, ok := s.pass.TypesInfo.ObjectOf(id).(*types.Var); ok {
				return s.isTimeType(obj.Type()) && sel.Sel.Name == "In"
			}
		}
	}
	return false
}

func (s *state) isTimeLocSetter(call *ast.CallExpr) bool {
	if fn := s.getFuncDecl(call); fn != nil {
		if funcDecl := s.findFuncDecl(fn); funcDecl != nil {
			var setsLocation bool
			ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
				if call, ok := n.(*ast.CallExpr); ok {
					if s.isInCall(call) || s.isTimeParseInLocation(call) {
						setsLocation = true
						return false
					}
				}
				return true
			})
			return setsLocation
		}
	}
	return false
}

func (s *state) findFuncDecl(fn *types.Func) *ast.FuncDecl {
	var found *ast.FuncDecl
	for _, file := range s.pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if fd, ok := n.(*ast.FuncDecl); ok {
				if s.pass.TypesInfo.ObjectOf(fd.Name) == fn {
					found = fd
					return false
				}
			}
			return true
		})
		if found != nil {
			break
		}
	}
	return found
}

func (s *state) getFuncDecl(call *ast.CallExpr) *types.Func {
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		if obj := s.pass.TypesInfo.ObjectOf(fun); obj != nil {
			if fn, ok := obj.(*types.Func); ok {
				return fn
			}
		}
	case *ast.SelectorExpr:
		if obj := s.pass.TypesInfo.ObjectOf(fun.Sel); obj != nil {
			if fn, ok := obj.(*types.Func); ok {
				return fn
			}
		}
	}
	return nil
}
