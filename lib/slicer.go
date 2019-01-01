package lib

import "errors"

type Equaler interface {
	Equal(Equaler) (bool, error)
}

func Contains(sl []Equaler, el Equaler) (bool, error) {
	for _, slEl := range sl {
		isEqual, err := slEl.Equal(el)
		if err != nil {
			return false, err
		}

		if isEqual {
			return true, nil
		}
	}

	return false, nil
}

func GetUnion(sl1, sl2 []Equaler) ([]Equaler, error) {
	result := make([]Equaler, 0)

	for _, sl1El := range sl1 {
		for _, sl2El := range sl2 {
			areEqual, err := sl1El.Equal(sl2El)
			if err != nil {
				return nil, err
			}
			if areEqual {
				result = append(result, sl1El)
				break
			}
		}
	}

	return result, nil
}

func InFirstOnly(sl1, sl2 []Equaler) ([]Equaler, error) {
	result := make([]Equaler, 0)

	for _, sl1El := range sl1 {
		found := false
		for _, sl2El := range sl2 {
			areEqual, err := sl1El.Equal(sl2El)
			if err != nil {
				return nil, err
			}
			if areEqual {
				found = true
				break
			}
		}

		if !found {
			result = append(result, sl1El)
		}
	}

	return result, nil
}

func FirstOrDefault(sl []interface{}, f func(interface{}) bool) interface{} {
	for _, slEl := range sl {
		if f(slEl) {
			return slEl
		}
	}

	return nil
}

func First(sl []interface{}, f func(interface{}) bool) (interface{}, error) {
	first := FirstOrDefault(sl, f)

	if first == nil {
		return nil, errors.New("Not found")
	}

	return first, nil
}

func Where(sl []interface{}, f func(interface{}) bool) []interface{} {
	res := make([]interface{}, 0)

	for _, slEl := range sl {
		if f(slEl) {
			res = append(res, slEl)
		}
	}

	return res
}

func Select(sl []interface{}, f func(interface{}) interface{}) []interface{} {
	res := make([]interface{}, len(sl))

	for i := range sl {
		res[i] = f(sl[i])
	}

	return res
}
