//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/flagset.go
//

package vflag

import "time"

// AutoHelp registers auto-help flags using GNU conventions.
//
// If shortName is not zero, a short flag (e.g., `-h`) is added to ShortFlags.
// If longName is not empty, a long flag (e.g., `--help`) is added to LongFlags.
func (fs *FlagSet) AutoHelp(shortName byte, longName string, helpText ...string) {
	value := ValueAutoHelp{}
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagAutoHelp(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagAutoHelp(value, longName, helpText...))
	}
}

// BoolVar registers boolean flags using GNU conventions.
//
// If shortName is not zero, a short flag (e.g., `-v`) is added to ShortFlags.
// If longName is not empty, a long flag (e.g., `--verbose`) is added to LongFlags.
func (fs *FlagSet) BoolVar(vp *bool, shortName byte, longName string, helpText ...string) {
	value := NewValueBool(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagBool(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagBool(value, longName, helpText...))
	}
}

// DurationVar registers duration flags using GNU conventions.
//
// If shortName is not zero, a short flag (e.g., `-t`) is added to ShortFlags.
// If longName is not empty, a long flag (e.g., `--timeout`) is added to LongFlags.
func (fs *FlagSet) DurationVar(vp *time.Duration, shortName byte, longName string, helpText ...string) {
	value := NewValueDuration(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagDuration(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagDurationRequired(value, longName, helpText...))
	}
}

// Float64Var registers float64 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Float64Var(vp *float64, shortName byte, longName string, helpText ...string) {
	value := NewValueFloat64(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagFloat64(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagFloat64Required(value, longName, helpText...))
	}
}

// IntVar registers int flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) IntVar(vp *int, shortName byte, longName string, helpText ...string) {
	value := NewValueInt(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagInt(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagIntRequired(value, longName, helpText...))
	}
}

// Int8Var registers int8 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Int8Var(vp *int8, shortName byte, longName string, helpText ...string) {
	value := NewValueInt8(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagInt8(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagInt8Required(value, longName, helpText...))
	}
}

// Int16Var registers int16 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Int16Var(vp *int16, shortName byte, longName string, helpText ...string) {
	value := NewValueInt16(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagInt16(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagInt16Required(value, longName, helpText...))
	}
}

// Int32Var registers int32 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Int32Var(vp *int32, shortName byte, longName string, helpText ...string) {
	value := NewValueInt32(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagInt32(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagInt32Required(value, longName, helpText...))
	}
}

// Int64Var registers int64 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Int64Var(vp *int64, shortName byte, longName string, helpText ...string) {
	value := NewValueInt64(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagInt64(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagInt64Required(value, longName, helpText...))
	}
}

// StringVar registers string flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) StringVar(vp *string, shortName byte, longName string, helpText ...string) {
	value := NewValueString(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagString(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagStringRequired(value, longName, helpText...))
	}
}

// StringSliceVar registers string slice flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) StringSliceVar(vp *[]string, shortName byte, longName string, helpText ...string) {
	value := NewValueStringSlice(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagStringSlice(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagStringSliceRequired(value, longName, helpText...))
	}
}

// UintVar registers uint flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) UintVar(vp *uint, shortName byte, longName string, helpText ...string) {
	value := NewValueUint(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagUint(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagUintRequired(value, longName, helpText...))
	}
}

// Uint8Var registers uint8 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Uint8Var(vp *uint8, shortName byte, longName string, helpText ...string) {
	value := NewValueUint8(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagUint8(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagUint8Required(value, longName, helpText...))
	}
}

// Uint16Var registers uint16 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Uint16Var(vp *uint16, shortName byte, longName string, helpText ...string) {
	value := NewValueUint16(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagUint16(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagUint16Required(value, longName, helpText...))
	}
}

// Uint32Var registers uint32 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Uint32Var(vp *uint32, shortName byte, longName string, helpText ...string) {
	value := NewValueUint32(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagUint32(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagUint32Required(value, longName, helpText...))
	}
}

// Uint64Var registers uint64 flags using GNU conventions.
//
// If shortName is not zero, a short flag is added to ShortFlags.
// If longName is not empty, a long flag is added to LongFlags.
func (fs *FlagSet) Uint64Var(vp *uint64, shortName byte, longName string, helpText ...string) {
	value := NewValueUInt64(vp)
	if shortName != 0 {
		fs.ShortFlags = append(fs.ShortFlags, NewShortFlagUint64(value, shortName, helpText...))
	}
	if longName != "" {
		fs.LongFlags = append(fs.LongFlags, NewLongFlagUint64Required(value, longName, helpText...))
	}
}
