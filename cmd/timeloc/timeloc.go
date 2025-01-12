// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"github.com/kitimark/timeloc"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(timeloc.Analyzer)
}
