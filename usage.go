//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/usage.go
//

package vflag

import (
	"fmt"
	"io"
	"strings"

	"github.com/bassosimone/must"
	"github.com/bassosimone/runtimex"
	"github.com/bassosimone/textwrap"
)

// PrintUsageString writes the usage string to the given [io.Writer].
//
// This function panics if writing to the [io.Writer] fails.
func (fs *FlagSet) PrintUsageString(w io.Writer) {
	fs.UsagePrinter.PrintUsageString(fs, w)
}

// PrintUsageError writes the usage error that occurred to the given [io.Writer].
//
// This function panics if writing to the [io.Writer] fails.
//
// If auto-help has been used, this function also prints a hint for the user.
func (fs *FlagSet) PrintUsageError(w io.Writer, err error) {
	fs.UsagePrinter.PrintUsageError(fs, w, err)
}

func (up *DefaultUsagePrinter) flagsName(fset *FlagSet) (output string) {
	if f := fset.Flags(); len(f) > 0 {
		output = " [flags]"
	}
	return output
}

// HelpInvocation returns the string with which to obtain help.
func (fs *FlagSet) HelpInvocation() string {
	for _, fx := range fs.flags {
		if _, ok := fx.Value.(ValueAutoHelp); !ok {
			continue
		}

		var (
			prefix string
			name   string
		)
		if fx.LongPrefix != "" && fx.LongName != "" {
			prefix, name = fx.LongPrefix, fx.LongName
		} else if fx.ShortPrefix != "" && fx.ShortName != 0 {
			prefix, name = fx.ShortPrefix, string(fx.ShortName)
		}

		// Note that [*FlagSet.AddFlag] rejects empty flags
		runtimex.Assert(prefix != "" && name != "")

		return fs.ProgramName + " " + prefix + name
	}
	return ""
}

func (up *DefaultUsagePrinter) positionalArgumentsUsage(fset *FlagSet) (output string) {
	minArgs, maxArgs := fset.MinPositionalArgs, fset.MaxPositionalArgs
	if maxArgs >= minArgs && maxArgs > 0 {
		output = up.PositionalArgumentsUsage
		switch {
		case output != "":
			// nothing
		case minArgs == 0 && maxArgs == 1:
			output = "[arg]"
		case minArgs == 1 && maxArgs == 1:
			output = "arg"
		case minArgs == 0 && maxArgs > 1:
			output = "[arg ...]"
		default:
			output = "arg [arg ...]"
		}
		output = " " + output
	}
	return
}

// UsagePrinter is the interface used to print the usage.
type UsagePrinter interface {
	PrintUsageString(fs *FlagSet, w io.Writer)
	PrintUsageError(fs *FlagSet, w io.Writer, err error)
}

// Constants controlling text formatting
const (
	wrapAtColumn = 72
	indent4      = "    "
	indent8      = indent4 + indent4
)

// DefaultUsagePrinter is the default [UsagePrinter] implementation.
//
// Construct using [NewDefaultUsagePrinter].
//
// # Usage Format
//
// The default we use follows this pattern:
//
//	Usage
//
//	    curl [flags] [--] URL [URL ...]
//
//	Description
//
//	    Transfer 1+ URLs using the HTTP/HTTPS protocol.
//
//	Flags
//
//	    -o FILE, --output FILE (default: `-`)
//
//	        Write output to the given file.
//
//	    -s, --silent[=true|false]
//
//	        Disable emitting output.
//
//	Examples
//
//	    Fetch the homepage of the www.example.com website:
//
//	        curl -so /dev/null https://www.example.com/
//
// # Help Hint Format
//
// The template we use follows this pattern:
//
//	curl: unknown flag --verbose
//	curl: try `curl --help` for more help.
//
// We only print the help hint on the given [io.Writer] if the user has
// configured a [*Flag] containing a [ValueAutoHelp] [Value].
type DefaultUsagePrinter struct {
	// Description contains the program description paragraphs used when printing the usage.
	//
	// [NewDefaultUsagePrinter] initializes this field to an empty slice.
	//
	// The [*DefaultUsagePrinter.PrintUsageString] method will treat each paragraph as independent
	// and word wrap it to 72 characters removing leading spaces. However, if
	// a paragraph starts with 4 spaces, the method will assume the user intends to
	// emit a verbatim block and will not word wrap it.
	Description []string

	// Example contains the examples paragraphs used when printing the usage.
	//
	// [NewDefaultUsagePrinter] initializes this field to an empty slice.
	//
	// The [*DefaultUsagePrinter.PrintUsageString] method will treat each paragraph as independent
	// and word wrap it to 72 characters removing leading spaces. However, if
	// a paragraph starts with 4 spaces, the method will assume the user intends to
	// emit a verbatim block and will not word wrap it.
	Example []string

	// PositionalArgumentsUsage is the usage string for postional arguments.
	//
	// [NewDefaultUsagePrinter] initializes this field to "". If this value is empty,
	// when printing help we use "", arg" or "args..." depending on whether
	// zero, one, or multiple positional arguments are possible.
	PositionalArgumentsUsage string
}

