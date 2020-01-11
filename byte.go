package main

import "errors"

type ByteSlice []byte

func (r ByteSlice) FirstOrDefault(f func(byte) bool) byte {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal byte
	return defVal
}
func (r ByteSlice) First(f func(byte) bool) (byte, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal byte
	return defVal, errors.New("Not found")
}
func (r ByteSlice) Where(f func(byte) bool) []byte {
	res := make([]byte, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r ByteSlice) Select(f func(byte) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r ByteSlice) Page(number int64, perPage int64) ([]byte, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []byte{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r ByteSlice) Any(f func(byte) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r ByteSlice) Contains(el byte) (bool, error) {
	for _, slEl := range r {
		if slEl == el {
			return true, nil
		}
	}
	return false, nil
}
func (r ByteSlice) GetUnion(sl2 []byte) ([]byte, error) {
	result := make([]byte, 0)
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
func (r ByteSlice) InFirstOnly(sl2 []byte) ([]byte, error) {
	result := make([]byte, 0)
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
