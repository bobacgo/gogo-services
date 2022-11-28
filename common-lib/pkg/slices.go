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
	tempM := make(map[T]struct{}, 1)
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

func (s *dataStream[T]) Each(fn func(i int, o T)) {
	for i, v := range s.v {
		fn(i, v)
	}
}

func (s *dataStream[T]) List() []T {
	return s.v
}
