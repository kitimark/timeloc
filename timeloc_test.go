// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package timeloc

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestTimeLocAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "timelocusage")
}
