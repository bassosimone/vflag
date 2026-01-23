// SPDX-License-Identifier: GPL-3.0-or-later

package vflag

import (
	"testing"
	"time"

	"github.com/bassosimone/flagparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLongFlagUsage(t *testing.T) {
	t.Run("bool flag has optional argument", func(t *testing.T) {
		var v bool
		lf := NewLongFlagBool(NewValueBool(&v), "verbose", "Enable verbose output.")
		assert.Equal(t, "--verbose[=true|false]", lf.Usage())
	})

	t.Run("string flag has argument name", func(t *testing.T) {
		var v string
		lf := NewLongFlagStringRequired(NewValueString(&v), "output", "Write to file.")
		assert.Equal(t, "--output STRING", lf.Usage())
	})

	t.Run("argument name extracted from description", func(t *testing.T) {
		var v string
		lf := NewLongFlagStringRequired(NewValueString(&v), "output", "Write to `FILE`.")
		assert.Equal(t, "--output FILE", lf.Usage())
	})

	t.Run("custom prefix", func(t *testing.T) {
		var v bool
		lf := NewLongFlagBool(NewValueBool(&v), "verbose", "Enable verbose output.")
		lf.Prefix = "+"
		assert.Equal(t, "+verbose[=true|false]", lf.Usage())
	})

	t.Run("optional string flag", func(t *testing.T) {
		var v string = "default"
		lf := NewLongFlagStringOptional(NewValueString(&v), "config", "Config path.")
		assert.Equal(t, "--config[=STRING]", lf.Usage())
	})
}

func TestLongFlagMakeOptionAutoHelp(t *testing.T) {
	lf := NewLongFlagAutoHelp(ValueAutoHelp{}, "help", "Show help.")
	opt := lf.MakeOption(lf)

	require.NotNil(t, opt)
	assert.Equal(t, flagparser.OptionTypeEarlyArgumentNone, opt.Type)
	assert.Equal(t, "--", opt.Prefix)
	assert.Equal(t, "help", opt.Name)
}

func TestLongFlagMakeOptionBool(t *testing.T) {
	var v bool
	lf := NewLongFlagBool(NewValueBool(&v), "verbose", "Enable verbose.")
	opt := lf.MakeOption(lf)

	require.NotNil(t, opt)
	assert.Equal(t, flagparser.OptionTypeStandaloneArgumentOptional, opt.Type)
	assert.Equal(t, "--", opt.Prefix)
	assert.Equal(t, "verbose", opt.Name)
	assert.Equal(t, "true", opt.DefaultValue)
}

func TestLongFlagMakeOptionWithRequiredValue(t *testing.T) {
	var v string
	lf := NewLongFlagStringRequired(NewValueString(&v), "output", "Output file.")
	opt := lf.MakeOption(lf)

	require.NotNil(t, opt)
	assert.Equal(t, flagparser.OptionTypeStandaloneArgumentRequired, opt.Type)
	assert.Equal(t, "--", opt.Prefix)
	assert.Equal(t, "output", opt.Name)
}

func TestLongFlagMakeOptionWithOptionalValue(t *testing.T) {
	var v string = "/dns-query"
	lf := NewLongFlagStringOptional(NewValueString(&v), "https", "Enable HTTPS.")
	opt := lf.MakeOption(lf)

	require.NotNil(t, opt)
	assert.Equal(t, flagparser.OptionTypeStandaloneArgumentOptional, opt.Type)
	assert.Equal(t, "--", opt.Prefix)
	assert.Equal(t, "https", opt.Name)
	assert.Equal(t, "/dns-query", opt.DefaultValue)
}

func TestNewLongFlagAutoHelp(t *testing.T) {
	lf := NewLongFlagAutoHelp(ValueAutoHelp{}, "help", "Show help.", "Extra info.")

	assert.Equal(t, "help", lf.Name)
	assert.Equal(t, "--", lf.Prefix)
	assert.Equal(t, "", lf.ArgumentName)
	assert.Equal(t, []string{"Show help.", "Extra info."}, lf.Description)
	_, ok := lf.Value.(ValueAutoHelp)
	assert.True(t, ok)
}

func TestNewLongFlagBool(t *testing.T) {
	var v bool
	lf := NewLongFlagBool(NewValueBool(&v), "verbose", "Enable verbose.")

	assert.Equal(t, "verbose", lf.Name)
	assert.Equal(t, "--", lf.Prefix)
	assert.Equal(t, "[=true|false]", lf.ArgumentName)
}

func TestNewLongFlagDurationRequired(t *testing.T) {
	var v time.Duration
	lf := NewLongFlagDurationRequired(NewValueDuration(&v), "timeout", "Set timeout.")

	assert.Equal(t, "timeout", lf.Name)
	assert.Equal(t, "--", lf.Prefix)
	assert.Equal(t, " DURATION", lf.ArgumentName)
}

