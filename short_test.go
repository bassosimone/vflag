// SPDX-License-Identifier: GPL-3.0-or-later

package vflag

import (
	"testing"
	"time"

	"github.com/bassosimone/flagparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortFlagUsage(t *testing.T) {
	t.Run("bool flag has no argument name", func(t *testing.T) {
		var v bool
		sf := NewShortFlagBool(NewValueBool(&v), 'v', "Enable verbose output.")
		assert.Equal(t, "-v", sf.Usage())
	})

	t.Run("string flag has argument name", func(t *testing.T) {
		var v string
		sf := NewShortFlagString(NewValueString(&v), 'o', "Write to file.")
		assert.Equal(t, "-o STRING", sf.Usage())
	})

	t.Run("argument name extracted from description", func(t *testing.T) {
		var v string
		sf := NewShortFlagString(NewValueString(&v), 'o', "Write to `FILE`.")
		assert.Equal(t, "-o FILE", sf.Usage())
	})

	t.Run("custom prefix", func(t *testing.T) {
		var v bool
		sf := NewShortFlagBool(NewValueBool(&v), 'v', "Enable verbose output.")
		sf.Prefix = "+"
		assert.Equal(t, "+v", sf.Usage())
	})
}

func TestArgumentNameFromDocsOrDefault(t *testing.T) {
	t.Run("empty description uses default", func(t *testing.T) {
		result := argumentNameFromDocsOrDefault(nil, " STRING")
		assert.Equal(t, " STRING", result)
	})

	t.Run("description without backticks uses default", func(t *testing.T) {
		result := argumentNameFromDocsOrDefault([]string{"Write to file."}, " STRING")
		assert.Equal(t, " STRING", result)
	})

	t.Run("description with backticks extracts name", func(t *testing.T) {
		result := argumentNameFromDocsOrDefault([]string{"Write to `FILE`."}, " STRING")
		assert.Equal(t, " FILE", result)
	})

	t.Run("backticks with hyphen", func(t *testing.T) {
		result := argumentNameFromDocsOrDefault([]string{"Set `URL-PATH`."}, " STRING")
		assert.Equal(t, " URL-PATH", result)
	})

	t.Run("backticks with underscore", func(t *testing.T) {
		result := argumentNameFromDocsOrDefault([]string{"Set `FILE_NAME`."}, " STRING")
		assert.Equal(t, " FILE_NAME", result)
	})

	t.Run("backticks with numbers", func(t *testing.T) {
		result := argumentNameFromDocsOrDefault([]string{"Set `VALUE123`."}, " STRING")
		assert.Equal(t, " VALUE123", result)
	})

	t.Run("only applies when default has leading space", func(t *testing.T) {
		// When default doesn't have leading space, extraction is skipped
		result := argumentNameFromDocsOrDefault([]string{"Write to `FILE`."}, "STRING")
		assert.Equal(t, "STRING", result)
	})
}

func TestShortFlagMakeOptionAutoHelp(t *testing.T) {
	sf := NewShortFlagAutoHelp(ValueAutoHelp{}, 'h', "Show help.")
	opt := sf.MakeOption(sf)

	require.NotNil(t, opt)
	assert.Equal(t, flagparser.OptionTypeEarlyArgumentNone, opt.Type)
	assert.Equal(t, "-", opt.Prefix)
	assert.Equal(t, "h", opt.Name)
}

func TestShortFlagMakeOptionBool(t *testing.T) {
	var v bool
	sf := NewShortFlagBool(NewValueBool(&v), 'v', "Enable verbose.")
	opt := sf.MakeOption(sf)

	require.NotNil(t, opt)
	assert.Equal(t, flagparser.OptionTypeGroupableArgumentNone, opt.Type)
	assert.Equal(t, "-", opt.Prefix)
	assert.Equal(t, "v", opt.Name)
}

func TestShortFlagMakeOptionWithValue(t *testing.T) {
	var v string
	sf := NewShortFlagString(NewValueString(&v), 'o', "Output file.")
	opt := sf.MakeOption(sf)

	require.NotNil(t, opt)
	assert.Equal(t, flagparser.OptionTypeGroupableArgumentRequired, opt.Type)
	assert.Equal(t, "-", opt.Prefix)
	assert.Equal(t, "o", opt.Name)
}

func TestNewShortFlagAutoHelp(t *testing.T) {
	sf := NewShortFlagAutoHelp(ValueAutoHelp{}, 'h', "Show help.", "Extra info.")

	assert.Equal(t, byte('h'), sf.Name)
	assert.Equal(t, "-", sf.Prefix)
	assert.Equal(t, "", sf.ArgumentName)
	assert.Equal(t, []string{"Show help.", "Extra info."}, sf.Description)
	_, ok := sf.Value.(ValueAutoHelp)
	assert.True(t, ok)
}

func TestNewShortFlagBool(t *testing.T) {
	var v bool
	sf := NewShortFlagBool(NewValueBool(&v), 'v', "Enable verbose.")

	assert.Equal(t, byte('v'), sf.Name)
	assert.Equal(t, "-", sf.Prefix)
	assert.Equal(t, "", sf.ArgumentName)
}

