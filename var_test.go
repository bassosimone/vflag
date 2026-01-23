// SPDX-License-Identifier: GPL-3.0-or-later

package vflag

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlagSetVarAutoHelp(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		fs.AutoHelp('h', "help", "Print help and exit.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify short flag
		short := fs.ShortFlags[0]
		assert.Equal(t, byte('h'), short.Name)
		assert.Equal(t, "-", short.Prefix)
		assert.Equal(t, []string{"Print help and exit."}, short.Description)
		_, ok := short.Value.(ValueAutoHelp)
		assert.True(t, ok)

		// Verify long flag
		long := fs.LongFlags[0]
		assert.Equal(t, "help", long.Name)
		assert.Equal(t, "--", long.Prefix)
		assert.Equal(t, []string{"Print help and exit."}, long.Description)
		_, ok = long.Value.(ValueAutoHelp)
		assert.True(t, ok)
	})

	t.Run("short only", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		fs.AutoHelp('h', "", "Print help and exit.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 0)
	})

	t.Run("long only", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		fs.AutoHelp(0, "help", "Print help and exit.")

		require.Len(t, fs.ShortFlags, 0)
		require.Len(t, fs.LongFlags, 1)
	})
}

func TestFlagSetVarBool(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value bool
		fs.BoolVar(&value, 'v', "verbose", "Enable verbose output.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify short flag
		short := fs.ShortFlags[0]
		assert.Equal(t, byte('v'), short.Name)
		assert.Equal(t, "-", short.Prefix)
		assert.Equal(t, "", short.ArgumentName)

		// Verify long flag
		long := fs.LongFlags[0]
		assert.Equal(t, "verbose", long.Name)
		assert.Equal(t, "--", long.Prefix)
		assert.Equal(t, "[=true|false]", long.ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, short.Value.Set("true"))
		assert.Equal(t, "true", long.Value.String())
		assert.True(t, value)
	})

	t.Run("short only", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value bool
		fs.BoolVar(&value, 'v', "", "Enable verbose output.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 0)
	})

	t.Run("long only", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value bool
		fs.BoolVar(&value, 0, "verbose", "Enable verbose output.")

		require.Len(t, fs.ShortFlags, 0)
		require.Len(t, fs.LongFlags, 1)
	})
}

func TestFlagSetVarDuration(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value time.Duration
		fs.DurationVar(&value, 't', "timeout", "Set timeout.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " DURATION", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " DURATION", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("5s"))
		assert.Equal(t, "5s", fs.LongFlags[0].Value.String())
		assert.Equal(t, 5*time.Second, value)
	})
}

func TestFlagSetVarFloat64(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value float64
		fs.Float64Var(&value, 'r', "ratio", "Set ratio.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " FLOAT64", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " FLOAT64", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("3.14"))
		assert.Equal(t, "3.14", fs.LongFlags[0].Value.String())
		assert.Equal(t, 3.14, value)
	})
}

func TestFlagSetVarInt(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value int
		fs.IntVar(&value, 'n', "count", "Set count.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " INT", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " INT", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("42"))
		assert.Equal(t, "42", fs.LongFlags[0].Value.String())
		assert.Equal(t, 42, value)
	})
}

func TestFlagSetVarInt8(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value int8
		fs.Int8Var(&value, 'b', "batch", "Set batch size.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " INT8", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " INT8", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("10"))
		assert.Equal(t, "10", fs.LongFlags[0].Value.String())
		assert.Equal(t, int8(10), value)
	})
}

func TestFlagSetVarInt16(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value int16
		fs.Int16Var(&value, 'p', "port", "Set port.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " INT16", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " INT16", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("8080"))
		assert.Equal(t, "8080", fs.LongFlags[0].Value.String())
		assert.Equal(t, int16(8080), value)
	})
}

func TestFlagSetVarInt32(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value int32
		fs.Int32Var(&value, 'i', "index", "Set index.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " INT32", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " INT32", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("100"))
		assert.Equal(t, "100", fs.LongFlags[0].Value.String())
		assert.Equal(t, int32(100), value)
	})
}

func TestFlagSetVarInt64(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value int64
		fs.Int64Var(&value, 's', "size", "Set size.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " INT64", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " INT64", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("1000"))
		assert.Equal(t, "1000", fs.LongFlags[0].Value.String())
		assert.Equal(t, int64(1000), value)
	})
}

func TestFlagSetVarString(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value string
		fs.StringVar(&value, 'o', "output", "Set output file.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " STRING", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " STRING", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("output.txt"))
		assert.Equal(t, "output.txt", fs.LongFlags[0].Value.String())
		assert.Equal(t, "output.txt", value)
	})
}

func TestFlagSetVarStringSlice(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value []string
		fs.StringSliceVar(&value, 'H', "header", "Set header.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " STRING", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " STRING", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("X-Custom: value"))
		assert.Equal(t, []string{"X-Custom: value"}, value)
	})
}

func TestFlagSetVarUint(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value uint
		fs.UintVar(&value, 'u', "users", "Set users.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " UINT", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " UINT", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("50"))
		assert.Equal(t, "50", fs.LongFlags[0].Value.String())
		assert.Equal(t, uint(50), value)
	})
}

func TestFlagSetVarUint8(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value uint8
		fs.Uint8Var(&value, 'q', "queue", "Set queue size.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " UINT8", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " UINT8", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("128"))
		assert.Equal(t, "128", fs.LongFlags[0].Value.String())
		assert.Equal(t, uint8(128), value)
	})
}

func TestFlagSetVarUint16(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value uint16
		fs.Uint16Var(&value, 'm', "max", "Set max.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " UINT16", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " UINT16", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("1000"))
		assert.Equal(t, "1000", fs.LongFlags[0].Value.String())
		assert.Equal(t, uint16(1000), value)
	})
}

func TestFlagSetVarUint32(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value uint32
		fs.Uint32Var(&value, 'c', "cache", "Set cache size.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " UINT32", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " UINT32", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("4096"))
		assert.Equal(t, "4096", fs.LongFlags[0].Value.String())
		assert.Equal(t, uint32(4096), value)
	})
}

func TestFlagSetVarUint64(t *testing.T) {
	t.Run("both short and long", func(t *testing.T) {
		fs := NewFlagSet("prog", ContinueOnError)
		var value uint64
		fs.Uint64Var(&value, 'l', "limit", "Set limit.")

		require.Len(t, fs.ShortFlags, 1)
		require.Len(t, fs.LongFlags, 1)

		// Verify argument names
		assert.Equal(t, " UINT64", fs.ShortFlags[0].ArgumentName)
		assert.Equal(t, " UINT64", fs.LongFlags[0].ArgumentName)

		// Verify shared value by setting one and checking the other
		require.NoError(t, fs.ShortFlags[0].Value.Set("999999"))
		assert.Equal(t, "999999", fs.LongFlags[0].Value.String())
		assert.Equal(t, uint64(999999), value)
	})
}
