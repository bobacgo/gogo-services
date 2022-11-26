package pkg

type Slices[T comparable] struct{}

func (s Slices[T]) Distinct(arr []T) (res []T) {
	if len(arr) == 0 {
		return res
	}
	tempM := make(map[T]struct{}, 1)
	t := struct{}{}
	for _, v := range arr {
		tempV := v
		if _, ok := tempM[tempV]; !ok {
			tempM[tempV] = t
			res = append(res, tempV)
		}
	}
	return res
}