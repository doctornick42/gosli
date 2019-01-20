package main

import "errors"

type Float64Slice []float64

func (r Float64Slice) FirstOrDefault(f func(float64) bool) float64 {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal float64
	return defVal
}
func (r Float64Slice) First(f func(float64) bool) (float64, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal float64
	return defVal, errors.New("Not found")
}
func (r Float64Slice) Where(f func(float64) bool) []float64 {
	res := make([]float64, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r Float64Slice) Select(f func(float64) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r Float64Slice) Page(number int64, perPage int64) ([]float64, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []float64{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r Float64Slice) Any(f func(float64) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r Float64Slice) Contains(el float64) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r Float64Slice) GetUnion(sl2 []float64) ([]float64, error) {
	result := make([]float64, 0)
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
func (r Float64Slice) InFirstOnly(sl2 []float64) ([]float64, error) {
	result := make([]float64, 0)
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
