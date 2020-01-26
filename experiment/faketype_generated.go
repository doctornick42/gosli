package experiment

import (
	"errors"
	lib "github.com/doctornick42/gosli/lib"
)

type FakeTypeSlice []FakeType

func (r FakeTypeSlice) FirstOrDefault(f func(FakeType) bool) FakeType {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal FakeType
	return defVal
}
func (r FakeTypeSlice) First(f func(FakeType) bool) (FakeType, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal FakeType
	return defVal, errors.New("Not found")
}
func (r FakeTypeSlice) Where(f func(FakeType) bool) FakeTypeSlice {
	res := make([]FakeType, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return FakeTypeSlice(res)
}
func (r FakeTypeSlice) Select(f func(FakeType) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r FakeTypeSlice) Page(number int64, perPage int64) (FakeTypeSlice, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []FakeType{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return FakeTypeSlice(r[first:last]), nil
}
func (r FakeTypeSlice) Any(f func(FakeType) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r FakeTypeSlice) sliceToEqualers() []lib.Equaler {
	equalerSl := make([]lib.Equaler, len(r))
	for i := range r {
		equalerSl[i] = &r[i]
	}
	return equalerSl
}
func (r FakeTypeSlice) Contains(el FakeType) (bool, error) {
	equalerSl := r.sliceToEqualers()
	return lib.Contains(equalerSl, &el)
}
func (r FakeTypeSlice) processSliceOperation(sl2 FakeTypeSlice, f func([]lib.Equaler, []lib.Equaler) ([]lib.Equaler, error)) (FakeTypeSlice, error) {
	equalerSl1 := r.sliceToEqualers()
	equalerSl2 := sl2.sliceToEqualers()
	untypedRes, err := f(equalerSl1, equalerSl2)
	if err != nil {
		return nil, err
	}
	res := make([]FakeType, len(untypedRes))
	for i := range untypedRes {
		res[i] = *untypedRes[i].(*FakeType)
	}
	return FakeTypeSlice(res), nil
}
func (r FakeTypeSlice) GetUnion(sl2 []FakeType) (FakeTypeSlice, error) {
	return r.processSliceOperation(sl2, lib.GetUnion)
}
func (r FakeTypeSlice) InFirstOnly(sl2 []FakeType) (FakeTypeSlice, error) {
	return r.processSliceOperation(sl2, lib.InFirstOnly)
}
