package api

type Resource[T any] interface {
	Data() T
}

type ListResource[T any] interface {
	Resource[[]T]
	Next() (ListResource[T], error)
	Previous() (ListResource[T], error)
}
