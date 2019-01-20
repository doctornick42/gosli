package main

import "errors"

type RuneSlice []rune

func (r RuneSlice) FirstOrDefault(f func(rune) bool) rune {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal rune
	return defVal
}
func (r RuneSlice) First(f func(rune) bool) (rune, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal rune
	return defVal, errors.New("Not found")
}
func (r RuneSlice) Where(f func(rune) bool) []rune {
	res := make([]rune, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r RuneSlice) Select(f func(rune) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r RuneSlice) Page(number int64, perPage int64) ([]rune, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []rune{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r RuneSlice) Any(f func(rune) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r RuneSlice) Contains(el rune) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r RuneSlice) GetUnion(sl2 []rune) ([]rune, error) {
	result := make([]rune, 0)
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
func (r RuneSlice) InFirstOnly(sl2 []rune) ([]rune, error) {
	result := make([]rune, 0)
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
