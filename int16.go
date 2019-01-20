package main

import "errors"

type Int16Slice []int16

func (r Int16Slice) FirstOrDefault(f func(int16) bool) int16 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal int16
	return defVal
}
func (r Int16Slice) First(f func(int16) bool) (int16, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal int16
	return defVal, errors.New("Not found")
}
func (r Int16Slice) Where(f func(int16) bool) []int16 {
	res := make([]int16, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Int16Slice) Select(f func(int16) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Int16Slice) Page(number int64, perPage int64) ([]int16, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int16{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Int16Slice) Any(f func(int16) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Int16Slice) Contains(el int16) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Int16Slice) GetUnion(sl2 []int16) ([]int16, error) {
	result := make([]int16, 0)
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
func (r Int16Slice) InFirstOnly(sl2 []int16) ([]int16, error) {
	result := make([]int16, 0)
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
