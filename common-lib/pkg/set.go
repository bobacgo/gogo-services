package pkg

var EmptyStruct = struct{}{}

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](size uint) *Set[T] {
	return &Set[T]{
		make(map[T]struct{}),
	}
}