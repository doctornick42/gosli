package main

import (
	"log"
	"os"

	"github.com/doctornick42/gosli/gen"
)

func main() {
	customGenerator := &gen.CustomGenerator{}
	err := customGenerator.Run(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
}
