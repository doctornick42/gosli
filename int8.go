package main

import "errors"

type Int8Slice []int8

func (r Int8Slice) FirstOrDefault(f func(int8) bool) int8 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal int8
	return defVal
}
func (r Int8Slice) First(f func(int8) bool) (int8, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal int8
	return defVal, errors.New("Not found")
}
func (r Int8Slice) Where(f func(int8) bool) []int8 {
	res := make([]int8, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Int8Slice) Select(f func(int8) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Int8Slice) Page(number int64, perPage int64) ([]int8, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int8{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Int8Slice) Any(f func(int8) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Int8Slice) Contains(el int8) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Int8Slice) GetUnion(sl2 []int8) ([]int8, error) {
	result := make([]int8, 0)
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
func (r Int8Slice) InFirstOnly(sl2 []int8) ([]int8, error) {
	result := make([]int8, 0)
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
