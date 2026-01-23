//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/flagset.go
//

package vflag

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bassosimone/flagparser"
	"github.com/bassosimone/runtimex"
)

// ShortFlag represents a short flag to parse.
//
// Short flags are single-character flags (e.g., `-v`, `-o`) that can be grouped
// together on the command line. For example, `-xvf` is equivalent to `-x -v -f`.
// When grouped, only the last flag in the group can take an argument (e.g., `-xvf FILE`).
//
// The first backtick-quoted uppercase name in the first Description entry (e.g.,
// "Write to `FILE`.") overrides the default ArgumentName in help output.
//
// The placeholder @DEFAULT_VALUE@ in Description entries is replaced with the
// current default value (via Value.String()) when printing help.
//
// Construct using [NewShortFlagBool], [NewShortFlagString], etc.
type ShortFlag struct {
	// Description contains the flag description paragraphs to use in the help.
	Description []string

	// ArgumentName is the name of the argument to use in the help.
	ArgumentName string

	// MakeOption constructs the [*flagparser.Option] to use.
	MakeOption func(fx *ShortFlag) *flagparser.Option

	// Name is the flag short name.
	Name byte

	// Prefix is the flag short prefix.
	Prefix string

	// Value is the flag [Value].
	Value Value
}

// argumentNameFromDocsOrDefault returns the `<name>` inside the first string in the
// documentation, if available, and otherwise returns the configured default.
func argumentNameFromDocsOrDefault(description []string, defaultValue string) (output string) {
	output = defaultValue
	if len(description) > 0 && strings.HasPrefix(output, " ") {
		re := regexp.MustCompile("`([A-Z0-9_-]+)`")
		m := re.FindStringSubmatch(description[0])
		if len(m) > 1 {
			output = " " + m[1]
		}
	}
	return
}

// Usage returns the short usage string for the [*ShortFlag].
//
// For example: `-v` or `-t TAG`.
func (fx *ShortFlag) Usage() string {
	argumentName := argumentNameFromDocsOrDefault(fx.Description, fx.ArgumentName)
	return fmt.Sprintf("%s%s%s", fx.Prefix, string(fx.Name), argumentName)
}

// ShortFlagMakeOptionAutoHelp returns the [*flagparser.Option] to use for auto help.
//
// This method panics if the name or prefix are empty.
func ShortFlagMakeOptionAutoHelp(fx *ShortFlag) *flagparser.Option {
	runtimex.Assert(fx.Prefix != "" && fx.Name != 0)
	return &flagparser.Option{
		Type:   flagparser.OptionTypeEarlyArgumentNone,
		Prefix: fx.Prefix,
		Name:   string(fx.Name),
	}
}

// NewShortFlagAutoHelp constructs a new [*ShortFlag] bound to a [ValueAutoHelp].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
func NewShortFlagAutoHelp(value ValueAutoHelp, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: "",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionAutoHelp,
		Prefix:       "-",
		Value:        value,
	}
}

// ShortFlagMakeOptionBool returns the [*flagparser.Option] to use for booleans.
//
// Short boolean flags are groupable and take no argument (e.g., `-v`, `-xvz`).
//
// This method panics if the name or prefix are empty.
func ShortFlagMakeOptionBool(fx *ShortFlag) *flagparser.Option {
	runtimex.Assert(fx.Prefix != "" && fx.Name != 0)
	return &flagparser.Option{
		Type:   flagparser.OptionTypeGroupableArgumentNone,
		Prefix: fx.Prefix,
		Name:   string(fx.Name),
	}
}

// NewShortFlagBool constructs a new [*ShortFlag] bound to a [ValueBool].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
func NewShortFlagBool(value ValueBool, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: "",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionBool,
		Prefix:       "-",
		Value:        value,
	}
}

// ShortFlagMakeOptionWithValue returns the [*flagparser.Option] to use for
// flags that require a value.
//
// Short flags with values are groupable and require an argument (e.g., `-o FILE`, `-xzf FILE`).
//
// This method panics if the name or prefix are empty.
func ShortFlagMakeOptionWithValue(fx *ShortFlag) *flagparser.Option {
	runtimex.Assert(fx.Prefix != "" && fx.Name != 0)
	return &flagparser.Option{
		Type:   flagparser.OptionTypeGroupableArgumentRequired,
		Prefix: fx.Prefix,
		Name:   string(fx.Name),
	}
}

// NewShortFlagDuration constructs a new [*ShortFlag] bound to a [ValueDuration].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` DURATION` by default.
func NewShortFlagDuration(value ValueDuration, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " DURATION",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagFloat64 constructs a new [*ShortFlag] bound to a [ValueFloat64].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` FLOAT64` by default.
func NewShortFlagFloat64(value ValueFloat64, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " FLOAT64",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagInt constructs a new [*ShortFlag] bound to a [ValueInt].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` INT` by default.
func NewShortFlagInt(value ValueInt, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " INT",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagInt8 constructs a new [*ShortFlag] bound to a [ValueInt8].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` INT8` by default.
func NewShortFlagInt8(value ValueInt8, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " INT8",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagInt16 constructs a new [*ShortFlag] bound to a [ValueInt16].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` INT16` by default.
func NewShortFlagInt16(value ValueInt16, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " INT16",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagInt32 constructs a new [*ShortFlag] bound to a [ValueInt32].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` INT32` by default.
func NewShortFlagInt32(value ValueInt32, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " INT32",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagInt64 constructs a new [*ShortFlag] bound to a [ValueInt64].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` INT64` by default.
func NewShortFlagInt64(value ValueInt64, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " INT64",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagString constructs a new [*ShortFlag] bound to a [ValueString].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` STRING` by default.
func NewShortFlagString(value ValueString, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " STRING",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagStringSlice constructs a new [*ShortFlag] bound to a [ValueStringSlice].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` STRING` by default.
func NewShortFlagStringSlice(value ValueStringSlice, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " STRING",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagUint constructs a new [*ShortFlag] bound to a [ValueUint].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` UINT` by default.
func NewShortFlagUint(value ValueUint, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " UINT",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagUint8 constructs a new [*ShortFlag] bound to a [ValueUint8].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` UINT8` by default.
func NewShortFlagUint8(value ValueUint8, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " UINT8",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagUint16 constructs a new [*ShortFlag] bound to a [ValueUint16].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` UINT16` by default.
func NewShortFlagUint16(value ValueUint16, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " UINT16",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagUint32 constructs a new [*ShortFlag] bound to a [ValueUint32].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` UINT32` by default.
func NewShortFlagUint32(value ValueUint32, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " UINT32",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}

// NewShortFlagUint64 constructs a new [*ShortFlag] bound to a [ValueUInt64].
//
// This constructor sets the flag prefix to `-`. If you need a different prefix,
// update the `Prefix` field in the returned [*ShortFlag] structure.
//
// The ArgumentName is set to ` UINT64` by default.
func NewShortFlagUint64(value ValueUInt64, name byte, helpText ...string) *ShortFlag {
	return &ShortFlag{
		Description:  helpText,
		ArgumentName: " UINT64",
		Name:         name,
		MakeOption:   ShortFlagMakeOptionWithValue,
		Prefix:       "-",
		Value:        value,
	}
}
