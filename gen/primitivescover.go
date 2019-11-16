package gen

import (
	"fmt"
	"log"

	. "github.com/dave/jennifer/jen"
)

type PrimitivesTestsGenerator struct{}

func (g *PrimitivesTestsGenerator) Run() error {
	availableTypes := []string{"int"}

	for _, typeName := range availableTypes {

		f := NewFile("main")
		g.generateInfrastructure(f, typeName)

		fakeOriginPath := fmt.Sprintf("fake.go")

		genFileName := g.getGeneratedFileName(fakeOriginPath, typeName)

		log.Printf("Generated filename: %s", genFileName)
		err := f.Save(genFileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *PrimitivesTestsGenerator) generateInfrastructure(f *File, typeName string) {
	generateTestInfrastructure(f, typeName)
}

func generateTestInfrastructure(f *File, typeName string) {
	uppercaseName := firstRuneToUpper(typeName)
	suiteStructName := uppercaseName + "TestSuite"

	f.Type().Id(suiteStructName).Struct(
		Qual("github.com/stretchr/testify/suite", "Suite"),
		Id("mockCtrl").Id("*").Qual("github.com/golang/mock/gomock", "Controller"),
	)

	factoryName := "Test" + uppercaseName + "Factory"

	f.Func().Id(factoryName).Params(
		Id("t").Id("*").Qual("testing", "T"),
	).Block(
		Id("suite.Run").Params(
			Id("t"),
			New(Id(suiteStructName)),
		),
	)
}

func (g *PrimitivesTestsGenerator) getGeneratedFileName(originFilePath, typeName string) string {
	return generateFileName(originFilePath, "test", typeName)
}
