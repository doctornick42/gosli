package main

import "errors"

type IntSlice []int

func (r IntSlice) FirstOrDefault(f func(int) bool) int {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal int
	return defVal
}
func (r IntSlice) First(f func(int) bool) (int, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal int
	return defVal, errors.New("Not found")
}
func (r IntSlice) Where(f func(int) bool) []int {
	res := make([]int, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r IntSlice) Select(f func(int) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r IntSlice) Page(number int64, perPage int64) ([]int, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r IntSlice) Any(f func(int) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r IntSlice) Contains(el int) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r IntSlice) GetUnion(sl2 []int) ([]int, error) {
	result := make([]int, 0)
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
func (r IntSlice) InFirstOnly(sl2 []int) ([]int, error) {
	result := make([]int, 0)
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
