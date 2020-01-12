package experiment

func (r *FakeType) equal(another *FakeType) bool {
	// `equal` method has to be implemented manually
	if r == nil && another == nil {
		return true
	}

	if (r == nil && another != nil) ||
		(r != nil && another == nil) {

		return false
	}

	return r.A == another.A &&
		r.B == another.B
}
