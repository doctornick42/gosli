package main

import (
	"log"
	"os"

	"github.com/doctornick42/gosli/gen"
)

const (
	primitivesRunner = "primitives"
)

func main() {

	var err error

	if os.Args[1] == primitivesRunner {
		primitivesGen := &gen.PrimitivesGenerator{}
		err = primitivesGen.Run()
	} else {
		customGenerator := &gen.CustomGenerator{}
		err = customGenerator.Run(os.Args[1:])
	}

	if err != nil {
		log.Fatal(err)
	}
}
