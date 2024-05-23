package codec

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytes_Encode(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		c := new(Bytes)

		var err error
		var encoded []byte
		expected := []byte("encode this")

		encoded, err = c.Encode(expected)
		require.Nil(t, err)
		require.EqualValues(t, expected, encoded)
	})

	t.Run("fail_string", func(t *testing.T) {
		c := new(Bytes)

		var err error
		expected := "encode this"

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_bool", func(t *testing.T) {
		c := new(Bytes)

		var err error
		expected := true

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_int", func(t *testing.T) {
		c := new(Bytes)

		var err error
		expected := 2048

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})
}

func TestBytes_Decode(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		c := new(Bytes)
		expected := []byte("decode this")

		decoded, err := c.Decode(expected)
		require.Nil(t, err)
		require.EqualValues(t, expected, decoded)
	})
}

func TestString_Encode(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		c := new(String)

		var err error
		var encoded []byte
		expected := "encode this"

		encoded, err = c.Encode(expected)
		require.Nil(t, err)
		require.EqualValues(t, expected, encoded)
	})

	t.Run("fail_bytes", func(t *testing.T) {
		c := new(String)

		var err error
		expected := []byte("encode this")

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_bool", func(t *testing.T) {
		c := new(String)

		var err error
		expected := true

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_int", func(t *testing.T) {
		c := new(String)

		var err error
		expected := 2048

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})
}

func TestString_Decode(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		c := new(String)
		expected := []byte("decode this")

		decoded, err := c.Decode(expected)
		require.Nil(t, err)
		require.EqualValues(t, expected, decoded)
	})
}

func TestInt64_Encode(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		c := new(Int64)

		var err error
		var encoded []byte
		var expected int64 = 2048

		encoded, err = c.Encode(expected)
		require.Nil(t, err)
		require.EqualValues(t, strconv.FormatInt(expected, 10), encoded)
	})

	t.Run("fail_int32", func(t *testing.T) {
		c := new(Int64)
		var expected int32 = 2048

		_, err := c.Encode(expected)
		require.NotNil(t, err)
	})
	t.Run("fail_int", func(t *testing.T) {
		c := new(Int64)
		var expected int = 2048

		_, err := c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_bytes", func(t *testing.T) {
		c := new(Int64)

		var err error
		expected := []byte("encode this")

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_string", func(t *testing.T) {
		c := new(Int64)

		var err error
		expected := 2048

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_bool", func(t *testing.T) {
		c := new(Int64)

		var err error
		expected := true

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})

	t.Run("fail_float32", func(t *testing.T) {
		c := new(Int64)

		var err error
		var expected float32 = 2048

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})
	t.Run("fail_float64", func(t *testing.T) {
		c := new(Int64)

		var err error
		var expected float64 = 2048

		_, err = c.Encode(expected)
		require.NotNil(t, err)
	})
}

func TestInt64_Decode(t *testing.T) {
	t.Run("succeed_int64", func(t *testing.T) {
		c := new(Int64)
		var expected int64 = 2048
		data := fmt.Sprintf("%d", expected)

		decoded, err := c.Decode([]byte(data))
		require.Nil(t, err)
		require.EqualValues(t, expected, decoded)
	})
	t.Run("succeed_int", func(t *testing.T) {
		c := new(Int64)
		var expected int = 2048
		data := fmt.Sprintf("%d", expected)

		decoded, err := c.Decode([]byte(data))
		require.Nil(t, err)
		require.EqualValues(t, expected, decoded)
	})
	t.Run("succeed_int32", func(t *testing.T) {
		c := new(Int64)
		var expected int32 = 2048
		data := fmt.Sprintf("%d", expected)

		decoded, err := c.Decode([]byte(data))
		require.Nil(t, err)
		require.EqualValues(t, expected, decoded)
	})

	t.Run("fail_float32", func(t *testing.T) {
		c := new(Int64)
		var expected float32 = 2048
		data := fmt.Sprintf("%f", expected)

		_, err := c.Decode([]byte(data))
		require.NotNil(t, err)
	})
	t.Run("fail_float64", func(t *testing.T) {
		c := new(Int64)
		var expected float64 = 2048
		data := fmt.Sprintf("%f", expected)

		_, err := c.Decode([]byte(data))
		require.NotNil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		c := new(Int64)
		var expected int64 = 2048
		data := fmt.Sprintf("decode %d", expected)

		_, err := c.Decode([]byte(data))
		require.NotNil(t, err)
	})
}
