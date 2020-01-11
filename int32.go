package main

import "errors"

type Int32Slice []int32

func (r Int32Slice) FirstOrDefault(f func(int32) bool) int32 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal int32
	return defVal
}
func (r Int32Slice) First(f func(int32) bool) (int32, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal int32
	return defVal, errors.New("Not found")
}
func (r Int32Slice) Where(f func(int32) bool) []int32 {
	res := make([]int32, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Int32Slice) Select(f func(int32) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Int32Slice) Page(number int64, perPage int64) ([]int32, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []int32{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Int32Slice) Any(f func(int32) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Int32Slice) Contains(el int32) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Int32Slice) GetUnion(sl2 []int32) ([]int32, error) {
	result := make([]int32, 0)
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
func (r Int32Slice) InFirstOnly(sl2 []int32) ([]int32, error) {
	result := make([]int32, 0)
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
