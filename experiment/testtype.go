package experiment

type TestType struct {
	A int
	B string
}

//signature is generated
//but implementation has to be done manually:
func (r *TestType) equal(another *TestType) bool {
	return r.A == another.A &&
		r.B == another.B
}
