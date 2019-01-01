package main

import (
	"fmt"

	"github.com/doctornick42/gosli/experiment"
)

func main() {
	sl := []*experiment.TestType{
		&experiment.TestType{
			A: 98,
			B: "Ninety eight",
		},
		&experiment.TestType{
			A: 157,
			B: "One hundred fifty seven",
		},
		&experiment.TestType{
			A: 4,
			B: "Four",
		},
	}

	existedEl := &experiment.TestType{
		A: 157,
		B: "One hundred fifty seven",
	}

	nonExistedEl := &experiment.TestType{
		A: 157,
		B: "whoopa!",
	}

	doesContain, err := experiment.TestTypeContains(sl, existedEl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Slice contain existed element? %v\r\n",
		doesContain)

	doesContain, err = experiment.TestTypeContains(sl, nonExistedEl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Slice contain non existed element? %v\r\n",
		doesContain)
}
