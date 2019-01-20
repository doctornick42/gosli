package gen

import (
	"fmt"
	"log"

	. "github.com/dave/jennifer/jen"
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

		f := NewFile(primitivesModuleName)
		g.generateInfrastructure(f, typeName)
		generateFirstOrDefault(f, typeName)
		generateFirst(f, typeName)
		generateWhere(f, typeName)
		generateSelect(f, typeName)
		generatePage(f, typeName)
		generateAny(f, typeName)
		g.generateContains(f, typeName)
		g.generateGetUnion(f, typeName)
		g.generateInFirstOnly(f, typeName)
		// g.generateEqualImplementation(f, "*"+typeName)

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

func (g *PrimitivesGenerator) generateInfrastructure(f *File, typeName string) {
	generateInfrastructure(f, typeName)
}

func (g *PrimitivesGenerator) getGeneratedFileName(originFilePath, typeName string) string {
	return generateFileName(originFilePath, "", typeName)
}

func (g *PrimitivesGenerator) generateFirst(f *File, typeName string) {
	generateFirst(f, typeName)
}

func (g *PrimitivesGenerator) generateContains(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("Contains").
		Params(
			Id("el").Id(typeName),
		).
		Params(
			Bool(),
			Error(),
		).
		Block(
			For(
				Id("_, slEl").Op(":=").Range().Id("r").Block(
					If(
						Id("slEl").Op("==").Id("el"),
					).Block(
						Return(True(), Nil()),
					),
				),
			),

			Return(False(), Nil()),
		)
}

func (g *PrimitivesGenerator) getEqualStatement(el1, el2 string, typeName string) []Code {
	if string(typeName[0]) == "*" {
		return []Code{
			Var().Id("areEqual").Bool(),

			If(
				Id(el1).Op("==").Nil().Op("&&").
					Id(el2).Op("==").Nil(),
			).Block(
				Id("areEqual").Op("=").True(),
			),

			If(
				Params(
					Id(el1).Op("!=").Nil().Op("&&").
						Id(el2).Op("==").Nil()).
					Op("||").
					Params(
						Id(el1).Op("==").Nil().Op("&&").
							Id(el2).Op("!=").Nil(),
					),
			).Block(
				Id("areEqual").Op("=").False(),
			),

			Id("areEqual").Op("=").Id("*" + el1).Op("==").Id("*" + el2),
		}
	}

	return []Code{
		Id("areEqual").Op(":=").Id(el1).Op("==").Id(el2),
	}
}

func (g *PrimitivesGenerator) generateGetUnion(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("GetUnion").
		Params(
			Id("sl2").Index().Id(typeName),
		).
		Params(
			Index().Id(typeName),
			Error(),
		).
		Block(
			Id("result").Op(":=").Make(Index().Id(typeName), Lit(0)),

			For(
				Id("_, sl1El").Op(":=").Range().Id("r").Block(
					For(
						Id("_, sl2El").Op(":=").Range().Id("sl2").Block(
							append(
								g.getEqualStatement("sl1El", "sl2El", typeName),
								If(
									Id("areEqual"),
								).Block(
									Id("result").Op("=").Append(Id("result"), Id("sl1El")),
								),
							)...,
						),
					),
				),
			),

			Return(Id("result"), Nil()),
		)
}

func (g *PrimitivesGenerator) generateInFirstOnly(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("InFirstOnly").
		Params(
			Id("sl2").Index().Id(typeName),
		).
		Params(
			Index().Id(typeName),
			Error(),
		).
		Block(
			Id("result").Op(":=").Make(Index().Id(typeName), Lit(0)),

			For(
				Id("_, sl1El").Op(":=").Range().Id("r").Block(
					Id("found").Op(":=").False(),

					For(
						Id("_, sl2El").Op(":=").Range().Id("sl2").Block(
							append(
								g.getEqualStatement("sl1El", "sl2El", typeName),
								If(
									Id("areEqual"),
								).Block(
									Id("found").Op("=").True(),
									Continue(),
								),
							)...,
						),
					),

					If(Id("!found")).Block(
						Id("result").Op("=").Append(Id("result"), Id("sl1El")),
					),
				),
			),

			Return(Id("result"), Nil()),
		)
}
