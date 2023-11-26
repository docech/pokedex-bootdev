package api

type Resource[T any] interface {
	Next() (Resource[T], error)
	Previous() (Resource[T], error)
	Data() T
}
