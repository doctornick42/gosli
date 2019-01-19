package primitives

import (
	"errors"
)

type Int32Slice []int32

func (r Int32Slice) FirstOrDefault(f func(int32) bool) *int32 {
	for _, slEl := range r {
		if f(slEl) {
			return &slEl
		}
	}
	return nil
}
func (r Int32Slice) First(f func(int32) bool) (int32, error) {
	first := r.FirstOrDefault(f)
	if first == nil {
		return 0, errors.New("Not found")
	}
	return *first, nil
}
func (r Int32Slice) Where(f func(int32) bool) []int32 {
	res := make([]int32, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Int32Slice) Select(f func(int32) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Int32Slice) Page(number int64, perPage int64) ([]int32, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int32{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Int32Slice) Any(f func(int32) bool) bool {
	first := r.FirstOrDefault(f)
	return first != nil
}

func (r Int32Slice) Contains(el int32) (bool, error) {
	return r.Any(func(i int32) bool {
		return el == i
	}), nil
}

func (r Int32Slice) processSliceOperation(sl2 Int32Slice, f func([]int32, []int32) ([]int32, error)) ([]int32, error) {
	untypedRes, err := f(r, sl2)
	if err != nil {
		return nil, err
	}
	res := make([]int32, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i]
	}
	return res, nil
}
func (r Int32Slice) GetUnion(sl2 []int32) ([]int32, error) {
	result := make([]int32, 0)

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
func (r Int32Slice) InFirstOnly(sl2 []int32) ([]int32, error) {
	result := make([]int32, 0)

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
