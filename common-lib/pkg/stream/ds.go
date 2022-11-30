package stream

type Stream[T comparable] chan T

func Of[T comparable](obj ...T) Stream[T] {
	ch := make(chan T, len(obj))
	go func() {
		for _, v := range obj {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func (s Stream[T]) Distinct() Stream[T] {
	ch := make(chan T)
	go func() {
		ptr := struct{}{}
		m := make(map[T]struct{})
		for v := range s {
			if _, ok := m[v]; !ok {
				ch <- v
				m[v] = ptr
			}
		}
		close(ch)
	}()
	return s
}
func (s Stream[T]) Filter(fn func(o T) bool) Stream[T] {
	ch := make(chan T)
	go func() {
		for v := range s {
			if fn(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

func (s Stream[T]) Each(fn func(i int, o T)) {
	idx := 0
	for v := range s {
		fn(idx, v)
		idx++
	}
}

func (s Stream[T]) List() (arr []T) {
	for v := range s {
		arr = append(arr, v)
	}
	return arr
}
