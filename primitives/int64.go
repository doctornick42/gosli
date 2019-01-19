package primitives

import (
	"errors"
)

type Int64Slice []int64

func (r Int64Slice) FirstOrDefault(f func(int64) bool) *int64 {
	for _, slEl := range r {
		if f(slEl) {
			return &slEl
		}
	}
	return nil
}
func (r Int64Slice) First(f func(int64) bool) (int64, error) {
	first := r.FirstOrDefault(f)
	if first == nil {
		return 0, errors.New("Not found")
	}
	return *first, nil
}
func (r Int64Slice) Where(f func(int64) bool) []int64 {
	res := make([]int64, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Int64Slice) Select(f func(int64) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Int64Slice) Page(number int64, perPage int64) ([]int64, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int64{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Int64Slice) Any(f func(int64) bool) bool {
	first := r.FirstOrDefault(f)
	return first != nil
}

func (r Int64Slice) Contains(el int64) (bool, error) {
	return r.Any(func(i int64) bool {
		return el == i
	}), nil
}

func (r Int64Slice) processSliceOperation(sl2 Int64Slice, f func([]int64, []int64) ([]int64, error)) ([]int64, error) {
	untypedRes, err := f(r, sl2)
	if err != nil {
		return nil, err
	}
	res := make([]int64, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i]
	}
	return res, nil
}
func (r Int64Slice) GetUnion(sl2 []int64) ([]int64, error) {
	result := make([]int64, 0)

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
func (r Int64Slice) InFirstOnly(sl2 []int64) ([]int64, error) {
	result := make([]int64, 0)

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
