package main

import "errors"

type Uint8Slice []uint8

func (r Uint8Slice) FirstOrDefault(f func(uint8) bool) uint8 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal uint8
	return defVal
}
func (r Uint8Slice) First(f func(uint8) bool) (uint8, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal uint8
	return defVal, errors.New("Not found")
}
func (r Uint8Slice) Where(f func(uint8) bool) []uint8 {
	res := make([]uint8, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Uint8Slice) Select(f func(uint8) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Uint8Slice) Page(number int64, perPage int64) ([]uint8, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []uint8{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Uint8Slice) Any(f func(uint8) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Uint8Slice) Contains(el uint8) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Uint8Slice) GetUnion(sl2 []uint8) ([]uint8, error) {
	result := make([]uint8, 0)
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
func (r Uint8Slice) InFirstOnly(sl2 []uint8) ([]uint8, error) {
	result := make([]uint8, 0)
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
