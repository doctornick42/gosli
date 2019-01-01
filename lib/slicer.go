package lib

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
