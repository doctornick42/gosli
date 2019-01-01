package experiment

func FirstOrDefault(sl []*TestType, f func(*TestType) bool) *TestType {
	for _, slEl := range sl {
		if f(slEl) {
			return slEl
		}
	}
	return nil
}
