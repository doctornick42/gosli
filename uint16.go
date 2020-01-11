package main

import "errors"

type Uint16Slice []uint16

func (r Uint16Slice) FirstOrDefault(f func(uint16) bool) uint16 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal uint16
	return defVal
}
func (r Uint16Slice) First(f func(uint16) bool) (uint16, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal uint16
	return defVal, errors.New("Not found")
}
func (r Uint16Slice) Where(f func(uint16) bool) []uint16 {
	res := make([]uint16, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Uint16Slice) Select(f func(uint16) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Uint16Slice) Page(number int64, perPage int64) ([]uint16, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []uint16{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Uint16Slice) Any(f func(uint16) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Uint16Slice) Contains(el uint16) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Uint16Slice) GetUnion(sl2 []uint16) ([]uint16, error) {
	result := make([]uint16, 0)
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
func (r Uint16Slice) InFirstOnly(sl2 []uint16) ([]uint16, error) {
	result := make([]uint16, 0)
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
