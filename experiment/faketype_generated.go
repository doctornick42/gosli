package experiment

import (
	"errors"
	lib "github.com/doctornick42/gosli/lib"
)

var fakeType_var fakeType_interface

type fakeType_struct struct{}
type fakeType_interface interface {
	Where(sl []*FakeType, f func(*FakeType) bool) []*FakeType
	FirstOrDefault(sl []*FakeType, f func(*FakeType) bool) *FakeType
	First(sl []*FakeType, f func(*FakeType) bool) (*FakeType, error)
	Select(sl []*FakeType, f func(*FakeType) interface{}) []interface{}
	Page(sl []*FakeType, number int64, perPage int64) ([]*FakeType, error)
	Contains(sl []*FakeType, el *FakeType) (bool, error)
	GetUnion(sl1, sl2 []*FakeType) ([]*FakeType, error)
	InFirstOnly(sl1, sl2 []*FakeType) ([]*FakeType, error)
}

func FakeTypeSlice() fakeType_interface {
	if fakeType_var == nil {
		fakeType_var = &fakeType_struct{}
	}
	return fakeType_var
}
func (r *fakeType_struct) FirstOrDefault(sl []*FakeType, f func(*FakeType) bool) *FakeType {
	for _, slEl := range sl {
		if f(slEl) {
			return slEl
		}
	}
	return nil
}
func (r *fakeType_struct) First(sl []*FakeType, f func(*FakeType) bool) (*FakeType, error) {
	first := r.FirstOrDefault(sl, f)
	if first == nil {
		return nil, errors.New("Not found")
	}
	return first, nil
}
func (r *fakeType_struct) Where(sl []*FakeType, f func(*FakeType) bool) []*FakeType {
	res := make([]*FakeType, 0)
	for _, slEl := range sl {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r *fakeType_struct) Select(sl []*FakeType, f func(*FakeType) interface{}) []interface{} {
	res := make([]interface{}, len(sl))
	for i := range sl {
		res[i] = f(sl[i])
	}
	return res
}
func (r *fakeType_struct) Page(sl []*FakeType, number int64, perPage int64) ([]*FakeType, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(sl)) {
		return []*FakeType{}, nil
	}
	last := first + perPage
	if last > int64(len(sl)) {
		last = int64(len(sl))
	}
	return sl[first:last], nil
}
func (r *fakeType_struct) sliceToEqualers(sl []*FakeType) []lib.Equaler {
	equalerSl := make([]lib.Equaler, len(sl))
	for i := range sl {
		equalerSl[i] = sl[i]
	}
	return equalerSl
}
func (r *fakeType_struct) Contains(sl []*FakeType, el *FakeType) (bool, error) {
	equalerSl := r.sliceToEqualers(sl)
	return lib.Contains(equalerSl, el)
}
func (r *fakeType_struct) processSliceOperation(sl1, sl2 []*FakeType, f func([]lib.Equaler, []lib.Equaler) ([]lib.Equaler, error)) ([]*FakeType, error) {
	equalerSl1 := r.sliceToEqualers(sl1)
	equalerSl2 := r.sliceToEqualers(sl2)
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
func (r *fakeType_struct) GetUnion(sl1, sl2 []*FakeType) ([]*FakeType, error) {
	return r.processSliceOperation(sl1, sl2, lib.GetUnion)
}
func (r *fakeType_struct) InFirstOnly(sl1, sl2 []*FakeType) ([]*FakeType, error) {
	return r.processSliceOperation(sl1, sl2, lib.InFirstOnly)
}
func (r *FakeType) Equal(another lib.Equaler) (bool, error) {
	anotherCasted, ok := another.(*FakeType)
	if !ok {
		return false, errors.New("Types mismatch")
	}
	return r.equal(anotherCasted), nil
}
