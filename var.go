//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/flagset.go
//

package vflag

import "time"

// AutoHelp combines [NewFlagAutoHelp] and [*FlagSet.AddFlag].
func (fs *FlagSet) AutoHelp(shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagAutoHelp(ValueAutoHelp{}, shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// BoolVar combines [NewFlagBool] and [*FlagSet.AddFlag].
func (fs *FlagSet) BoolVar(vp *bool, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagBool(NewValueBool(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// DurationVar combines [NewFlagDuration] and [*FlagSet.AddFlag].
func (fs *FlagSet) DurationVar(vp *time.Duration, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagDuration(NewValueDuration(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Float64Var combines [NewFlagFloat64] and [*FlagSet.AddFlag].
func (fs *FlagSet) Float64Var(vp *float64, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagFloat64(NewValueFloat64(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// IntVar combines [NewFlagInt] and [*FlagSet.AddFlag].
func (fs *FlagSet) IntVar(vp *int, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagInt(NewValueInt(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Int8Var combines [NewFlagInt8] and [*FlagSet.AddFlag].
func (fs *FlagSet) Int8Var(vp *int8, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagInt8(NewValueInt8(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Int16Var combines [NewFlagInt16] and [*FlagSet.AddFlag].
func (fs *FlagSet) Int16Var(vp *int16, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagInt16(NewValueInt16(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Int32Var combines [NewFlagInt32] and [*FlagSet.AddFlag].
func (fs *FlagSet) Int32Var(vp *int32, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagInt32(NewValueInt32(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Int64Var combines [NewFlagInt64] and [*FlagSet.AddFlag].
func (fs *FlagSet) Int64Var(vp *int64, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagInt64(NewValueInt64(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// StringVar combines [NewFlagString] and [*FlagSet.AddFlag].
func (fs *FlagSet) StringVar(vp *string, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagString(NewValueString(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// StringSliceVar combines [NewFlagStringSlice] and [*FlagSet.AddFlag].
func (fs *FlagSet) StringSliceVar(vp *[]string, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagStringSlice(NewValueStringSlice(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// UintVar combines [NewFlagUint] and [*FlagSet.AddFlag].
func (fs *FlagSet) UintVar(vp *uint, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagUint(NewValueUint(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Uint8Var combines [NewFlagUint8] and [*FlagSet.AddFlag].
func (fs *FlagSet) Uint8Var(vp *uint8, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagUint8(NewValueUint8(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Uint16Var combines [NewFlagUint16] and [*FlagSet.AddFlag].
func (fs *FlagSet) Uint16Var(vp *uint16, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagUint16(NewValueUint16(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Uint32Var combines [NewFlagUint32] and [*FlagSet.AddFlag].
func (fs *FlagSet) Uint32Var(vp *uint32, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagUint32(NewValueUint32(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}

// Uint64Var combines [NewFlagUInt64] and [*FlagSet.AddFlag].
func (fs *FlagSet) Uint64Var(vp *uint64, shortName byte, longName string, helpText ...string) *Flag {
	fx := NewFlagUInt64(NewValueUInt64(vp), shortName, longName, helpText...)
	fs.AddFlag(fx)
	return fx
}
