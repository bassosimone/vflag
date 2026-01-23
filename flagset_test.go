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

func TestFlagSetParsePanicsOnDuplicateName(t *testing.T) {
	fset := NewFlagSet("test", ContinueOnError)

	// Add a short flag named 'v'
	var verbose bool
	fset.ShortFlags = append(fset.ShortFlags, NewShortFlagBool(
		NewValueBool(&verbose), 'v', "Enable verbose output.",
	))

	// Add a long flag also named 'v' - this should panic
	var version bool
	fset.LongFlags = append(fset.LongFlags, NewLongFlagBool(
		NewValueBool(&version), "v", "Show version.",
	))

	assert.Panics(t, func() {
		fset.Parse([]string{})
	})
}

func TestFlagSetParseShortFlagPanicsOnEmptyPrefix(t *testing.T) {
	fset := NewFlagSet("test", ContinueOnError)
	var verbose bool
	sf := NewShortFlagBool(NewValueBool(&verbose), 'v', "Enable verbose output.")
	sf.Prefix = "" // break it
	fset.ShortFlags = append(fset.ShortFlags, sf)

	assert.Panics(t, func() {
		fset.Parse([]string{})
	})
}

func TestFlagSetParseLongFlagPanicsOnEmptyPrefix(t *testing.T) {
	fset := NewFlagSet("test", ContinueOnError)
	var verbose bool
	lf := NewLongFlagBool(NewValueBool(&verbose), "verbose", "Enable verbose output.")
	lf.Prefix = "" // break it
	fset.LongFlags = append(fset.LongFlags, lf)

	assert.Panics(t, func() {
		fset.Parse([]string{})
	})
}
