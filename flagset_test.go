//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/flagset_test.go
//

package vflag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagSetSetInvalidValue(t *testing.T) {
	fset := NewFlagSet("test", ContinueOnError)
	var value int64
	fset.Int64Var(&value, 0, "count")
	err := fset.Parse([]string{"--count", "not-a-valid-number"})
	assert.Error(t, err)
}

func TestFlagSetMaybeHandleError(t *testing.T) {
	t.Run("PanicOnError", func(t *testing.T) {
		fset := NewFlagSet("test", PanicOnError)
		assert.Panics(t, func() {
			fset.Parse([]string{"--unknown"})
		})
	})

	t.Run("ContinueOnError", func(t *testing.T) {
		fset := NewFlagSet("test", ContinueOnError)
		fset.AutoHelp('h', "help", "Show this help message and exit.")
		err := fset.Parse([]string{"--help"})
		assert.ErrorIs(t, err, ErrHelp)
	})
}

func TestFlagSetAddInvalidFlagPanics(t *testing.T) {
	fset := NewFlagSet("test", ExitOnError)
	assert.Panics(t, func() {
		fset.AddFlag(&Flag{}) // add completely empty flag
	})
}
