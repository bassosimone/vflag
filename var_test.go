// SPDX-License-Identifier: GPL-3.0-or-later

package vflag

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertFlagSetVar(t *testing.T, fs *FlagSet, got *Flag, expected *Flag) {
	t.Helper()

	require.Len(t, fs.Flags(), 1)
	assert.Same(t, got, fs.Flags()[0])

	// We need to sweat a bit because function equality is tricky and we need to
	// perform this specific comparison separately
	gotCopy := *got
	expectedCopy := *expected
	gotCopy.MakeOptions = nil
	expectedCopy.MakeOptions = nil
	assert.Equal(t, expectedCopy, gotCopy)
	assertFuncPointerEqual(t, expected.MakeOptions, got.MakeOptions)
}

func assertFuncPointerEqual(t *testing.T, expected any, got any) {
	t.Helper()

	expectedValue := reflect.ValueOf(expected)
	gotValue := reflect.ValueOf(got)
	if assert.True(t, expectedValue.IsValid(), "expected function is invalid") &&
		assert.True(t, gotValue.IsValid(), "got function is invalid") {
		assert.Equal(t, expectedValue.Pointer(), gotValue.Pointer())
	}
}

func TestFlagSetVarAutoHelp(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	got := fs.AutoHelp('h', "help", "Print help and exit.")

	expected := NewFlagAutoHelp(ValueAutoHelp{}, 'h', "help", "Print help and exit.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarBool(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value bool
	got := fs.BoolVar(&value, 'v', "verbose", "Enable verbose output.")

	expected := NewFlagBool(NewValueBool(&value), 'v', "verbose", "Enable verbose output.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarDuration(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value time.Duration
	got := fs.DurationVar(&value, 't', "timeout", "Set timeout.")

	expected := NewFlagDuration(NewValueDuration(&value), 't', "timeout", "Set timeout.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarFloat64(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value float64
	got := fs.Float64Var(&value, 'r', "ratio", "Set ratio.")

	expected := NewFlagFloat64(NewValueFloat64(&value), 'r', "ratio", "Set ratio.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarInt(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value int
	got := fs.IntVar(&value, 'n', "count", "Set count.")

	expected := NewFlagInt(NewValueInt(&value), 'n', "count", "Set count.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarInt8(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value int8
	got := fs.Int8Var(&value, 'b', "batch", "Set batch size.")

	expected := NewFlagInt8(NewValueInt8(&value), 'b', "batch", "Set batch size.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarInt16(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value int16
	got := fs.Int16Var(&value, 'p', "port", "Set port.")

	expected := NewFlagInt16(NewValueInt16(&value), 'p', "port", "Set port.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarInt32(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value int32
	got := fs.Int32Var(&value, 'i', "index", "Set index.")

	expected := NewFlagInt32(NewValueInt32(&value), 'i', "index", "Set index.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarInt64(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value int64
	got := fs.Int64Var(&value, 's', "size", "Set size.")

	expected := NewFlagInt64(NewValueInt64(&value), 's', "size", "Set size.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarString(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value string
	got := fs.StringVar(&value, 'o', "output", "Set output file.")

	expected := NewFlagString(NewValueString(&value), 'o', "output", "Set output file.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarStringSlice(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value []string
	got := fs.StringSliceVar(&value, 'H', "header", "Set header.")

	expected := NewFlagStringSlice(NewValueStringSlice(&value), 'H', "header", "Set header.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarUint(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value uint
	got := fs.UintVar(&value, 'u', "users", "Set users.")

	expected := NewFlagUint(NewValueUint(&value), 'u', "users", "Set users.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarUint8(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value uint8
	got := fs.Uint8Var(&value, 'q', "queue", "Set queue size.")

	expected := NewFlagUint8(NewValueUint8(&value), 'q', "queue", "Set queue size.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarUint16(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value uint16
	got := fs.Uint16Var(&value, 'm', "max", "Set max.")

	expected := NewFlagUint16(NewValueUint16(&value), 'm', "max", "Set max.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarUint32(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value uint32
	got := fs.Uint32Var(&value, 'c', "cache", "Set cache size.")

	expected := NewFlagUint32(NewValueUint32(&value), 'c', "cache", "Set cache size.")
	assertFlagSetVar(t, fs, got, expected)
}

func TestFlagSetVarUint64(t *testing.T) {
	fs := NewFlagSet("prog", ContinueOnError)
	var value uint64
	got := fs.Uint64Var(&value, 'l', "limit", "Set limit.")

	expected := NewFlagUInt64(NewValueUInt64(&value), 'l', "limit", "Set limit.")
	assertFlagSetVar(t, fs, got, expected)
}
