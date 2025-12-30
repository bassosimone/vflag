//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/usage.go
//

package vflag

import (
	"fmt"
	"io"

	"github.com/bassosimone/runtimex"
	"github.com/bassosimone/textwrap"
)

// UsageFlag is [*Flag] as seen by [UsagePrinter].
type UsageFlag struct {
	// Description contains the paragraphs inside the description.
	Description []string

	// IsAutoHelp indicates that this is the auto help flag.
	IsAutoHelp bool

	// Long is the long flag to use (e.g. "--verbose[=BOOL]", "--file FILE")
	Long string

	// Short is the short flag to use (e.g. "-v", "-f FILE")
	Short string

	// Value is the current flag value.
	Value string
}

// UsageFlagSet is a [*FlagSet] view that allows obtaining the
// strings for printing the usage using a [UsagePrinter].
type UsageFlagSet struct {
	Set *FlagSet
}

// Description returns the paragraphs inside the description.
func (ufs UsageFlagSet) Description() []string {
	return ufs.Set.Description
}

// Example returns the paragraphs inside the example.
func (ufs UsageFlagSet) Example() []string {
	return ufs.Set.Example
}

// Flags returns the flags we should print.
func (ufs UsageFlagSet) Flags() (output []UsageFlag) {
	for _, entry := range ufs.Set.flags {
		var longOption string
		if entry.LongName != "" && entry.LongPrefix != "" {
			longArgumentName := entry.LongArgumentName
			longOption = fmt.Sprintf("%s%s%s", entry.LongPrefix, entry.LongName, longArgumentName)
		}

		var shortOption string
		if entry.ShortName != 0 && entry.ShortPrefix != "" {
			shortArgumentName := entry.ShortArgumentName
			shortOption = fmt.Sprintf("%s%s%s", entry.ShortPrefix, string(entry.ShortName), shortArgumentName)
		}

		if longOption == "" && shortOption == "" {
			continue
		}

		value := entry.Value.String()
		_, isAutoHelp := entry.Value.(ValueAutoHelp)
		output = append(output, UsageFlag{
			Description: entry.Description,
			IsAutoHelp:  isAutoHelp,
			Long:        longOption,
			Short:       shortOption,
			Value:       value,
		})
	}
	return
}

func (ufs UsageFlagSet) FlagsName() (output string) {
	if f := ufs.Set.flags; len(f) > 0 {
		output = " [flags]"
	}
	return output
}

// HelpFlag returns the flag to use to get the help screen.
//
// The return value is empty if [*FlagSet.AutoHelp] has not been used.
func (ufs UsageFlagSet) HelpFlag() string {
	for _, fx := range ufs.Set.flags {
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
		} else {
			continue
		}

		return ufs.Set.ProgramName + " " + prefix + name
	}
	return ""
}

// PositionalArgumentsUsage returns the string to print for
// the positional arguments (e.g. "", " [args]").
func (ufs UsageFlagSet) PositionalArgumentsUsage() (output string) {
	minArgs, maxArgs := ufs.Set.MinPositionalArgs, ufs.Set.MaxPositionalArgs
	if minArgs >= 0 && maxArgs > 0 && maxArgs >= minArgs {
		output = ufs.Set.PositionalArgumentsUsage
		switch {
		case output == "" && minArgs == 0 && maxArgs == 1:
			output = "arg"
		case output == "" && minArgs == 0 && maxArgs > 1:
			output = "[args...]"
		case output == "" && minArgs > 0 && maxArgs > 1:
			output = "arg [arg ...]"
		}
		output = " " + output
	}
	return
}

// ProgramName returns the program name.
func (ufs UsageFlagSet) ProgramName() string {
	return ufs.Set.ProgramName
}

// UsagePrinter is the interface used to print the usage.
type UsagePrinter interface {
	PrintUsage(w io.Writer, fs UsageFlagSet)
	PrintHelpHint(w io.Writer, fs UsageFlagSet)
}

// DefaultUsagePrinter is the default [UsagePrinter] implementation.
//
// # Usage Format
//
// The default template we use follows this pattern:
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
//	    -o FILE
//
//	    --output FILE
//
//	        Write output to the given file.
//
//	    -s
//
//	    --silent[=true|false]
//
//	        Disable emitting output.
//
//	Examples
//
//	    curl -so /dev/null https://www.example.com/
//
//	        Fetches the homepage of the www.example.com website.
//
// # Help Hint Format
//
// The default template we use follows this pattern:
//
//	hint: Try `curl --help' for more help.\n
//
// We only print something on the given [io.Writer] if the user has
// configured a [*Flag] containing a [ValueAutoHelp] [Value].
type DefaultUsagePrinter struct{}

// PrintUsage implements [UsagePrinter].
//
// This method panics on I/O error.
func (p DefaultUsagePrinter) PrintUsage(w io.Writer, fs UsageFlagSet) {
	const wrapAtColumn = 72

	p.div0(w, "Usage")
	p.div0(w, fmt.Sprintf("    %s%s%s", fs.ProgramName(), fs.FlagsName(), fs.PositionalArgumentsUsage()))

	if description := fs.Description(); len(description) > 0 {
		p.div0(w, "Description")
		for _, dentry := range description {
			p.div0(w, textwrap.Do(dentry, wrapAtColumn, "    "))
		}
	}

	if flags := fs.Flags(); len(flags) > 0 {
		p.div0(w, "Flags")
		for _, fentry := range flags {
			short, long, value := fentry.Short, fentry.Long, fentry.Value
			defaultValue := fmt.Sprintf(" (default: `%s`)", value)
			if fentry.IsAutoHelp {
				defaultValue = ""
			}
			switch {
			case short != "" && long != "":
				p.div0(w, fmt.Sprintf("    %s, %s%s", short, long, defaultValue))
			case short != "":
				p.div0(w, fmt.Sprintf("    %s%s", short, defaultValue))
			case long != "":
				p.div0(w, fmt.Sprintf("    %s%s", long, defaultValue))
			default:
				runtimex.Assert(false)
			}
			for _, dentry := range fentry.Description {
				p.div0(w, textwrap.Do(dentry, wrapAtColumn, "        "))
			}
		}
	}

	if example := fs.Example(); len(example) > 0 {
		p.div0(w, "Examples")
		for _, eentry := range example {
			p.div0(w, textwrap.Do(eentry, wrapAtColumn, "    "))
		}
	}

	_ = runtimex.PanicOnError1(fmt.Fprintf(w, "\n"))
}

// PrintHelpHint implements [UsagePrinter].
//
// This method panics on I/O error.
func (p DefaultUsagePrinter) PrintHelpHint(w io.Writer, fs UsageFlagSet) {
	if hf := fs.HelpFlag(); hf != "" {
		_ = runtimex.PanicOnError1(fmt.Fprintf(w, "hint: try `%s' for more help.\n", hf))
	}
}

func (p DefaultUsagePrinter) div0(w io.Writer, value string) {
	_ = runtimex.PanicOnError1(fmt.Fprintf(w, "\n%s\n", value))
}