// PrintUsageString implements [vflag.UsagePrinter].
//
// This method panics on I/O error.
func (up *DefaultUsagePrinter) PrintUsageString(fset *FlagSet, w io.Writer) {
	// ## Usage
	up.div0(w, "Usage")
	up.div0(w, fmt.Sprintf("    %s%s%s", fset.ProgramName, up.flagsName(fset), up.positionalArgumentsUsage(fset)))

	// ## Description
	if description := up.Description; len(description) > 0 {
		up.div0(w, "Description")
		for _, entry := range description {
			up.div1(w, entry)
		}
	}

	// ## Flags
	if flags := fset.Flags(); len(flags) > 0 {
		up.div0(w, "Flags")
		for _, fentry := range flags {
			short, long, value := fentry.UsageShort(), fentry.UsageLong(), fentry.Value
			defaultValue := fmt.Sprintf(" (default: `%s`)", value)
			if _, ok := fentry.Value.(ValueAutoHelp); ok {
				defaultValue = ""
			}
			var formatted string
			switch {
			case short != "" && long != "":
				formatted = fmt.Sprintf("    %s, %s%s", short, long, defaultValue)
			case short != "":
				if _, ok := fentry.Value.(ValueBool); ok {
					defaultValue = ""
				}
				formatted = fmt.Sprintf("    %s%s", short, defaultValue)
			case long != "":
				formatted = fmt.Sprintf("    %s%s", long, defaultValue)
			}
			runtimex.Assert(formatted != "") // [*vflag.FlagSet.AddFlag] rejects empty flags
			up.div0(w, formatted)
			for _, dentry := range fentry.Description {
				up.div0(w, textwrap.Do(dentry, wrapAtColumn, indent8))
			}
		}
	}

	// ## Example
	if example := up.Example; len(example) > 0 {
		up.div0(w, "Examples")
		for _, entry := range example {
			up.div1(w, entry)
		}
	}

	must.Fprintf(w, "\n")
}

// PrintUsageError implements [vflag.UsagePrinter].
//
// This method panics on I/O error.
func (up *DefaultUsagePrinter) PrintUsageError(fset *FlagSet, w io.Writer, err error) {
	programName := fset.ProgramName
	must.Fprintf(w, "%s: %s\n", programName, err.Error())
	if cmdline := fset.HelpInvocation(); cmdline != "" {
		must.Fprintf(w, "%s: try `%s' for more help.\n", programName, cmdline)
	}
}

func (up *DefaultUsagePrinter) div1(w io.Writer, entry string) {
	if strings.HasPrefix(entry, indent4) {
		up.div0(w, indent4+entry)
		return
	}
	up.div0(w, textwrap.Do(entry, wrapAtColumn, indent4))
}

func (up *DefaultUsagePrinter) div0(w io.Writer, value string) {
	must.Fprintf(w, "\n%s\n", value)
}

// NewDefaultUsagePrinter constructs a new [*DefaultUsagePrinter].
func NewDefaultUsagePrinter() *DefaultUsagePrinter {
	return &DefaultUsagePrinter{}
}

// AddDescription adds a paragraph to the current description.
func (up *DefaultUsagePrinter) AddDescription(values ...string) {
	up.Description = append(up.Description, values...)
}

// AddExamples adds a paragraph to the current examples.
func (up *DefaultUsagePrinter) AddExamples(values ...string) {
	up.Example = append(up.Example, values...)
}

var _ UsagePrinter = &DefaultUsagePrinter{}
