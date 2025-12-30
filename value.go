//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/value.go
//

package vflag

import (
	"strconv"
	"time"
)

// Value represents a writable flag value.
//
// Construct using [NewValueBool], [NewValueInt], etc.
type Value interface {
	// Set sets the value of the flag.
	//
	// This method MAY be called multiple times if
	// the command-line flag is repeated.
	Set(value string) error
}

// ValueAutoHelp is a sentinel value associated with the user
// requesting for help using the command line.
type ValueAutoHelp struct{}

var _ Value = ValueAutoHelp{}

// Set implements [Value].
func (v ValueAutoHelp) Set(value string) error {
	if value == "" {
		value = "true"
	}
	_, err := strconv.ParseBool(value)
	return err
}

// ValueBool implements [Value] for bool.
//
// Construct using [NewValueBool].
type ValueBool struct {
	vp *bool
}

// NewValueBool constructs a new [ValueBool] using an underlying bool.
func NewValueBool(vp *bool) ValueBool {
	return ValueBool{vp}
}

var _ Value = ValueBool{}

// Set implements [Value].
func (v ValueBool) Set(value string) error {
	if value == "" {
		value = "true"
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	*v.vp = parsed
	return nil
}

// ValueDuration implements [Value] for [time.Duration].
//
// Construct using [NewValueDuration].
type ValueDuration struct {
	vp *time.Duration
}

// NewValueDuration constructs a new [ValueDuration] using an underlying [time.Duration].
func NewValueDuration(vp *time.Duration) ValueDuration {
	return ValueDuration{vp}
}

var _ Value = ValueDuration{}

// Set implements [Value].
func (v ValueDuration) Set(value string) error {
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*v.vp = parsed
	return nil
}

// ValueFloat64 implements [Value] for float64.
//
// Construct using [NewValueFloat64].
type ValueFloat64 struct {
	vp *float64
}

// NewValueFloat64 constructs a new [ValueFloat64] using an underlying float64.
func NewValueFloat64(vp *float64) ValueFloat64 {
	return ValueFloat64{vp}
}

var _ Value = ValueFloat64{}

// Set implements [Value].
func (v ValueFloat64) Set(value string) error {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	*v.vp = parsed
	return nil
}

// ValueInt implements [Value] for int.
//
// Construct using [NewValueInt].
type ValueInt struct {
	vp *int
}

// NewValueInt constructs a new [ValueInt] using an underlying int.
func NewValueInt(vp *int) ValueInt {
	return ValueInt{vp}
}

var _ Value = ValueInt{}

// Set implements [Value].
func (v ValueInt) Set(value string) error {
	parsed, err := strconv.ParseInt(value, 10, strconv.IntSize)
	if err != nil {
		return err
	}
	*v.vp = int(parsed)
	return nil
}

// ValueInt8 implements [Value] for int8.
//
// Construct using [NewValueInt8].
type ValueInt8 struct {
	vp *int8
}

// NewValueInt8 constructs a new [ValueInt8] using an underlying int8.
func NewValueInt8(vp *int8) ValueInt8 {
	return ValueInt8{vp}
}

var _ Value = ValueInt8{}

// Set implements [Value].
func (v ValueInt8) Set(value string) error {
	parsed, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return err
	}
	*v.vp = int8(parsed)
	return nil
}

// ValueInt16 implements [Value] for int16.
//
// Construct using [NewValueInt16].
type ValueInt16 struct {
	vp *int16
}

// NewValueInt16 constructs a new [ValueInt16] using an underlying int16.
func NewValueInt16(vp *int16) ValueInt16 {
	return ValueInt16{vp}
}

var _ Value = ValueInt16{}

// Set implements [Value].
func (v ValueInt16) Set(value string) error {
	parsed, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return err
	}
	*v.vp = int16(parsed)
	return nil
}

// ValueInt32 implements [Value] for int32.
//
// Construct using [NewValueInt32].
type ValueInt32 struct {
	vp *int32
}

// NewValueInt32 constructs a new [ValueInt32] using an underlying int32.
func NewValueInt32(vp *int32) ValueInt32 {
	return ValueInt32{vp}
}

var _ Value = ValueInt32{}

