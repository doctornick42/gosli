package primitives

import (
	"errors"
)

type IntSlice []int

func (r IntSlice) FirstOrDefault(f func(int) bool) *int {
	for _, slEl := range r {
		if f(slEl) {
			return &slEl
		}
	}
	return nil
}
func (r IntSlice) First(f func(int) bool) (int, error) {
	first := r.FirstOrDefault(f)
	if first == nil {
		return 0, errors.New("Not found")
	}
	return *first, nil
}
func (r IntSlice) Where(f func(int) bool) []int {
	res := make([]int, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r IntSlice) Select(f func(int) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r IntSlice) Page(number int64, perPage int64) ([]int, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r IntSlice) Any(f func(int) bool) bool {
	first := r.FirstOrDefault(f)
	return first != nil
}

func (r IntSlice) Contains(el int) (bool, error) {
	return r.Any(func(i int) bool {
		return el == i
	}), nil
}

func (r IntSlice) processSliceOperation(sl2 IntSlice, f func([]int, []int) ([]int, error)) ([]int, error) {
	untypedRes, err := f(r, sl2)
	if err != nil {
		return nil, err
	}
	res := make([]int, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i]
	}
	return res, nil
}
func (r IntSlice) GetUnion(sl2 []int) ([]int, error) {
	result := make([]int, 0)

	for _, sl1El := range r {
		for _, sl2El := range sl2 {
			if sl1El == sl2El {
				result = append(result, sl1El)
				break
			}
		}
	}

	return result, nil
}
func (r IntSlice) InFirstOnly(sl2 []int) ([]int, error) {
	result := make([]int, 0)

	for _, sl1El := range r {
		found := false
		for _, sl2El := range sl2 {
			if sl1El == sl2El {
				found = true
				break
			}
		}

		if !found {
			result = append(result, sl1El)
		}
	}

	return result, nil
}
