package main

import (
	"log"

	"github.com/doctornick42/gosli/gen"
)

func main() {
	//customGenerator := &gen.CustomGenerator{}
	//err := customGenerator.Run(os.Args[1:])

	primitivesGen := &gen.PrimitivesGenerator{}
	err := primitivesGen.Run()

	if err != nil {
		log.Fatal(err)
	}
}
