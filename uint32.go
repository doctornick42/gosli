package main

import "errors"

type Uint32Slice []uint32

func (r Uint32Slice) FirstOrDefault(f func(uint32) bool) uint32 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal uint32
	return defVal
}
func (r Uint32Slice) First(f func(uint32) bool) (uint32, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal uint32
	return defVal, errors.New("Not found")
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
	_, err := r.First(f)
	return err == nil
}
func (r Uint32Slice) Contains(el uint32) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Uint32Slice) GetUnion(sl2 []uint32) ([]uint32, error) {
	result := make([]uint32, 0)
	for _, sl1El := range r {
		for _, sl2El := range sl2 {
			areEqual := sl1El == sl2El
			if areEqual {
				result = append(result, sl1El)
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
			areEqual := sl1El == sl2El
			if areEqual {
				found = true
				continue
			}
		}
		if !found {
			result = append(result, sl1El)
		}
	}
	return result, nil
}