func TestNewShortFlagDuration(t *testing.T) {
	var v time.Duration
	sf := NewShortFlagDuration(NewValueDuration(&v), 't', "Set timeout.")

	assert.Equal(t, byte('t'), sf.Name)
	assert.Equal(t, "-", sf.Prefix)
	assert.Equal(t, " DURATION", sf.ArgumentName)
}

func TestNewShortFlagFloat64(t *testing.T) {
	var v float64
	sf := NewShortFlagFloat64(NewValueFloat64(&v), 'r', "Set ratio.")

	assert.Equal(t, byte('r'), sf.Name)
	assert.Equal(t, " FLOAT64", sf.ArgumentName)
}

func TestNewShortFlagInt(t *testing.T) {
	var v int
	sf := NewShortFlagInt(NewValueInt(&v), 'n', "Set count.")

	assert.Equal(t, byte('n'), sf.Name)
	assert.Equal(t, " INT", sf.ArgumentName)
}

func TestNewShortFlagInt8(t *testing.T) {
	var v int8
	sf := NewShortFlagInt8(NewValueInt8(&v), 'b', "Set batch.")

	assert.Equal(t, byte('b'), sf.Name)
	assert.Equal(t, " INT8", sf.ArgumentName)
}

func TestNewShortFlagInt16(t *testing.T) {
	var v int16
	sf := NewShortFlagInt16(NewValueInt16(&v), 'p', "Set port.")

	assert.Equal(t, byte('p'), sf.Name)
	assert.Equal(t, " INT16", sf.ArgumentName)
}

func TestNewShortFlagInt32(t *testing.T) {
	var v int32
	sf := NewShortFlagInt32(NewValueInt32(&v), 'i', "Set index.")

	assert.Equal(t, byte('i'), sf.Name)
	assert.Equal(t, " INT32", sf.ArgumentName)
}

func TestNewShortFlagInt64(t *testing.T) {
	var v int64
	sf := NewShortFlagInt64(NewValueInt64(&v), 's', "Set size.")

	assert.Equal(t, byte('s'), sf.Name)
	assert.Equal(t, " INT64", sf.ArgumentName)
}

func TestNewShortFlagString(t *testing.T) {
	var v string
	sf := NewShortFlagString(NewValueString(&v), 'o', "Set output.")

	assert.Equal(t, byte('o'), sf.Name)
	assert.Equal(t, " STRING", sf.ArgumentName)
}

func TestNewShortFlagStringSlice(t *testing.T) {
	var v []string
	sf := NewShortFlagStringSlice(NewValueStringSlice(&v), 'H', "Set header.")

	assert.Equal(t, byte('H'), sf.Name)
	assert.Equal(t, " STRING", sf.ArgumentName)
}

func TestNewShortFlagUint(t *testing.T) {
	var v uint
	sf := NewShortFlagUint(NewValueUint(&v), 'u', "Set users.")

	assert.Equal(t, byte('u'), sf.Name)
	assert.Equal(t, " UINT", sf.ArgumentName)
}

func TestNewShortFlagUint8(t *testing.T) {
	var v uint8
	sf := NewShortFlagUint8(NewValueUint8(&v), 'q', "Set queue.")

	assert.Equal(t, byte('q'), sf.Name)
	assert.Equal(t, " UINT8", sf.ArgumentName)
}

func TestNewShortFlagUint16(t *testing.T) {
	var v uint16
	sf := NewShortFlagUint16(NewValueUint16(&v), 'm', "Set max.")

	assert.Equal(t, byte('m'), sf.Name)
	assert.Equal(t, " UINT16", sf.ArgumentName)
}

func TestNewShortFlagUint32(t *testing.T) {
	var v uint32
	sf := NewShortFlagUint32(NewValueUint32(&v), 'c', "Set cache.")

	assert.Equal(t, byte('c'), sf.Name)
	assert.Equal(t, " UINT32", sf.ArgumentName)
}

func TestNewShortFlagUint64(t *testing.T) {
	var v uint64
	sf := NewShortFlagUint64(NewValueUInt64(&v), 'l', "Set limit.")

	assert.Equal(t, byte('l'), sf.Name)
	assert.Equal(t, " UINT64", sf.ArgumentName)
}

func TestShortFlagMakeOptionPanicsOnEmptyPrefix(t *testing.T) {
	var v bool
	sf := NewShortFlagBool(NewValueBool(&v), 'v', "Verbose.")
	sf.Prefix = ""

	assert.Panics(t, func() {
		sf.MakeOption(sf)
	})
}

func TestShortFlagMakeOptionPanicsOnZeroName(t *testing.T) {
	var v bool
	sf := NewShortFlagBool(NewValueBool(&v), 'v', "Verbose.")
	sf.Name = 0

	assert.Panics(t, func() {
		sf.MakeOption(sf)
	})
}
