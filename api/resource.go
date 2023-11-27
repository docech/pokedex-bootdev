package api

type Resource[T any] interface {
	Data() T
}

type DetailResource[ID any, T any] interface {
	Resource[T]
	Detail(id ID) error
}

type ListResource[T any] interface {
	Resource[[]T]
	Next() error
	Previous() error
}
