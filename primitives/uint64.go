package primitives

import (
	"errors"
)

type Uint64Slice []uint64

func (r Uint64Slice) FirstOrDefault(f func(uint64) bool) *uint64 {
	for _, slEl := range r {
		if f(slEl) {
			return &slEl
		}
	}
	return nil
}
func (r Uint64Slice) First(f func(uint64) bool) (uint64, error) {
	first := r.FirstOrDefault(f)
	if first == nil {
		return 0, errors.New("Not found")
	}
	return *first, nil
}
func (r Uint64Slice) Where(f func(uint64) bool) []uint64 {
	res := make([]uint64, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Uint64Slice) Select(f func(uint64) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Uint64Slice) Page(number int64, perPage int64) ([]uint64, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []uint64{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Uint64Slice) Any(f func(uint64) bool) bool {
	first := r.FirstOrDefault(f)
	return first != nil
}

func (r Uint64Slice) Contains(el uint64) (bool, error) {
	return r.Any(func(i uint64) bool {
		return el == i
	}), nil
}

func (r Uint64Slice) processSliceOperation(sl2 Uint64Slice, f func([]uint64, []uint64) ([]uint64, error)) ([]uint64, error) {
	untypedRes, err := f(r, sl2)
	if err != nil {
		return nil, err
	}
	res := make([]uint64, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i]
	}
	return res, nil
}
func (r Uint64Slice) GetUnion(sl2 []uint64) ([]uint64, error) {
	result := make([]uint64, 0)

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
func (r Uint64Slice) InFirstOnly(sl2 []uint64) ([]uint64, error) {
	result := make([]uint64, 0)

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
