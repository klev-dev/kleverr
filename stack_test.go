package kleverr

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStackError(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		gerr := New("abc")
		serr := Get(gerr)
		require.NotNil(t, serr)

		fmt.Println(serr.Print())
		require.Contains(t, serr.Print(), "stack_test")
	})

	t.Run("nowrap", func(t *testing.T) {
		gerr := Ret(nil)
		require.NoError(t, gerr)
	})

	t.Run("wrap", func(t *testing.T) {
		err := errors.New("cde")
		gerr := Ret(err)
		require.ErrorIs(t, gerr, err)

		serr := Get(gerr)
		require.NotNil(t, serr)

		fmt.Println(serr.Print())
		require.Contains(t, serr.Print(), "stack_test")
	})

	t.Run("mwrap", func(t *testing.T) {
		err := errors.New("efg")
		var makeErr = func(err error) error {
			return Ret(err)
		}
		gerr := Ret(makeErr(err))
		require.ErrorIs(t, gerr, err)

		serr := Get(gerr)
		require.NotNil(t, serr)

		fmt.Println(serr.Print())
		require.Contains(t, serr.Print(), "stack_test")
		require.True(t, strings.Contains(serr.Print(), "stack_test"))
	})
	t.Run("nwrap", func(t *testing.T) {
		err := errors.New("cde")
		str, gerr := Ret1[string](err)
		require.ErrorIs(t, gerr, err)
		require.Zero(t, str)

		serr := Get(gerr)
		require.NotNil(t, serr)

		fmt.Println(serr.Print())
		require.Contains(t, serr.Print(), "stack_test")
	})
}
