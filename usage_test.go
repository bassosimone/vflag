// SPDX-License-Identifier: GPL-3.0-or-later

package vflag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsageFlagSetFlags(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value bool
	fs.AddFlag(NewFlagBool(NewValueBool(&value), 0, "", "Ignored flag."))

	flags := UsageFlagSet{Set: fs}.Flags()
	require.Empty(t, flags)
}

func TestUsageFlagSetHelpFlag(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	fs.AddFlag(&Flag{
		Value: ValueAutoHelp{},
	})

	output := UsageFlagSet{Set: fs}.HelpFlag()
	require.Empty(t, output)
}

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
			fs.PositionalArgumentsUsage = tc.usage

			output := UsageFlagSet{Set: fs}.PositionalArgumentsUsage()
			require.Equal(t, tc.wantOutput, output)
		})
	}
}
