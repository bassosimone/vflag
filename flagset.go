//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/flagset.go
//

package vflag

import (
	"errors"
	"io"
	"os"

	"github.com/bassosimone/flagparser"
	"github.com/bassosimone/runtimex"
)

// ErrorHandling controls [*FlagSet.Parse] error handling.
type ErrorHandling int

// These constants define the allowed [ErrorHandling] values.
const (
	// ContinueOnError causes [*FlagSet] to return the parse error.
	ContinueOnError = ErrorHandling(iota)

	// ExitOnError causes [*Flagset] to call Exit with code 2 on error.
	ExitOnError

	// PanicOnError causes [*FlagSet] to panic on error.
	PanicOnError
)

// FlagSet allows to parse flags from the command line. The zero value is not
// ready to use. Construct using the [NewFlagSet] constructor.
//
// Note: a [*FlagSet] where you have not added any flags through methods
// like [*FlagSet.BoolVar], [*FlagSet.StringVar], etc defaults to parsing options
// using GNU conventions. That is, if you run the related program with:
//
//	./program --verbose
//
// The [*FlagSet] will recognize `--verbose` as a syntactically valid flag
// that has not been configured and print an "unknown flag" error.
type FlagSet struct {
	// Description contains the program description paragraphs used when printing the usage.
	//
	// [NewFlagSet] initializes this field to an empty slice.
	Description []string

	// DisablePermute disable the permutation of options and arguments.
	//
	// [NewFlagSet] initializes this field to false.
	//
	// Consider the following command line args:
	//
	// 	https://www.google.com/ -H 'Host: google.com'
	//
	// The default behavior is to permute this to:
	//
	// 	-H 'Host: google.com' https://www.google.com/
	//
	// However, when DisablePermute is true, we keep the command
	// line unmodified. While permuting is a nice-to-have property
	// in general, consider instead the following case:
	//
	// 	foreach -kx git status -v
	//
	// With permutation, this command line would become:
	//
	// 	-kx -v foreach git status
	//
	// This is not the desired behavior if the foreach command
	// takes another command and its options as arguments.
	//
	// To make the above command line work with permutation, a
	// user would instead need to write this:
	//
	// 	foreach -kx -- git status -v
	//
	// By setting DisablePermute to true, the `--` separator
	// becomes unnecessary and the UX is improved.
	DisablePermute bool

	// Example contains the examples paragraphs used when printing the usage.
	//
	// [NewFlagSet] initializes this field to an empty slice.
	//
	// The [*FlagSet.PrintUsage] method will treat each paragraph as independent
	// and word wrap it to 72 characters removing leading spaces. However, if
	// a paragraph starts with 4 spaces, the method will assume the user intends to
	// emit a verbatim block and will not word wrap it.
	Example []string

	// Exit is the function to call with the [ExitOnError] policy.
	//
	// [NewFlagSet] initializes this field to [os.Exit].
	Exit func(status int)

	// ErrorHandling is the [ErrorHandling] policy.
	//
	// [NewFlagSet] initializes this field to [ContinueOnError].
	ErrorHandling ErrorHandling

	// MaxPositionalArgs is the maximum number of positional arguments.
	//
	// [NewFlagSet] initializes this field to 0.
	//
	// The default configuration, thus, allows for no positional
	// arguments to be on the command line.
	MaxPositionalArgs int

	// MinPositionalArgs is the minimum number of positional arguments.
	//
	// [NewFlagSet] initializes this field to 0.
	//
	// The default configuration, thus, allows for no positional
	// arguments to be on the command line.
	MinPositionalArgs int

	// OptionsArgumentsSeparator separates options and arguments.
	//
	// [NewFlagSet] initializes this field to "--".
	//
	// The default configuration is compatible with the GNU standards
	// where "--" instructs getopt to stop processing flags and to treat
	// all the remaining entries as positional arguments.
	OptionsArgumentsSeparator string

	// PositionalArgumentsUsage is the usage string for postional arguments.
	//
	// [NewFlagSet] initializes this field to "". If this value is empty,
	// when printing help we use "", arg" or "args..." depending on whether
	// zero, one, or multiple positional arguments are possible.
	PositionalArgumentsUsage string

	// ProgramName is the program name.
	//
	// [NewFlagSet] initializes this field to the given program name.
	ProgramName string

	// Stderr is the [io.Writer] to use as the stderr.
	//
	// [NewFlagSet] initializes this field to [os.Stderr].
	//
	// We use this field with [ExitOnError] policy.
	Stderr io.Writer

	// Stdout is the [io.Writer] to use as the stdout.
	//
	// [NewFlagSet] initializes this field to [os.Stdout].
	//
	// We use this field with [ExitOnError] policy.
	Stdout io.Writer

	// UsagePrinter is the [UsagePrinter] to use.
	//
	// [NewFlagSet] initializes this field to [DefaultUsagePrinter].
	//
	// We use this field with [ExitOnError] policy.
	UsagePrinter UsagePrinter

	// flags contains the [*Flag] to parse.
	flags []*Flag

	// positional buffers the positional arguments.
	positionals []string
}

