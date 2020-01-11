package main

import "errors"

type BoolSlice []bool

func (r BoolSlice) FirstOrDefault(f func(bool) bool) bool {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal bool
	return defVal
}
func (r BoolSlice) First(f func(bool) bool) (bool, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal bool
	return defVal, errors.New("Not found")
}
func (r BoolSlice) Where(f func(bool) bool) []bool {
	res := make([]bool, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r BoolSlice) Select(f func(bool) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r BoolSlice) Page(number int64, perPage int64) ([]bool, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []bool{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r BoolSlice) Any(f func(bool) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r BoolSlice) Contains(el bool) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r BoolSlice) GetUnion(sl2 []bool) ([]bool, error) {
	result := make([]bool, 0)
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
func (r BoolSlice) InFirstOnly(sl2 []bool) ([]bool, error) {
	result := make([]bool, 0)
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
