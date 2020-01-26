package experiment

func (r *FakeType) equal(another *FakeType) bool {
	if r == nil && another == nil {
		return true
	}
	if (r == nil && another != nil) || (r != nil && another == nil) {
		return false
	}
	// `equal` method has to be implemented manually
	return r.A == another.A &&
		r.B == another.B
}
