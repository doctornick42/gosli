package primitives

import (
	"errors"
)

type UintSlice []uint

func (r UintSlice) FirstOrDefault(f func(uint) bool) *uint {
	for _, slEl := range r {
		if f(slEl) {
			return &slEl
		}
	}
	return nil
}
func (r UintSlice) First(f func(uint) bool) (uint, error) {
	first := r.FirstOrDefault(f)
	if first == nil {
		return 0, errors.New("Not found")
	}
	return *first, nil
}
func (r UintSlice) Where(f func(uint) bool) []uint {
	res := make([]uint, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r UintSlice) Select(f func(uint) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r UintSlice) Page(number uint64, perPage uint64) ([]uint, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > uint64(len(r)) {
		return []uint{}, nil
	}
	last := first + perPage
	if last > uint64(len(r)) {
		last = uint64(len(r))
	}
	return r[first:last], nil
}
func (r UintSlice) Any(f func(uint) bool) bool {
	first := r.FirstOrDefault(f)
	return first != nil
}

func (r UintSlice) Contains(el uint) (bool, error) {
	return r.Any(func(i uint) bool {
		return el == i
	}), nil
}

func (r UintSlice) processSliceOperation(sl2 UintSlice, f func([]uint, []uint) ([]uint, error)) ([]uint, error) {
	untypedRes, err := f(r, sl2)
	if err != nil {
		return nil, err
	}
	res := make([]uint, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i]
	}
	return res, nil
}
func (r UintSlice) GetUnion(sl2 []uint) ([]uint, error) {
	result := make([]uint, 0)

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
func (r UintSlice) InFirstOnly(sl2 []uint) ([]uint, error) {
	result := make([]uint, 0)

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
