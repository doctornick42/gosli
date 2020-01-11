package main

import "errors"

type StringSlice []string

func (r StringSlice) FirstOrDefault(f func(string) bool) string {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal string
	return defVal
}
func (r StringSlice) First(f func(string) bool) (string, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal string
	return defVal, errors.New("Not found")
}
func (r StringSlice) Where(f func(string) bool) []string {
	res := make([]string, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r StringSlice) Select(f func(string) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r StringSlice) Page(number int64, perPage int64) ([]string, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []string{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r StringSlice) Any(f func(string) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r StringSlice) Contains(el string) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r StringSlice) GetUnion(sl2 []string) ([]string, error) {
	result := make([]string, 0)
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
func (r StringSlice) InFirstOnly(sl2 []string) ([]string, error) {
	result := make([]string, 0)
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
