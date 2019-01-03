package experiment

import (
	"testing"
)

func BenchmarkFirst(b *testing.B) {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	for n := 0; n < b.N; n++ {
		var filter func(*FakeType) bool

		if n%2 == 0 {
			filter = func(t *FakeType) bool {
				return t.A == 2
			}
		} else {
			filter = func(t *FakeType) bool {
				return t.A == 123
			}
		}

		FakeTypeSlice().First(sl, filter)
	}
}

func BenchmarkFirstOrDefault(b *testing.B) {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	for n := 0; n < b.N; n++ {
		var filter func(*FakeType) bool

		if n%2 == 0 {
			filter = func(t *FakeType) bool {
				return t.A == 2
			}
		} else {
			filter = func(t *FakeType) bool {
				return t.A == 123
			}
		}

		FakeTypeSlice().FirstOrDefault(sl, filter)
	}
}

func BenchmarkSelect(b *testing.B) {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	for n := 0; n < b.N; n++ {
		filter := func(t *FakeType) interface{} {
			return struct {
				smthng string
			}{
				smthng: t.B,
			}
		}

		FakeTypeSlice().Select(sl, filter)
	}
}

func BenchmarkWhere(b *testing.B) {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	for n := 0; n < b.N; n++ {
		var filter func(*FakeType) bool

		if n%2 == 0 {
			filter = func(t *FakeType) bool {
				return t.A > 2
			}
		} else {
			filter = func(t *FakeType) bool {
				return t.A > 1000
			}
		}

		FakeTypeSlice().Where(sl, filter)
	}
}

func BenchmarkContains(b *testing.B) {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	for n := 0; n < b.N; n++ {
		var el *FakeType

		if n%2 == 0 {
			el = &FakeType{
				A: 2,
				B: "two",
			}
		} else {
			el = &FakeType{
				A: 100500,
				B: "whoa!",
			}
		}

		FakeTypeSlice().Contains(sl, el)
	}
}

func BenchmarkGetUnion(b *testing.B) {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	sl2WithUnion := []*FakeType{
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
		&FakeType{
			A: 4,
			B: "four",
		},
	}

	sl2WithoutUnion := []*FakeType{
		&FakeType{
			A: 4,
			B: "four",
		},
		&FakeType{
			A: 5,
			B: "five",
		},
	}

	for n := 0; n < b.N; n++ {
		var sl2 []*FakeType

		if n%2 == 0 {
			sl2 = sl2WithUnion
		} else {
			sl2 = sl2WithoutUnion
		}

		FakeTypeSlice().GetUnion(sl, sl2)
	}
}

func BenchmarkInFirstOnly(b *testing.B) {
	sl := []*FakeType{
		&FakeType{
			A: 1,
			B: "one",
		},
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
	}

	sl2WithUnion := []*FakeType{
		&FakeType{
			A: 2,
			B: "two",
		},
		&FakeType{
			A: 3,
			B: "three",
		},
		&FakeType{
			A: 4,
			B: "four",
		},
	}

	sl2WithoutUnion := []*FakeType{
		&FakeType{
			A: 4,
			B: "four",
		},
		&FakeType{
			A: 5,
			B: "five",
		},
	}

	for n := 0; n < b.N; n++ {
		var sl2 []*FakeType

		if n%2 == 0 {
			sl2 = sl2WithUnion
		} else {
			sl2 = sl2WithoutUnion
		}

		FakeTypeSlice().InFirstOnly(sl, sl2)
	}
}
