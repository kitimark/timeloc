// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package gclplugin

import (
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/plugin-module-register/register"
	"github.com/kitimark/timeloc"
)

func init() {
	register.Plugin("timeloc", NewPlugin)
}

type timeLocPlugin struct{}

// NewPlugin creates a new instance of the timeloc plugin
func NewPlugin(conf any) (register.LinterPlugin, error) {
	return &timeLocPlugin{}, nil
}

// BuildAnalyzers returns the analyzers this plugin provides
func (p *timeLocPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{timeloc.Analyzer}, nil
}

// GetLoadMode returns how much information should be loaded for analysis
func (p *timeLocPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
