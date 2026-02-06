//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/flagset.go
//

package vflag

import (
	"fmt"

	"github.com/bassosimone/flagparser"
	"github.com/bassosimone/runtimex"
)

// LongFlag represents a long flag to parse.
//
// Long flags are multi-character flags (e.g., `--verbose`, `--output`) that cannot
// be grouped together. They are standalone and can accept values either with an
// equals sign (`--output=FILE`) or as a separate argument (`--output FILE`).
//
// The first backtick-quoted uppercase name in the first Description entry (e.g.,
// "Write to `FILE`.") overrides the default ArgumentName in help output.
//
// The placeholder @DEFAULT_VALUE@ in Description entries is replaced with the
// current default value (via Value.String()) when printing help.
//
// Construct using [NewLongFlagBool], [NewLongFlagString], etc.
type LongFlag struct {
	// Description contains the flag description paragraphs to use in the help.
	Description []string

	// ArgumentName is the name of the argument to use in the help.
	ArgumentName string

	// DefaultValue is the default value to use when the flag is present but no
	// value is provided. This is only used by [LongFlagMakeOptionWithOptionalValue].
	// The value is captured at construction time from the bound variable.
	DefaultValue string

	// MakeOption constructs the [*flagparser.Option] to use.
	MakeOption func(fx *LongFlag) *flagparser.Option

	// Name is the flag long name.
	Name string

	// Prefix is the flag long prefix.
	Prefix string

	// Value is the flag [Value].
	Value Value
}

// Usage returns the usage string for the [*LongFlag].
//
// For example: `--verbose` or `--output FILE`.
func (fx *LongFlag) Usage() string {
	argumentName := argumentNameFromDocsOrDefault(fx.Description, fx.ArgumentName)
	return fmt.Sprintf("%s%s%s", fx.Prefix, fx.Name, argumentName)
}

// LongFlagMakeOptionAutoHelp returns the [*flagparser.Option] to use for auto help.
//
// This method panics if the name or prefix are empty.
func LongFlagMakeOptionAutoHelp(fx *LongFlag) *flagparser.Option {
	runtimex.Assert(fx.Prefix != "" && fx.Name != "")
	return &flagparser.Option{
		Type:   flagparser.OptionTypeEarlyArgumentNone,
		Prefix: fx.Prefix,
		Name:   fx.Name,
	}
}

// NewLongFlagAutoHelp constructs a new [*LongFlag] bound to a [ValueAutoHelp].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
func NewLongFlagAutoHelp(value ValueAutoHelp, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: "",
		Name:         name,
		MakeOption:   LongFlagMakeOptionAutoHelp,
		Prefix:       "--",
		Value:        value,
	}
}

// LongFlagMakeOptionBool returns the [*flagparser.Option] to use for booleans.
//
// Long boolean flags are standalone and take an optional argument (e.g., `--verbose`,
// `--verbose=true`, `--verbose=false`). When no argument is provided, the default
// value is "true".
//
// This method panics if the name or prefix are empty.
func LongFlagMakeOptionBool(fx *LongFlag) *flagparser.Option {
	runtimex.Assert(fx.Prefix != "" && fx.Name != "")
	return &flagparser.Option{
		Type:         flagparser.OptionTypeStandaloneArgumentOptional,
		Prefix:       fx.Prefix,
		Name:         fx.Name,
		DefaultValue: "true",
	}
}

// NewLongFlagBool constructs a new [*LongFlag] bound to a [ValueBool].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to `[=true|false]` by default.
func NewLongFlagBool(value ValueBool, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: "[=true|false]",
		Name:         name,
		MakeOption:   LongFlagMakeOptionBool,
		Prefix:       "--",
		Value:        value,
	}
}

