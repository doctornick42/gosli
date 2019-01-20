package experiment

import (
	"errors"
	lib "github.com/doctornick42/gosli/lib"
)

type FakeTypePointerSlice []*FakeType

func (r FakeTypePointerSlice) FirstOrDefault(f func(*FakeType) bool) *FakeType {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	return nil
}
func (r FakeTypePointerSlice) First(f func(*FakeType) bool) (*FakeType, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal *FakeType
	return defVal, errors.New("Not found")
}
func (r FakeTypePointerSlice) Where(f func(*FakeType) bool) []*FakeType {
	res := make([]*FakeType, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r FakeTypePointerSlice) Select(f func(*FakeType) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r FakeTypePointerSlice) Page(number int64, perPage int64) ([]*FakeType, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []*FakeType{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return r[first:last], nil
}
func (r FakeTypePointerSlice) Any(f func(*FakeType) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r FakeTypePointerSlice) sliceToEqualers() []lib.Equaler {
	equalerSl := make([]lib.Equaler, len(r))
	for i := range r {
		equalerSl[i] = r[i]
	}
	return equalerSl
}
func (r FakeTypePointerSlice) Contains(el *FakeType) (bool, error) {
	equalerSl := r.sliceToEqualers()
	return lib.Contains(equalerSl, el)
}
func (r FakeTypePointerSlice) processSliceOperation(sl2 FakeTypePointerSlice, f func([]lib.Equaler, []lib.Equaler) ([]lib.Equaler, error)) ([]*FakeType, error) {
	equalerSl1 := r.sliceToEqualers()
	equalerSl2 := sl2.sliceToEqualers()
	untypedRes, err := f(equalerSl1, equalerSl2)
	if err != nil {
		return nil, err
	}
	res := make([]*FakeType, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i].(*FakeType)
	}
	return res, nil
}
func (r FakeTypePointerSlice) GetUnion(sl2 []*FakeType) ([]*FakeType, error) {
	return r.processSliceOperation(sl2, lib.GetUnion)
}
func (r FakeTypePointerSlice) InFirstOnly(sl2 []*FakeType) ([]*FakeType, error) {
	return r.processSliceOperation(sl2, lib.InFirstOnly)
}
func (r *FakeType) Equal(another lib.Equaler) (bool, error) {
	anotherCasted, ok := another.(*FakeType)
	if !ok {
		return false, errors.New("Types mismatch")
	}
	return r.equal(anotherCasted), nil
}
