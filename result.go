package option

type (
	Result[T any] interface {
		Or(result Result[T]) Result[T]
		Ok() Option[T]
		IsOk() bool
		IsErr() bool
		UnwrapErr(f func(err error) error) error
		UnwrapOr(v T) T
	}
	Err[T any] struct {
		err error
	}
	Ok[T any] struct {
		v T
	}
)

func NewErr[T any](err error) Result[T] {
	return &Err[T]{err: err}
}

func (e *Err[T]) Ok() Option[T] {
	return NewNone[T]()
}

func (e *Err[T]) IsOk() bool {
	return false
}

func (e *Err[T]) IsErr() bool {
	return true
}

func (e *Err[T]) UnwrapErr(f func(err error) error) error {
	return f(e.err)
}

func (e *Err[T]) UnwrapOr(v T) T {
	return v
}

func (e *Err[T]) Or(result Result[T]) Result[T] {
	return result
}

// -----------------

func NewOk[T any](v T) Result[T] {
	return &Ok[T]{v: v}
}

func (o *Ok[T]) Ok() Option[T] {
	return NewSome(o.v)
}

func (o *Ok[T]) IsOk() bool {
	return true
}

func (o *Ok[T]) IsErr() bool {
	return false
}

func (o *Ok[T]) UnwrapErr(_ func(err error) error) error {
	return nil
}

func (o *Ok[T]) UnwrapOr(_ T) T {
	return o.v
}

func (o *Ok[T]) Or(_ Result[T]) Result[T] {
	return o
}
