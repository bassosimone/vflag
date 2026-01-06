// SPDX-License-Identifier: GPL-3.0-or-later

package vflag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsageFlagSetPositionalArgumentsUsage(t *testing.T) {
	cases := []struct {
		name       string
		minArgs    int
		maxArgs    int
		usage      string
		wantOutput string
	}{
		{
			name:       "zero-max",
			minArgs:    0,
			maxArgs:    0,
			usage:      "",
			wantOutput: "",
		},

		{
			name:       "default-single-optional",
			minArgs:    0,
			maxArgs:    1,
			usage:      "",
			wantOutput: " [arg]",
		},

		{
			name:       "default-multiple-optional",
			minArgs:    0,
			maxArgs:    2,
			usage:      "",
			wantOutput: " [arg ...]",
		},

		{
			name:       "default-multiple-required",
			minArgs:    1,
			maxArgs:    2,
			usage:      "",
			wantOutput: " arg [arg ...]",
		},

		{
			name:       "default-required-single",
			minArgs:    1,
			maxArgs:    1,
			usage:      "",
			wantOutput: " arg",
		},

		{
			name:       "custom-single",
			minArgs:    0,
			maxArgs:    1,
			usage:      "URL",
			wantOutput: " URL",
		},

		{
			name:       "custom-multiple",
			minArgs:    1,
			maxArgs:    3,
			usage:      "FILE...",
			wantOutput: " FILE...",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFlagSet("prog", ContinueOnError)
			fs.MinPositionalArgs = tc.minArgs
			fs.MaxPositionalArgs = tc.maxArgs
			usage := NewDefaultUsagePrinter()
			usage.PositionalArgumentsUsage = tc.usage
			fs.UsagePrinter = usage
			output := usage.positionalArgumentsUsage(fs)
			require.Equal(t, tc.wantOutput, output)
		})
	}
}
