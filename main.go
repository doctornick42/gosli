package main

import (
	"fmt"
	"log"

	"os"

	"github.com/doctornick42/gosli/experiment"
	"github.com/doctornick42/gosli/gen"
)

//argsWithoutProg := os.Args[1:]
func main() {
	err := gen.Run(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	//checkSomething()
}

func checkSomething() {
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

	first, _ := experiment.First(sl, func(t *experiment.TestType) bool {
		return t.A == 157
	})

	log.Printf("First: %v", first)

	firstOrDefault := experiment.FirstOrDefault(sl, func(t *experiment.TestType) bool {
		return t.A == 999
	})

	log.Printf("FirstOrDefault: %v", firstOrDefault)

	where := experiment.Where(sl, func(t *experiment.TestType) bool {
		return t.A > 10
	})

	log.Printf("Where: %v", where)

	type AnotherStruct struct {
		Msg string
	}

	selectRes := experiment.Select(sl, func(t *experiment.TestType) interface{} {
		return &AnotherStruct{
			Msg: fmt.Sprintf("%v-%s", t.A, t.B),
		}
	})

	for _, s := range selectRes {
		log.Print(s.(*AnotherStruct).Msg)
	}
}

// func fakeStuff() {

// 	sl := []*experiment.TestType{
// 		&experiment.TestType{
// 			A: 98,
// 			B: "Ninety eight",
// 		},
// 		&experiment.TestType{
// 			A: 157,
// 			B: "One hundred fifty seven",
// 		},
// 		&experiment.TestType{
// 			A: 4,
// 			B: "Four",
// 		},
// 	}

// 	existedEl := &experiment.TestType{
// 		A: 157,
// 		B: "One hundred fifty seven",
// 	}

// 	nonExistedEl := &experiment.TestType{
// 		A: 157,
// 		B: "whoopa!",
// 	}

// 	doesContain, err := experiment.TestTypeContains(sl, existedEl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Slice contain existed element? %v\r\n",
// 		doesContain)

// 	doesContain, err = experiment.TestTypeContains(sl, nonExistedEl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Slice contain non existed element? %v\r\n",
// 		doesContain)
// }
