package experiment

import "fmt"

func ExampleRunChainedOperations() {
	original := []FakeType{
		FakeType{
			A: 1,
			B: "one",
		},
		FakeType{
			A: 2,
			B: "two",
		},
		FakeType{
			A: 3,
			B: "three",
		},
		FakeType{
			A: 4,
			B: "four",
		},
		FakeType{
			A: 5,
			B: "five",
		},
	}

	filter1 := func(f FakeType) bool {
		return f.A > 1
	}

	filter2 := func(f FakeType) bool {
		return f.B != "three"
	}

	result, err := FakeTypeSlice(original).
		Where(filter1).
		Where(filter2).
		Page(1, 2)

	/*
		expected result to be:
		FakeType{
			A: 2,
			B: "two",
		},
		FakeType{
			A: 4,
			B: "four",
		},
		FakeType{
			A: 5,
			B: "five",
		},
	*/

	if err != nil {
		fmt.Printf("error: %v\r\n", err)
	}

	fmt.Println(len(result))
	fmt.Println(result[0].B)

	// Output:
	// 2
	// two
}

func ExampleRunChainedOperationsForPointers() {
	original := []*FakeType{
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
		&FakeType{
			A: 4,
			B: "four",
		},
		&FakeType{
			A: 5,
			B: "five",
		},
	}

	filter1 := func(f *FakeType) bool {
		return f.A > 1
	}

	filter2 := func(f *FakeType) bool {
		return f.B != "three"
	}

	anyArg := func(f *FakeType) bool {
		return f.A == 3
	}

	result := FakeTypePSlice(original).
		Where(filter1).
		Where(filter2).
		Any(anyArg)

	fmt.Println(result)

	// Output:
	// false
}
