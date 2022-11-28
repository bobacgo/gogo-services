package mapset

// Set no-safe、unordered
type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable]() Set[T] {
	return Set[T]{
		m: make(map[T]struct{}),
	}
}

func Of[T comparable](obj ...T) Set[T] {
	s := New[T]()
	for _, o := range obj {
		s.Add(o)
	}
	return s
}

func (s Set[T]) Add(obj T) {
	s.m[obj] = struct{}{}
}

func (s Set[T]) Value() []T {
	es := make([]T, len(s.m))
	for k := range s.m {
		es = append(es, k)
	}
	return es
}

func (s Set[T]) Each(fn func(element T)) {
	for k := range s.m {
		fn(k)
	}
}