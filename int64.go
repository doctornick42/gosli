package main

import "errors"

type Int64Slice []int64

func (r Int64Slice) FirstOrDefault(f func(int64) bool) int64 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal int64
	return defVal
}
func (r Int64Slice) First(f func(int64) bool) (int64, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal int64
	return defVal, errors.New("Not found")
}
func (r Int64Slice) Where(f func(int64) bool) []int64 {
	res := make([]int64, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Int64Slice) Select(f func(int64) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Int64Slice) Page(number int64, perPage int64) ([]int64, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int64{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Int64Slice) Any(f func(int64) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Int64Slice) Contains(el int64) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Int64Slice) GetUnion(sl2 []int64) ([]int64, error) {
	result := make([]int64, 0)
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
func (r Int64Slice) InFirstOnly(sl2 []int64) ([]int64, error) {
	result := make([]int64, 0)
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
