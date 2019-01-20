package gen

import (
	"fmt"
	"log"

	"github.com/dave/jennifer/jen"
)

const (
	primitivesModuleName = "primitives"
)

var (
	AvailableTypes = []string{
		"int8",
	}
)

type PrimitivesGenerator struct{}

func (g *PrimitivesGenerator) Run() error {
	for _, typeName := range AvailableTypes {

		f := jen.NewFile(primitivesModuleName)
		g.generateInfrastructure(f, typeName)
		g.generateFirst(f, typeName)

		fakeOriginPath := fmt.Sprintf("%s/fake.go", primitivesModuleName)

		genFileName := g.getGeneratedFileName(fakeOriginPath, typeName)

		log.Printf("Generated filename: %s", genFileName)
		err := f.Save(genFileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *PrimitivesGenerator) generateInfrastructure(f *jen.File, typeName string) {
	generateInfrastructure(f, typeName)
}

func (g *PrimitivesGenerator) getGeneratedFileName(originFilePath, typeName string) string {
	return generateFileName(originFilePath, "", typeName)
}

func (g *PrimitivesGenerator) generateFirst(f *jen.File, typeName string) {
	generateFirst(f, typeName)
}
