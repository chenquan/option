package option

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption_IsNone(t *testing.T) {
	o := NewNone[int64]()
	assert.True(t, o.IsNone())
}

func TestOption_IsSome(t *testing.T) {
	some := NewSome(2.1)
	assert.True(t, some.IsSome())
}

func TestOption_Filter(t *testing.T) {
	some := NewSome(2.1)
	o := some.Filter(func(v float64) bool {
		return v > 1
	})
	assert.True(t, o.IsSome())

	o = o.Filter(func(v float64) bool {
		return v < 1
	})
	assert.False(t, o.IsSome())
	assert.True(t, o.IsNone())

	o = o.Filter(func(v float64) bool {
		return v < 1
	})
	assert.False(t, o.IsSome())
	assert.True(t, o.IsNone())
}

func TestOption_Or(t *testing.T) {
	assert.False(t, NewNone[int64]().Or(NewSome[int64](1)).IsNone())
	assert.Equal(t, int64(1), NewSome[int64](1).Or(NewSome[int64](2)).UnwrapOr(22))
}

func TestOption_UnwrapOr(t *testing.T) {
	assert.Equal(t, int64(1), NewNone[int64]().UnwrapOr(int64(1)))
	assert.Equal(t, int64(1), NewSome[int64](1).UnwrapOr(int64(3)))
}

func TestOption_OkOr(t *testing.T) {
	or := NewNone[int]().OkOr(errors.New("err"))
	assert.Error(t, or.(*Err[int]).err)

	or = NewSome[int](1).OkOr(errors.New("err"))
	assert.Equal(t, 1, or.(*Ok[int]).v)
}
