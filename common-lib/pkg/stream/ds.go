// Package stream is data stream ETL
package stream

// New
// Of
// Connect

// Distinct
// Filter
// Reverse
// Limit
// Skip

// Each
// List

var ptr = struct{}{}

type Stream[T comparable] chan T

// === in ======

// New slice to stream
func New[T comparable](arr []T) Stream[T] {
	ch := make(chan T, len(arr))
	go func() {
		for _, v := range arr {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// Of element to stream
func Of[T comparable](obj ...T) Stream[T] {
	return New(obj)
}

// Connect Multiple slice splicing
func Connect[T comparable](arr []T, arrs ...[]T) Stream[T] {
	for _, sli := range arrs {
		arr = append(arr, sli...)
	}
	return New(arr)
}

// operator

// Distinct 去重
func (s Stream[T]) Distinct() Stream[T] {
	ch := make(chan T)
	go func() {
		m := make(map[T]struct{})
		for v := range s {
			if _, ok := m[v]; !ok {
				ch <- v
				m[v] = ptr
			}
		}
		close(ch)
	}()
	return ch
}

// Filter 过滤
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

// Reverse 数据反转
func (s Stream[T]) Reverse() Stream[T] {
	var arr []T
	for v := range s {
		arr = append(arr, v)
	}

	ch := make(chan T)
	go func() {
		for i := len(arr) - 1; i >= 0; i-- {
			ch <- arr[i]
		}
		close(ch)
	}()
	return ch
}

func (s Stream[T]) Limit(n uint) Stream[T] {
	ch := make(chan T)
	go func() {
		count, isClose := 0, false
		for v := range s {
			if count < int(n) {
				ch <- v
			}
			if count == int(n)-1 {
				isClose = true
				close(ch)
			}
			count++
		}
		if !isClose { // n > len(data)
			close(ch)
		}
	}()
	return ch
}

func (s Stream[T]) Skip(n uint) Stream[T] {
	ch := make(chan T)
	go func() {
		count := 0
		for v := range s {
			if count >= int(n) {
				ch <- v
			}
			count++
		}
		close(ch)
	}()
	return ch
}

// === out ======

// Each for every element
func (s Stream[T]) Each(fn func(i int, o T)) {
	idx := 0
	for v := range s {
		fn(idx, v)
		idx++
	}
}

// List stream to slice
func (s Stream[T]) List() (arr []T) {
	for v := range s {
		arr = append(arr, v)
	}
	return arr
}