// LongFlagMakeOptionWithRequiredValue returns the [*flagparser.Option] to use for
// flags that require a value.
//
// Long flags with required values are standalone and require an argument
// (e.g., `--output FILE`, `--output=FILE`).
//
// This method panics if the name or prefix are empty.
func LongFlagMakeOptionWithRequiredValue(fx *LongFlag) *flagparser.Option {
	runtimex.Assert(fx.Prefix != "" && fx.Name != "")
	return &flagparser.Option{
		Type:   flagparser.OptionTypeStandaloneArgumentRequired,
		Prefix: fx.Prefix,
		Name:   fx.Name,
	}
}

// LongFlagMakeOptionWithOptionalValue returns the [*flagparser.Option] to use for
// flags that take an optional value.
//
// Long flags with optional values are standalone and accept an optional argument
// (e.g., `+https` uses the default, `+https=/custom` uses the provided value).
// The value must be provided using the `=` syntax; otherwise the default is used.
//
// This method panics if the name or prefix are empty.
func LongFlagMakeOptionWithOptionalValue(fx *LongFlag) *flagparser.Option {
	runtimex.Assert(fx.Prefix != "" && fx.Name != "")
	return &flagparser.Option{
		Type:         flagparser.OptionTypeStandaloneArgumentOptional,
		Prefix:       fx.Prefix,
		Name:         fx.Name,
		DefaultValue: fx.DefaultValue,
	}
}

// NewLongFlagDuration constructs a new [*LongFlag] bound to a [ValueDuration].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` DURATION` by default.
func NewLongFlagDuration(value ValueDuration, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " DURATION",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagFloat64 constructs a new [*LongFlag] bound to a [ValueFloat64].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` FLOAT64` by default.
func NewLongFlagFloat64(value ValueFloat64, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " FLOAT64",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagInt constructs a new [*LongFlag] bound to a [ValueInt].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` INT` by default.
func NewLongFlagInt(value ValueInt, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " INT",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagInt8 constructs a new [*LongFlag] bound to a [ValueInt8].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` INT8` by default.
func NewLongFlagInt8(value ValueInt8, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " INT8",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagInt16 constructs a new [*LongFlag] bound to a [ValueInt16].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` INT16` by default.
func NewLongFlagInt16(value ValueInt16, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " INT16",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagInt32 constructs a new [*LongFlag] bound to a [ValueInt32].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` INT32` by default.
func NewLongFlagInt32(value ValueInt32, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " INT32",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagInt64 constructs a new [*LongFlag] bound to a [ValueInt64].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` INT64` by default.
func NewLongFlagInt64(value ValueInt64, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " INT64",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagString constructs a new [*LongFlag] bound to a [ValueString].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` STRING` by default.
func NewLongFlagString(value ValueString, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " STRING",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagStringSlice constructs a new [*LongFlag] bound to a [ValueStringSlice].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` STRING` by default.
func NewLongFlagStringSlice(value ValueStringSlice, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " STRING",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagUint constructs a new [*LongFlag] bound to a [ValueUint].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` UINT` by default.
func NewLongFlagUint(value ValueUint, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " UINT",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagUint8 constructs a new [*LongFlag] bound to a [ValueUint8].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` UINT8` by default.
func NewLongFlagUint8(value ValueUint8, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " UINT8",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagUint16 constructs a new [*LongFlag] bound to a [ValueUint16].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` UINT16` by default.
func NewLongFlagUint16(value ValueUint16, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " UINT16",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagUint32 constructs a new [*LongFlag] bound to a [ValueUint32].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` UINT32` by default.
func NewLongFlagUint32(value ValueUint32, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " UINT32",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}

// NewLongFlagUint64 constructs a new [*LongFlag] bound to a [ValueUint64].
//
// This constructor sets the flag prefix to `--`. If you need a different prefix,
// update the `Prefix` field in the returned [*LongFlag] structure.
//
// The ArgumentName is set to ` UINT64` by default.
func NewLongFlagUint64(value ValueUint64, name string, helpText ...string) *LongFlag {
	return &LongFlag{
		Description:  helpText,
		ArgumentName: " UINT64",
		Name:         name,
		MakeOption:   LongFlagMakeOptionWithRequiredValue,
		Prefix:       "--",
		Value:        value,
	}
}