// Set implements [Value].
func (v ValueInt32) Set(value string) error {
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return err
	}
	*v.vp = int32(parsed)
	return nil
}

// ValueInt64 implements [Value] for int64.
//
// Construct using [NewValueInt64].
type ValueInt64 struct {
	vp *int64
}

// NewValueInt64 constructs a new [ValueInt64] using an underlying int64.
func NewValueInt64(vp *int64) ValueInt64 {
	return ValueInt64{vp}
}

var _ Value = ValueInt64{}

// Set implements [Value].
func (v ValueInt64) Set(value string) error {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	*v.vp = parsed
	return nil
}

// ValueString implements [Value] for string.
//
// Construct using [NewValueString].
type ValueString struct {
	vp *string
}

// NewValueString constructs a new [ValueString] using an underlying string.
func NewValueString(vp *string) ValueString {
	return ValueString{vp}
}

var _ Value = ValueString{}

// Set implements [Value].
func (v ValueString) Set(value string) error {
	*v.vp = value
	return nil
}

// ValueStringSlice implements [Value] for a string slice.
//
// Construct using [NewValueStringSlice].
type ValueStringSlice struct {
	vp *[]string
}

// NewValueStringSlice constructs a new [ValueStringSlice] using an underlying string slice.
func NewValueStringSlice(vp *[]string) ValueStringSlice {
	return ValueStringSlice{vp}
}

var _ Value = ValueStringSlice{}

// Set implements [Value].
func (v ValueStringSlice) Set(value string) error {
	*v.vp = append(*v.vp, value)
	return nil
}

// ValueUint implements [Value] for uint.
//
// Construct using [NewValueUint].
type ValueUint struct {
	vp *uint
}

// NewValueUint constructs a new [ValueUint] using an underlying uint.
func NewValueUint(vp *uint) ValueUint {
	return ValueUint{vp}
}

var _ Value = ValueUint{}

// Set implements [Value].
func (v ValueUint) Set(value string) error {
	parsed, err := strconv.ParseUint(value, 10, strconv.IntSize)
	if err != nil {
		return err
	}
	*v.vp = uint(parsed)
	return nil
}

// ValueUint8 implements [Value] for uint8.
//
// Construct using [NewValueUint8].
type ValueUint8 struct {
	vp *uint8
}

// NewValueUint8 constructs a new [ValueUint8] using an underlying uint8.
func NewValueUint8(vp *uint8) ValueUint8 {
	return ValueUint8{vp}
}

var _ Value = ValueUint8{}

// Set implements [Value].
func (v ValueUint8) Set(value string) error {
	parsed, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return err
	}
	*v.vp = uint8(parsed)
	return nil
}

// ValueUint16 implements [Value] for uint16.
//
// Construct using [NewValueUint16].
type ValueUint16 struct {
	vp *uint16
}

// NewValueUint16 constructs a new [ValueUint16] using an underlying uint16.
func NewValueUint16(vp *uint16) ValueUint16 {
	return ValueUint16{vp}
}

var _ Value = ValueUint16{}

// Set implements [Value].
func (v ValueUint16) Set(value string) error {
	parsed, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return err
	}
	*v.vp = uint16(parsed)
	return nil
}

// ValueUint32 implements [Value] for uint32.
//
// Construct using [NewValueUint32].
type ValueUint32 struct {
	vp *uint32
}

// NewValueUint32 constructs a new [ValueUint32] using an underlying uint32.
func NewValueUint32(vp *uint32) ValueUint32 {
	return ValueUint32{vp}
}

var _ Value = ValueUint32{}

// Set implements [Value].
func (v ValueUint32) Set(value string) error {
	parsed, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return err
	}
	*v.vp = uint32(parsed)
	return nil
}

// ValueUInt64 implements [Value] for uint64.
//
// Construct using [NewValueUInt64].
type ValueUInt64 struct {
	vp *uint64
}

// NewValueUInt64 constructs a new [ValueUInt64] using an underlying uint64.
func NewValueUInt64(vp *uint64) ValueUInt64 {
	return ValueUInt64{vp}
}

var _ Value = ValueUInt64{}

// Set implements [Value].
func (v ValueUInt64) Set(value string) error {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}
	*v.vp = parsed
	return nil
}
