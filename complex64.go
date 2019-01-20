package main

import "errors"

type Complex64Slice []complex64

func (r Complex64Slice) FirstOrDefault(f func(complex64) bool) complex64 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal complex64
	return defVal
}
func (r Complex64Slice) First(f func(complex64) bool) (complex64, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal complex64
	return defVal, errors.New("Not found")
}
func (r Complex64Slice) Where(f func(complex64) bool) []complex64 {
	res := make([]complex64, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Complex64Slice) Select(f func(complex64) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Complex64Slice) Page(number int64, perPage int64) ([]complex64, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []complex64{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Complex64Slice) Any(f func(complex64) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Complex64Slice) Contains(el complex64) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Complex64Slice) GetUnion(sl2 []complex64) ([]complex64, error) {
	result := make([]complex64, 0)
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
func (r Complex64Slice) InFirstOnly(sl2 []complex64) ([]complex64, error) {
	result := make([]complex64, 0)
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
