package kleverr

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStackError(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		gerr := New("abc")
		serr := GetStack(gerr)
		require.NotNil(t, serr)

		fmt.Println(serr.Print())
	})

	t.Run("nowrap", func(t *testing.T) {
		gerr := Ret(nil)
		require.NoError(t, gerr)
	})

	t.Run("wrap", func(t *testing.T) {
		err := errors.New("cde")
		gerr := Ret(err)
		require.ErrorIs(t, gerr, err)

		serr := GetStack(gerr)
		require.NotNil(t, serr)

		fmt.Println(serr.Print())
	})

	t.Run("mwrap", func(t *testing.T) {
		err := errors.New("efg")
		var makeErr = func(err error) error {
			return Ret(err)
		}
		gerr := Ret(makeErr(err))
		require.ErrorIs(t, gerr, err)

		serr := GetStack(gerr)
		require.NotNil(t, serr)

		fmt.Println(serr.Print())
	})
}
