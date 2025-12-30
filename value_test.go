// SPDX-License-Identifier: GPL-3.0-or-later

package vflag

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValueAutoHelp(t *testing.T) {
	value := ValueAutoHelp{}

	assert.Equal(t, "false", value.String())
	require.NoError(t, value.Set(""))
	assert.Equal(t, "false", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "false", value.String())
}

func TestValueBool(t *testing.T) {
	var raw bool
	value := NewValueBool(&raw)

	assert.Equal(t, "false", value.String())
	require.NoError(t, value.Set("true"))
	assert.Equal(t, "true", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "true", value.String())
}

func TestValueDuration(t *testing.T) {
	var raw time.Duration
	value := NewValueDuration(&raw)

	assert.Equal(t, "0s", value.String())
	require.NoError(t, value.Set("3s"))
	assert.Equal(t, "3s", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "3s", value.String())
}

func TestValueFloat64(t *testing.T) {
	var raw float64
	value := NewValueFloat64(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("1.25"))
	assert.Equal(t, "1.25", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "1.25", value.String())
}

func TestValueInt(t *testing.T) {
	var raw int
	value := NewValueInt(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("12"))
	assert.Equal(t, "12", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "12", value.String())
}

func TestValueInt8(t *testing.T) {
	var raw int8
	value := NewValueInt8(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "7", value.String())
}

func TestValueInt16(t *testing.T) {
	var raw int16
	value := NewValueInt16(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "7", value.String())
}

func TestValueInt32(t *testing.T) {
	var raw int32
	value := NewValueInt32(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "7", value.String())
}

func TestValueInt64(t *testing.T) {
	var raw int64
	value := NewValueInt64(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("nope"))
	assert.Equal(t, "7", value.String())
}

func TestValueString(t *testing.T) {
	var raw string
	value := NewValueString(&raw)

	assert.Equal(t, "", value.String())
	require.NoError(t, value.Set("hello"))
	assert.Equal(t, "hello", value.String())

	require.NoError(t, value.Set("not-a-number"))
	assert.Equal(t, "not-a-number", value.String())
}

func TestValueStringSlice(t *testing.T) {
	var raw []string
	value := NewValueStringSlice(&raw)

	assert.Equal(t, "", value.String())
	require.NoError(t, value.Set("a"))
	assert.Equal(t, "a", value.String())

	require.NoError(t, value.Set("not-a-number"))
	assert.Equal(t, "a,not-a-number", value.String())
}

func TestValueUint(t *testing.T) {
	var raw uint
	value := NewValueUint(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("-1"))
	assert.Equal(t, "7", value.String())
}

func TestValueUint8(t *testing.T) {
	var raw uint8
	value := NewValueUint8(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("-1"))
	assert.Equal(t, "7", value.String())
}

func TestValueUint16(t *testing.T) {
	var raw uint16
	value := NewValueUint16(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("-1"))
	assert.Equal(t, "7", value.String())
}

func TestValueUint32(t *testing.T) {
	var raw uint32
	value := NewValueUint32(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("-1"))
	assert.Equal(t, "7", value.String())
}

func TestValueUInt64(t *testing.T) {
	var raw uint64
	value := NewValueUInt64(&raw)

	assert.Equal(t, "0", value.String())
	require.NoError(t, value.Set("7"))
	assert.Equal(t, "7", value.String())

	require.Error(t, value.Set("-1"))
	assert.Equal(t, "7", value.String())
}
