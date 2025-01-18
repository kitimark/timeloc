// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package gclplugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/kitimark/timeloc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPlugin(t *testing.T) {
	t.Run("creates plugin without configuration", func(t *testing.T) {
		plugin, err := NewPlugin(nil)
		require.NoError(t, err)
		require.NotNil(t, plugin)

		require.Equal(t, register.LoadModeTypesInfo, plugin.GetLoadMode())
		analyzers, err := plugin.BuildAnalyzers()
		require.NoError(t, err)
		require.Len(t, analyzers, 1)
		require.Equal(t, timeloc.Analyzer, analyzers[0])
	})

	t.Run("creates plugin with empty configuration", func(t *testing.T) {
		plugin, err := NewPlugin(map[string]interface{}{})
		require.NoError(t, err)
		require.NotNil(t, plugin)

		require.Equal(t, register.LoadModeTypesInfo, plugin.GetLoadMode())
		analyzers, err := plugin.BuildAnalyzers()
		require.NoError(t, err)
		require.Len(t, analyzers, 1)
		require.Equal(t, timeloc.Analyzer, analyzers[0])
	})
}
