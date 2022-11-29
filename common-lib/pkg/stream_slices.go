package pkg

type dataStream[T comparable] struct {
	v []T
}

func StreamSlice[T comparable](obj []T) *dataStream[T] {
	return &dataStream[T]{v: obj}
}

func (s *dataStream[T]) Distinct() *dataStream[T] {
	if len(s.v) == 0 {
		return s
	}
	var data []T
	tempM := make(map[T]struct{})
	t := struct{}{}
	for _, v := range s.v {
		if _, ok := tempM[v]; !ok {
			tempM[v] = t
			data = append(data, v)
		}
	}
	s.v = data
	return s
}

func (s *dataStream[T]) Filter(fn func(o T) bool) *dataStream[T] {
	var data []T
	for _, v := range s.v {
		if fn(v) {
			data = append(data, v)
		}
	}
	s.v = data
	return s
}

func (s *dataStream[T]) Limit(n uint) *dataStream[T] {
	v := s.v[:n]
	s.v = v
	return s
}

func (s *dataStream[T]) Skip(n uint) *dataStream[T] {
	v := s.v[n:]
	s.v = v
	return s
}

func (s *dataStream[T]) Peek(fn func(o *T)) *dataStream[T] {
	for i := 0; i < len(s.v); i++ {
		fn(&s.v[i])
	}
	return s
}

func (s *dataStream[T]) Each(fn func(i int, o T)) {
	for i, v := range s.v {
		fn(i, v)
	}
}

func (s *dataStream[T]) List() []T {
	return s.v
}