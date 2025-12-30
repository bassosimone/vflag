//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/flagset.go
//

package vflag

import "github.com/bassosimone/flagparser"

// Flag represents a flag to parse.
//
// Construct using [NewFlagBool], [NewFlagInt], etc or implicitly
// through [*FlagSet.BoolVar], [*FlagSet.IntVar], etc.
type Flag struct {
	// Description contains the flag description paragraphs to use in the help.
	Description []string

	// LongArgumentName is the name of the long argument to use in the help.
	LongArgumentName string

	// LongName is the flag long name.
	LongName string

	// LongPrefix is the flag long prefix.
	LongPrefix string

	// MakeOptions constructs the [*flagparser.Option] slice to use.
	MakeOptions func(fx *Flag) []*flagparser.Option

	// ShortArgumentName is the name of the short argument to use in the help.
	ShortArgumentName string

	// ShortName is the flag short name.
	ShortName byte

	// ShortPrefix is the flag short prefix.
	ShortPrefix string

	// Value is the flag value.
	Value Value
}

var _ Value = &Flag{}

// Set implements [Value].
func (fx *Flag) Set(value string) error {
	return fx.Value.Set(value)
}

// FlagMakeOptionsAutoHelp returns the slice of [*flagparser.Option] to use for auto help.
//
// The return value has zero length if the [*Flag] argument is misconfigured.
func FlagMakeOptionsAutoHelp(fx *Flag) []*flagparser.Option {
	var options []*flagparser.Option
	if fx.ShortName != 0 && fx.ShortPrefix != "" {
		options = append(options, &flagparser.Option{
			Type:   flagparser.OptionTypeEarlyArgumentNone,
			Prefix: fx.ShortPrefix,
			Name:   string(fx.ShortName),
		})
	}
	if fx.LongName != "" && fx.LongPrefix != "" {
		options = append(options, &flagparser.Option{
			Type:   flagparser.OptionTypeEarlyArgumentNone,
			Prefix: fx.LongPrefix,
			Name:   fx.LongName,
		})
	}
	return options
}