// NewFlagSet returns a new [*FlagSet] instance. We use the given progname as
// the ProgramName field and the given handling as the ErrorHandling field. We
// initialize all the other fields using sensible defaults. We document these
// defaults in the [*FlagSet] documentation.
func NewFlagSet(progname string, handling ErrorHandling) *FlagSet {
	const (
		expectedFlags       = 32
		expectedPositionals = 8
	)
	return &FlagSet{
		Description:               []string{},
		DisablePermute:            false,
		Example:                   []string{},
		Exit:                      os.Exit,
		ErrorHandling:             handling,
		MaxPositionalArgs:         0,
		MinPositionalArgs:         0,
		OptionsArgumentsSeparator: "--",
		PositionalArgumentsUsage:  "",
		ProgramName:               progname,
		Stderr:                    os.Stderr,
		Stdout:                    os.Stdout,
		UsagePrinter:              DefaultUsagePrinter{},
		flags:                     make([]*Flag, 0, expectedFlags),
		positionals:               make([]string, 0, expectedPositionals),
	}
}

// AddDescription adds a paragraph to the current description.
func (fs *FlagSet) AddDescription(values ...string) {
	fs.Description = append(fs.Description, values...)
}

// AddExamples adds a paragraph to the current examples.
func (fs *FlagSet) AddExamples(values ...string) {
	fs.Example = append(fs.Example, values...)
}

// SetMinMaxPositionalArgs sets the minimum and maximum positional arguments.
func (fs *FlagSet) SetMinMaxPositionalArgs(minArgs, maxArgs int) {
	fs.MinPositionalArgs = minArgs
	fs.MaxPositionalArgs = maxArgs
}

// Args returns the positional arguments collected by [*FlagSet.Parse].
func (fs *FlagSet) Args() []string {
	return fs.positionals
}

// Flags returns the [*Flag] configured so far.
func (fs *FlagSet) Flags() []*Flag {
	return fs.flags
}

// AddFlag adds the given [*Flag] to the [*FlagSet].
func (fs *FlagSet) AddFlag(fx *Flag) {
	fs.flags = append(fs.flags, fx)
}

// Parse parses the given command line arguments, It assigns positional arguments
// and each flag [Value] as a side effect of parsing.
//
// The args MUST NOT contain the program name. That is, if there are no command
// line arguments beyond the program name, arguments must be empty.
//
// Depending on the [ErrorHandling] policy, on failure, this method may return the
// error, invoke [os.Exit], or call panic with the error that occurred.
func (fs *FlagSet) Parse(args []string) error {
	return fs.maybeHandleError(fs.parse(args))
}

// ErrHelp is the error returned in case the user requested for `help`.
//
// Use [*FlagSet.AutoHelp] to enable recognizing help flags.
//
// This error is never returned when using the [ExitOnError] policy.
var ErrHelp = errors.New("help requested")

func (fs *FlagSet) parse(args []string) error {
	// configure the command line parser
	px := &flagparser.Parser{
		DisablePermute:            fs.DisablePermute,
		MaxPositionalArguments:    fs.MaxPositionalArgs,
		MinPositionalArguments:    fs.MinPositionalArgs,
		OptionsArgumentsSeparator: fs.OptionsArgumentsSeparator,
		Options:                   []*flagparser.Option{},
	}
	pview := make(map[string]*Flag)
	for _, fx := range fs.flags {
		px.Options = append(px.Options, fx.MakeOptions(fx)...)
		if fx.ShortPrefix != "" && fx.ShortName != 0 {
			pview[string(fx.ShortName)] = fx
		}
		if fx.LongPrefix != "" && fx.LongName != "" {
			pview[fx.LongName] = fx
		}
	}

	// parse the command line
	values, err := px.Parse(args)
	if err != nil {
		return err
	}

	// map the parsed values back to options and positionals
	for _, value := range values {
		switch value := value.(type) {

		// positional argument: just add to the internal slice of positionals
		case flagparser.ValuePositionalArgument:
			fs.positionals = append(fs.positionals, value.Value)

		// option: find the corresponding value and attempt to set it
		case flagparser.ValueOption:
			// attempt to get the right parser view
			optname := value.Option.Name
			flag, found := pview[optname]
			runtimex.Assert(found) // should not happen

			// assign a value to the flag
			if err := flag.Value.Set(value.Value); err != nil {
				return err
			}

			// detect [ValueAutoHelp] and transform it to [ErrHelp]
			if _, ok := flag.Value.(ValueAutoHelp); ok {
				return ErrHelp
			}
		}
	}
	return nil
}

func (fs *FlagSet) maybeHandleError(err error) error {
	switch {
	case err == nil:
		return nil

	case fs.ErrorHandling == ContinueOnError:
		return err

	case fs.ErrorHandling == ExitOnError && errors.Is(err, ErrHelp):
		fs.PrintUsageString(fs.Stdout)
		fs.Exit(0)

	case fs.ErrorHandling == ExitOnError:
		fs.PrintUsageError(fs.Stderr, err)
		fs.Exit(2)
	}

	// We end up here for [PanicOnError] or whenever fs.Exit is so
	// broken that it does not actually exit.
	panic(err)
}
