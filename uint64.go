package main

import "errors"

type Uint64Slice []uint64

func (r Uint64Slice) FirstOrDefault(f func(uint64) bool) uint64 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal uint64
	return defVal
}
func (r Uint64Slice) First(f func(uint64) bool) (uint64, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal uint64
	return defVal, errors.New("Not found")
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
	_, err := r.First(f)
	return err == nil
}
func (r Uint64Slice) Contains(el uint64) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Uint64Slice) GetUnion(sl2 []uint64) ([]uint64, error) {
	result := make([]uint64, 0)
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
func (r Uint64Slice) InFirstOnly(sl2 []uint64) ([]uint64, error) {
	result := make([]uint64, 0)
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