// NewFlagAutoHelp constructs a new [*Flag] bound to a [ValueAutoHelp].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.AutoHelp] as a shortcut.
func NewFlagAutoHelp(value ValueAutoHelp, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  "",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsAutoHelp,
		ShortArgumentName: "",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// FlagMakeOptionsBool returns the slice of [*flagparser.Option] to use for booleans.
//
// The return value has zero length if the [*Flag] argument is misconfigured.
func FlagMakeOptionsBool(fx *Flag) []*flagparser.Option {
	var options []*flagparser.Option
	if fx.ShortName != 0 && fx.ShortPrefix != "" {
		options = append(options, &flagparser.Option{
			Type:   flagparser.OptionTypeGroupableArgumentNone,
			Prefix: fx.ShortPrefix,
			Name:   string(fx.ShortName),
		})
	}
	if fx.LongName != "" && fx.LongPrefix != "" {
		options = append(options, &flagparser.Option{
			Type:         flagparser.OptionTypeStandaloneArgumentOptional,
			Prefix:       fx.LongPrefix,
			Name:         fx.LongName,
			DefaultValue: "true",
		})
	}
	return options
}

// NewFlagBool constructs a new [*Flag] bound to a [ValueBool].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName to `[=BOOL]` and the ShortArgumentName
// to an empty values, and you may want to override these assignments.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.BoolVar] as a shortcut.
func NewFlagBool(value ValueBool, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  "[=BOOL]",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsBool,
		ShortArgumentName: "",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// FlagMakeOptionsWithValue returns the slice of [*flagparser.Option] to use for non-booleans.
//
// The return value has zero length if the [*Flag] argument is misconfigured.
func FlagMakeOptionsWithValue(fx *Flag) []*flagparser.Option {
	var options []*flagparser.Option
	if fx.ShortName != 0 && fx.ShortPrefix != "" {
		options = append(options, &flagparser.Option{
			Type:   flagparser.OptionTypeGroupableArgumentRequired,
			Prefix: fx.ShortPrefix,
			Name:   string(fx.ShortName),
		})
	}
	if fx.LongName != "" && fx.LongPrefix != "" {
		options = append(options, &flagparser.Option{
			Type:   flagparser.OptionTypeStandaloneArgumentRequired,
			Prefix: fx.LongPrefix,
			Name:   fx.LongName,
		})
	}
	return options
}

// NewFlagDuration constructs a new [*Flag] bound to a [ValueDuration].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` DURATION` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.DurationVar] as a shortcut.
func NewFlagDuration(value ValueDuration, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " DURATION",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " DURATION",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagFloat64 constructs a new [*Flag] bound to a [ValueFloat64].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` FLOAT64` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Float64Var] as a shortcut.
func NewFlagFloat64(value ValueFloat64, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " FLOAT64",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " FLOAT64",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagInt constructs a new [*Flag] bound to a [ValueInt].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` INT` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.IntVar] as a shortcut.
func NewFlagInt(value ValueInt, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " INT",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " INT",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagInt8 constructs a new [*Flag] bound to a [ValueInt8].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` INT8` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Int8Var] as a shortcut.
func NewFlagInt8(value ValueInt8, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " INT8",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " INT8",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagInt16 constructs a new [*Flag] bound to a [ValueInt16].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` INT16` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Int16Var] as a shortcut.
func NewFlagInt16(value ValueInt16, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " INT16",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " INT16",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagInt32 constructs a new [*Flag] bound to a [ValueInt32].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` INT32` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Int32Var] as a shortcut.
func NewFlagInt32(value ValueInt32, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " INT32",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " INT32",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagInt64 constructs a new [*Flag] bound to a [ValueInt64].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` INT64` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Int64Var] as a shortcut.
func NewFlagInt64(value ValueInt64, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " INT64",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " INT64",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagString constructs a new [*Flag] bound to a [ValueString].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` STRING` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.StringVar] as a shortcut.
func NewFlagString(value ValueString, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " STRING",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " STRING",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagStringSlice constructs a new [*Flag] bound to a [ValueStringSlice].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` STRING` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.StringSliceVar] as a shortcut.
func NewFlagStringSlice(value ValueStringSlice, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " STRING",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " STRING",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagUint constructs a new [*Flag] bound to a [ValueUint].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` UINT` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.UintVar] as a shortcut.
func NewFlagUint(value ValueUint, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " UINT",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " UINT",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagUint8 constructs a new [*Flag] bound to a [ValueUint8].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` UINT8` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Uint8Var] as a shortcut.
func NewFlagUint8(value ValueUint8, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " UINT8",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " UINT8",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagUint16 constructs a new [*Flag] bound to a [ValueUint16].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` UINT16` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Uint16Var] as a shortcut.
func NewFlagUint16(value ValueUint16, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " UINT16",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " UINT16",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagUint32 constructs a new [*Flag] bound to a [ValueUint32].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` UINT32` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Uint32Var] as a shortcut.
func NewFlagUint32(value ValueUint32, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " UINT32",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " UINT32",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}

// NewFlagUInt64 constructs a new [*Flag] bound to a [ValueUInt64].
//
// This constructor sets the long prefix to `--` and the short prefix
// to `-` following the GNU convention. If you need different prefixes,
// modify the related fields in the returned [*Flag] structure.
//
// Likewise, we set the LongArgumentName and ShortArgumentName to
// ` UINT64` and you may want to override this default.
//
// A zero shortName or an empty longName cause the respective short/long
// flag to not be configured. If both are empty, no flag is ever set.
//
// Consider using [*FlagSet.Uint64Var] as a shortcut.
func NewFlagUInt64(value ValueUInt64, shortName byte, longName string, helpText ...string) *Flag {
	return &Flag{
		Description:       helpText,
		LongArgumentName:  " UINT64",
		LongName:          longName,
		LongPrefix:        "--",
		MakeOptions:       FlagMakeOptionsWithValue,
		ShortArgumentName: " UINT64",
		ShortName:         shortName,
		ShortPrefix:       "-",
		Value:             value,
	}
}
