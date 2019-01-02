package main

import (
	"log"
	"os"

	"github.com/doctornick42/gosli/gen"
)

func main() {
	err := gen.Run(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
}
