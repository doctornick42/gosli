package experiment

import (
	"errors"
	lib "github.com/doctornick42/gosli/lib"
)

func FakeTypeFirstOrDefault(sl []*FakeType, f func(*FakeType) bool) *FakeType {
	for _, slEl := range sl {
		if f(slEl) {
			return slEl
		}
	}
	return nil
}
func FakeTypeFirst(sl []*FakeType, f func(*FakeType) bool) (*FakeType, error) {
	first := FakeTypeFirstOrDefault(sl, f)
	if first == nil {
		return nil, errors.New("Not found")
	}
	return first, nil
}
func FakeTypeWhere(sl []*FakeType, f func(*FakeType) bool) []*FakeType {
	res := make([]*FakeType, 0)
	for _, slEl := range sl {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return res
}
func (r *FakeType) Equal(another lib.Equaler) (bool, error) {
	anotherCasted, ok := another.(*FakeType)
	if !ok {
		return false, errors.New("Types mismatch")
	}
	return r.equal(anotherCasted), nil
}
func FakeTypeSelect(sl []*FakeType, f func(*FakeType) interface{}) []interface{} {
	res := make([]interface{}, len(sl))
	for i := range sl {
		res[i] = f(sl[i])
	}
	return res
}
func FakeTypeSliceToEqualers(sl []*FakeType) []lib.Equaler {
	equalerSl := make([]lib.Equaler, len(sl))
	for i := range sl {
		equalerSl[i] = sl[i]
	}
	return equalerSl
}
func FakeTypeSliceToInterfacesSlice(sl []*FakeType) []interface{} {
	equalerSl := make([]interface{}, len(sl))
	for i := range sl {
		equalerSl[i] = sl[i]
	}
	return equalerSl
}
func FakeTypeContains(sl []*FakeType, el *FakeType) (bool, error) {
	equalerSl := FakeTypeSliceToEqualers(sl)
	return lib.Contains(equalerSl, el)
}
func FakeTypeProcessSliceOperation(sl1, sl2 []*FakeType, f func([]lib.Equaler, []lib.Equaler) ([]lib.Equaler, error)) ([]*FakeType, error) {
	equalerSl1 := FakeTypeSliceToEqualers(sl1)
	equalerSl2 := FakeTypeSliceToEqualers(sl2)
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
func FakeTypeGetUnion(sl1, sl2 []*FakeType) ([]*FakeType, error) {
	return FakeTypeProcessSliceOperation(sl1, sl2, lib.GetUnion)
}
func FakeTypeInFirstOnly(sl1, sl2 []*FakeType) ([]*FakeType, error) {
	return FakeTypeProcessSliceOperation(sl1, sl2, lib.InFirstOnly)
}
