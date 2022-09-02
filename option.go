package option

type (
	Option[T any] interface {
		OkOr(err error) Result[T]
		Or(o Option[T]) Option[T]
		UnwrapOr(v T) T
		IsNone() bool
		IsSome() bool
		Filter(filter func(v T) bool) Option[T]
	}
	Some[T any] struct {
		v T
	}
	None[T any] struct {
	}
)

func NewSome[T any](v T) Option[T] {
	return &Some[T]{v: v}
}

func NewNone[T any]() Option[T] {
	return &None[T]{}
}

// -----------------

func (n None[T]) OkOr(err error) Result[T] {
	return NewErr[T](err)
}

func (n None[T]) IsNone() bool {
	return true
}

func (n None[T]) IsSome() bool {
	return false
}

func (n None[T]) UnwrapOr(v T) T {
	return v
}

func (n None[T]) Or(o Option[T]) Option[T] {
	return o
}

func (n None[T]) Filter(_ func(v T) bool) Option[T] {
	return n
}

// -----------------

func (o *Some[T]) Or(_ Option[T]) Option[T] {
	return o
}

func (o *Some[T]) IsNone() bool {
	return false
}

func (o *Some[T]) IsSome() bool {
	return true
}

func (o *Some[T]) Filter(filter func(v T) bool) Option[T] {
	if filter(o.v) {
		return o
	}

	return NewNone[T]()
}

func (o *Some[T]) UnwrapOr(_ T) T {
	return o.v
}

func (o *Some[T]) OkOr(_ error) Result[T] {
	return NewOk[T](o.v)
}
