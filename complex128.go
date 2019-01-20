package main

import "errors"

type Complex128Slice []complex128

func (r Complex128Slice) FirstOrDefault(f func(complex128) bool) complex128 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal complex128
	return defVal
}
func (r Complex128Slice) First(f func(complex128) bool) (complex128, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal complex128
	return defVal, errors.New("Not found")
}
func (r Complex128Slice) Where(f func(complex128) bool) []complex128 {
	res := make([]complex128, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Complex128Slice) Select(f func(complex128) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Complex128Slice) Page(number int64, perPage int64) ([]complex128, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []complex128{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Complex128Slice) Any(f func(complex128) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Complex128Slice) Contains(el complex128) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Complex128Slice) GetUnion(sl2 []complex128) ([]complex128, error) {
	result := make([]complex128, 0)
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
func (r Complex128Slice) InFirstOnly(sl2 []complex128) ([]complex128, error) {
	result := make([]complex128, 0)
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
