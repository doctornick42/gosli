package main

import "errors"

type UintSlice []uint

func (r UintSlice) FirstOrDefault(f func(uint) bool) uint {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal uint
	return defVal
}
func (r UintSlice) First(f func(uint) bool) (uint, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal uint
	return defVal, errors.New("Not found")
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
func (r UintSlice) Page(number int64, perPage int64) ([]uint, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []uint{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r UintSlice) Any(f func(uint) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r UintSlice) Contains(el uint) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r UintSlice) GetUnion(sl2 []uint) ([]uint, error) {
	result := make([]uint, 0)
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
func (r UintSlice) InFirstOnly(sl2 []uint) ([]uint, error) {
	result := make([]uint, 0)
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
