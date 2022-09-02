package option

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk_IsOk(t *testing.T) {
	assert.True(t, NewOk[int](1).IsOk())
}

func TestOk_Ok(t *testing.T) {
	assert.EqualValues(t, &Ok[int]{
		v: 1,
	}, NewOk(1).Ok())

	assert.EqualValues(t, NewNone[int](), NewErr[int](errors.New("1")).Ok())
}

func TestOk_Or(t *testing.T) {
	assert.Equal(t, 1, NewOk(1).Or(NewOk(2)).(*Ok[int]).v)
}
func TestOk_IsErr(t *testing.T) {
	assert.False(t, NewOk[int](1).IsErr())
}

func TestOk_UnwrapErr(t *testing.T) {
	err := NewOk[int](1).UnwrapErr(func(err error) error {
		return nil
	})
	assert.NoError(t, err)
}

func TestOk_UnwrapOr(t *testing.T) {
	assert.Equal(t, 1, NewOk[int](1).UnwrapOr(100))
}

func TestErr_IsErr(t *testing.T) {
	assert.True(t, NewErr[int](errors.New("1")).IsErr())
}

func TestErr_UnwrapErr(t *testing.T) {
	err := errors.New("err")
	assert.ErrorIs(t, NewErr[int](err).UnwrapErr(func(err error) error {
		return fmt.Errorf("any %w", err)
	}), err)

	assert.Equal(t, err, NewErr[int](err).UnwrapErr(func(err error) error {
		return err
	}))

}

func TestErr_UnwrapOr(t *testing.T) {
	assert.Equal(t, 1, NewErr[int](errors.New("1")).UnwrapOr(1))
}

func TestErr_IsOk(t *testing.T) {
	assert.False(t, NewErr[int](errors.New("1")).IsOk())
}

func TestErr_Ok(t *testing.T) {
	assert.Equal(t, NewNone[int](), NewErr[int](errors.New("1")).Ok())
}

func TestErr_Or(t *testing.T) {
	assert.Equal(t, 2, NewErr[int](errors.New("")).Or(NewOk(2)).(*Ok[int]).v)
}
