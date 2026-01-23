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
// Note: a [*FlagSet] where you have not added any flags defaults to parsing
// options using GNU conventions. That is, if you run the related program with:
//
//	./program --verbose
//
// The [*FlagSet] will recognize `--verbose` as a syntactically valid flag
// that has not been configured and print an "unknown flag" error.
type FlagSet struct {
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

	// ErrorHandling is the [ErrorHandling] policy.
	//
	// [NewFlagSet] initializes this field to [ContinueOnError].
	ErrorHandling ErrorHandling

	// Exit is the function to call with the [ExitOnError] policy.
	//
	// [NewFlagSet] initializes this field to [os.Exit].
	Exit func(status int)

	// LongFlags contains the long flags to parse.
	//
	// Long flags are multi-character flags (e.g., `--verbose`, `--output`)
	// that cannot be grouped together.
	LongFlags []*LongFlag

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

	// ProgramName is the program name.
	//
	// [NewFlagSet] initializes this field to the given program name.
	ProgramName string

	// ShortFlags contains the short flags to parse.
	//
	// Short flags are single-character flags (e.g., `-v`, `-o`) that can be
	// grouped together on the command line.
	ShortFlags []*ShortFlag

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
	// [NewFlagSet] initializes this field to an empty [*DefaultUsagePrinter]
	// instance. Explicitly construct a [UsagePrinter], for example using
	// [NewDefaultUsagePrinter], and assign the field, to customize the help.
	//
	// We use this field with [ExitOnError] policy.
	UsagePrinter UsagePrinter

	// positionals buffers the positional arguments.
	positionals []string
}

// NewFlagSet returns a new [*FlagSet] instance. We use the given progname as
// the ProgramName field and the given handling as the ErrorHandling field. We
// initialize all the other fields using sensible defaults. We document these
// defaults in the [*FlagSet] documentation.
func NewFlagSet(progname string, handling ErrorHandling) *FlagSet {
	const (
		expectedLongFlags   = 16
		expectedPositionals = 8
		expectedShortFlags  = 16
	)
	return &FlagSet{
		DisablePermute:            false,
		ErrorHandling:             handling,
		Exit:                      os.Exit,
		LongFlags:                 make([]*LongFlag, 0, expectedLongFlags),
		MaxPositionalArgs:         0,
		MinPositionalArgs:         0,
		OptionsArgumentsSeparator: "--",
		ProgramName:               progname,
		ShortFlags:                make([]*ShortFlag, 0, expectedShortFlags),
		Stderr:                    os.Stderr,
		Stdout:                    os.Stdout,
		UsagePrinter:              &DefaultUsagePrinter{},
		positionals:               make([]string, 0, expectedPositionals),
	}
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

// Parse parses the given command line arguments, It assigns positional arguments
// and each flag [Value] as a side effect of parsing.
//
// The args MUST NOT contain the program name. That is, if there are no command
// line arguments beyond the program name, arguments must be empty.
//
// Depending on the [ErrorHandling] policy, on failure, this method may return the
// error, invoke [os.Exit], or call panic with the error that occurred.
//
// This method panics if a long flag has the same name as a short flag.
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

	// build options and value map from short flags
	pview := make(map[string]Value)
	for _, fx := range fs.ShortFlags {
		opt := fx.MakeOption(fx)
		px.Options = append(px.Options, opt)
		pview[opt.Name] = fx.Value
	}

	// build options and value map from long flags
	for _, fx := range fs.LongFlags {
		opt := fx.MakeOption(fx)
		_, found := pview[opt.Name]
		runtimex.Assert(!found)
		px.Options = append(px.Options, opt)
		pview[opt.Name] = fx.Value
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
			optname := value.Option.Name
			val, found := pview[optname]
			runtimex.Assert(found) // should not happen

			// assign a value to the flag
			if err := val.Set(value.Value); err != nil {
				return err
			}

			// detect [ValueAutoHelp] and transform it to [ErrHelp]
			if _, ok := val.(ValueAutoHelp); ok {
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
