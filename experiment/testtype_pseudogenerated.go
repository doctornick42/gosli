package experiment

//all code after this line is pseudo-generated:

// func sliceToEqualers(sl []*TestType) []lib.Equaler {
// 	equalerSl := make([]lib.Equaler, len(sl))
// 	for i := range sl {
// 		equalerSl[i] = sl[i]
// 	}

// 	return equalerSl
// }

// func sliceToInterfacesSlice(sl []*TestType) []interface{} {
// 	equalerSl := make([]interface{}, len(sl))
// 	for i := range sl {
// 		equalerSl[i] = sl[i]
// 	}

// 	return equalerSl
// }

// func TestTypeContains(sl []*TestType, el *TestType) (bool, error) {
// 	equalerSl := sliceToEqualers(sl)
// 	return lib.Contains(equalerSl, el)
// }

// func TestTypeGetUnion(sl1, sl2 []*TestType) ([]*TestType, error) {
// 	return processSliceOperation(sl1, sl2, lib.GetUnion)
// }

// func TestTypeInFirstOnly(sl1, sl2 []*TestType) ([]*TestType, error) {
// 	return processSliceOperation(sl1, sl2, lib.InFirstOnly)
// }

// func processSliceOperation(sl1, sl2 []*TestType,
// 	f func([]lib.Equaler, []lib.Equaler) ([]lib.Equaler, error)) ([]*TestType, error) {

// 	equalerSl1 := sliceToEqualers(sl1)
// 	equalerSl2 := sliceToEqualers(sl2)

// 	untypedRes, err := f(equalerSl1, equalerSl2)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := make([]*TestType, len(untypedRes))
// 	for i := range untypedRes {
// 		res[i] = untypedRes[i].(*TestType)
// 	}

// 	return res, nil
// }
