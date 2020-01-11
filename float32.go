package main

import "errors"

type Float32Slice []float32

func (r Float32Slice) FirstOrDefault(f func(float32) bool) float32 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal float32
	return defVal
}
func (r Float32Slice) First(f func(float32) bool) (float32, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal float32
	return defVal, errors.New("Not found")
}
func (r Float32Slice) Where(f func(float32) bool) []float32 {
	res := make([]float32, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Float32Slice) Select(f func(float32) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Float32Slice) Page(number int64, perPage int64) ([]float32, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []float32{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Float32Slice) Any(f func(float32) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Float32Slice) Contains(el float32) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Float32Slice) GetUnion(sl2 []float32) ([]float32, error) {
	result := make([]float32, 0)
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
func (r Float32Slice) InFirstOnly(sl2 []float32) ([]float32, error) {
	result := make([]float32, 0)
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
