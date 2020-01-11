package main

import "errors"

type UintptrSlice []uintptr

func (r UintptrSlice) FirstOrDefault(f func(uintptr) bool) uintptr {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal uintptr
	return defVal
}
func (r UintptrSlice) First(f func(uintptr) bool) (uintptr, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal uintptr
	return defVal, errors.New("Not found")
}
func (r UintptrSlice) Where(f func(uintptr) bool) []uintptr {
	res := make([]uintptr, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r UintptrSlice) Select(f func(uintptr) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r UintptrSlice) Page(number int64, perPage int64) ([]uintptr, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []uintptr{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r UintptrSlice) Any(f func(uintptr) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r UintptrSlice) Contains(el uintptr) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r UintptrSlice) GetUnion(sl2 []uintptr) ([]uintptr, error) {
	result := make([]uintptr, 0)
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
func (r UintptrSlice) InFirstOnly(sl2 []uintptr) ([]uintptr, error) {
	result := make([]uintptr, 0)
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
