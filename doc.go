//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/doc.go
//

/*
Package vflag provides facilities for command-line flag parsing with support for both
short and long options with an API similar to the stdlib [flag] package.

The [*FlagSet] type represents a set of command-line flags and provides methods
to define boolean and string flags with various combinations of short (-f) and
long (--flag) option names. The package supports GNU-style flag parsing with
customizable option-specific prefixes and option-arguments separator.

The [NewFlagSet] function creates a new flag set with configurable error
handling behavior. Use [*FlagSet.BoolVar], [*FlagSet.StringVar], and their variants
to define flags, then call [*FlagSet.Parse] to parse command-line arguments.

The package provides comprehensive flag definition methods including [*FlagSet.BoolVar],
[*FlagSet.Int64Var], [*FlagSet.StringVar], etc. that accept existing pointers to variables
holding initial-default values. The [*FlagSet.AutoHelp] method helps to automatically
generate and handle help flags (typically `-h` and `--help`).
*/
package vflag
