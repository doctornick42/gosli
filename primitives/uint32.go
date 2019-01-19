package primitives

import (
	"errors"
)

type Uint32Slice []uint32

func (r Uint32Slice) FirstOrDefault(f func(uint32) bool) *uint32 {
	for _, slEl := range r {
		if f(slEl) {
			return &slEl
		}
	}
	return nil
}
func (r Uint32Slice) First(f func(uint32) bool) (uint32, error) {
	first := r.FirstOrDefault(f)
	if first == nil {
		return 0, errors.New("Not found")
	}
	return *first, nil
}
func (r Uint32Slice) Where(f func(uint32) bool) []uint32 {
	res := make([]uint32, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Uint32Slice) Select(f func(uint32) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Uint32Slice) Page(number int64, perPage int64) ([]uint32, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []uint32{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Uint32Slice) Any(f func(uint32) bool) bool {
	first := r.FirstOrDefault(f)
	return first != nil
}

func (r Uint32Slice) Contains(el uint32) (bool, error) {
	return r.Any(func(i uint32) bool {
		return el == i
	}), nil
}

func (r Uint32Slice) processSliceOperation(sl2 Uint32Slice, f func([]uint32, []uint32) ([]uint32, error)) ([]uint32, error) {
	untypedRes, err := f(r, sl2)
	if err != nil {
		return nil, err
	}
	res := make([]uint32, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i]
	}
	return res, nil
}
func (r Uint32Slice) GetUnion(sl2 []uint32) ([]uint32, error) {
	result := make([]uint32, 0)

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
func (r Uint32Slice) InFirstOnly(sl2 []uint32) ([]uint32, error) {
	result := make([]uint32, 0)

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
