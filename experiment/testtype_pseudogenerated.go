package experiment

//all code after this line is pseudo-generated:

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