func TestNewLongFlagFloat64Required(t *testing.T) {
	var v float64
	lf := NewLongFlagFloat64Required(NewValueFloat64(&v), "ratio", "Set ratio.")

	assert.Equal(t, "ratio", lf.Name)
	assert.Equal(t, " FLOAT64", lf.ArgumentName)
}

func TestNewLongFlagIntRequired(t *testing.T) {
	var v int
	lf := NewLongFlagIntRequired(NewValueInt(&v), "count", "Set count.")

	assert.Equal(t, "count", lf.Name)
	assert.Equal(t, " INT", lf.ArgumentName)
}

func TestNewLongFlagInt8Required(t *testing.T) {
	var v int8
	lf := NewLongFlagInt8Required(NewValueInt8(&v), "batch", "Set batch.")

	assert.Equal(t, "batch", lf.Name)
	assert.Equal(t, " INT8", lf.ArgumentName)
}

func TestNewLongFlagInt16Required(t *testing.T) {
	var v int16
	lf := NewLongFlagInt16Required(NewValueInt16(&v), "port", "Set port.")

	assert.Equal(t, "port", lf.Name)
	assert.Equal(t, " INT16", lf.ArgumentName)
}

func TestNewLongFlagInt32Required(t *testing.T) {
	var v int32
	lf := NewLongFlagInt32Required(NewValueInt32(&v), "index", "Set index.")

	assert.Equal(t, "index", lf.Name)
	assert.Equal(t, " INT32", lf.ArgumentName)
}

func TestNewLongFlagInt64Required(t *testing.T) {
	var v int64
	lf := NewLongFlagInt64Required(NewValueInt64(&v), "size", "Set size.")

	assert.Equal(t, "size", lf.Name)
	assert.Equal(t, " INT64", lf.ArgumentName)
}

func TestNewLongFlagStringRequired(t *testing.T) {
	var v string
	lf := NewLongFlagStringRequired(NewValueString(&v), "output", "Set output.")

	assert.Equal(t, "output", lf.Name)
	assert.Equal(t, " STRING", lf.ArgumentName)
}

func TestNewLongFlagStringOptional(t *testing.T) {
	var v string = "default-value"
	lf := NewLongFlagStringOptional(NewValueString(&v), "config", "Config path.")

	assert.Equal(t, "config", lf.Name)
	assert.Equal(t, "[=STRING]", lf.ArgumentName)
	assert.Equal(t, "default-value", lf.DefaultValue)
}

func TestNewLongFlagStringSliceRequired(t *testing.T) {
	var v []string
	lf := NewLongFlagStringSliceRequired(NewValueStringSlice(&v), "header", "Set header.")

	assert.Equal(t, "header", lf.Name)
	assert.Equal(t, " STRING", lf.ArgumentName)
}

func TestNewLongFlagUintRequired(t *testing.T) {
	var v uint
	lf := NewLongFlagUintRequired(NewValueUint(&v), "users", "Set users.")

	assert.Equal(t, "users", lf.Name)
	assert.Equal(t, " UINT", lf.ArgumentName)
}

func TestNewLongFlagUint8Required(t *testing.T) {
	var v uint8
	lf := NewLongFlagUint8Required(NewValueUint8(&v), "queue", "Set queue.")

	assert.Equal(t, "queue", lf.Name)
	assert.Equal(t, " UINT8", lf.ArgumentName)
}

func TestNewLongFlagUint16Required(t *testing.T) {
	var v uint16
	lf := NewLongFlagUint16Required(NewValueUint16(&v), "max", "Set max.")

	assert.Equal(t, "max", lf.Name)
	assert.Equal(t, " UINT16", lf.ArgumentName)
}

func TestNewLongFlagUint32Required(t *testing.T) {
	var v uint32
	lf := NewLongFlagUint32Required(NewValueUint32(&v), "cache", "Set cache.")

	assert.Equal(t, "cache", lf.Name)
	assert.Equal(t, " UINT32", lf.ArgumentName)
}

func TestNewLongFlagUint64Required(t *testing.T) {
	var v uint64
	lf := NewLongFlagUint64Required(NewValueUInt64(&v), "limit", "Set limit.")

	assert.Equal(t, "limit", lf.Name)
	assert.Equal(t, " UINT64", lf.ArgumentName)
}

func TestLongFlagMakeOptionPanicsOnEmptyPrefix(t *testing.T) {
	var v bool
	lf := NewLongFlagBool(NewValueBool(&v), "verbose", "Verbose.")
	lf.Prefix = ""

	assert.Panics(t, func() {
		lf.MakeOption(lf)
	})
}

func TestLongFlagMakeOptionPanicsOnEmptyName(t *testing.T) {
	var v bool
	lf := NewLongFlagBool(NewValueBool(&v), "verbose", "Verbose.")
	lf.Name = ""

	assert.Panics(t, func() {
		lf.MakeOption(lf)
	})
}
