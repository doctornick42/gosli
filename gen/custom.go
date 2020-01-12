package gen

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
)

type CustomGenerator struct{}

func (g *CustomGenerator) run(typeName, moduleName, originFilePath string) error {
	f := NewFile(moduleName)
	generateInfrastructure(f, typeName)
	generateFirstOrDefault(f, typeName)
	generateFirst(f, typeName)
	generateWhere(f, typeName)
	generateSelect(f, typeName)
	generatePage(f, typeName)
	generateAny(f, typeName)
	g.generateSliceToEqualers(f, typeName)
	g.generateContains(f, typeName)
	g.generateProcessSliceOperation(f, typeName)
	g.generateGetUnion(f, typeName)
	g.generateInFirstOnly(f, typeName)
	g.generateEqualImplementation(f, typeName)

	//pureTypeName := strings.TrimLeft(typeName, "*")

	genFileName := g.getGeneratedFileName(originFilePath, typeName)

	log.Printf("Generated filename: %s", genFileName)
	err := f.Save(genFileName)
	if err != nil {
		return err
	}

	genFileName = g.getEqualGeneratedFileName(originFilePath, typeName)
	if _, err := os.Stat(genFileName); os.IsNotExist(err) {
		f = NewFile(moduleName)
		g.generateEqualToFillManually(f, typeName)

		log.Printf("Generated filename: %s", genFileName)
		return f.Save(genFileName)
	}

	return nil
}

func (g *CustomGenerator) Run(args []string) error {
	if len(args) < 2 {
		return errors.New("Wrong amount of arguments")
	}

	originFilePath := args[0]
	typeName := args[1]

	moduleName, err := g.getModuleName(originFilePath)
	if err != nil {
		return err
	}

	log.Printf("Module name: %s", moduleName)

	err = g.run(typeName, moduleName, originFilePath)
	if err != nil {
		return err
	}

	return g.run("*"+typeName, moduleName, originFilePath)
}

func (g *CustomGenerator) getModuleName(originFilePath string) (string, error) {
	f, err := os.Open(originFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	firstLine := ""
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "package") {
			firstLine = sc.Text()
			break
		}
	}

	if len(firstLine) == 0 {
		return "", errors.New("Package name not found in the specified file")
	}

	firstLineSplitted := strings.Split(firstLine, " ")
	return firstLineSplitted[len(firstLineSplitted)-1], nil
}

func (g *CustomGenerator) getGeneratedFileName(originFilePath, typeName string) string {
	suffix := "generated"
	if typeName[0] == '*' {
		suffix = "p_" + suffix
		typeName = strings.TrimLeft(typeName, "*")
	}
	return generateFileName(originFilePath, suffix, typeName)
}

func (g *CustomGenerator) getEqualGeneratedFileName(originFilePath, typeName string) string {
	suffix := "equal"
	if typeName[0] == '*' {
		suffix = "p_" + suffix
		typeName = strings.TrimLeft(typeName, "*")
	}
	return generateFileName(originFilePath, suffix, typeName)
}

func (g *CustomGenerator) generateEqualToFillManually(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(typeName),
		).Id("equal").
		Params(
			Id("another").Id(typeName),
		).
		Bool().
		Block(
			Comment("`equal` method has to be implemented manually"),
		)
}

func (g *CustomGenerator) generateEqualImplementation(f *File, typeName string) {
	f.Func().
		Params(Id("r").Id(typeName)).
		Id("Equal").
		Params(
			Id("another").Qual("github.com/doctornick42/gosli/lib", "Equaler"),
		).
		Params(
			Bool(),
			Error(),
		).
		Block(
			Id("anotherCasted, ok").Op(":=").Id("another").Dot(fmt.Sprintf("(%s)", typeName)),

			If(
				Id("!ok"),
			).Block(
				Return(False(), Qual("errors", "New").Call(Lit("Types mismatch"))),
			),

			Return(Id("r.equal").Call(Id("anotherCasted")), Nil()),
		)
}

func (g *CustomGenerator) generateSliceToEqualers(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("sliceToEqualers").
		Params().
		Index().Qual("github.com/doctornick42/gosli/lib", "Equaler").
		Block(
			Id("equalerSl").Op(":=").Make(Id("[]lib.Equaler"), Len(Id("r"))),

			For(
				Id("i").Op(":=").Range().Id("r"),
			).Block(
				Id("equalerSl[i]").Op("=").Id("r[i]"),
			),

			Return(Id("equalerSl")),
		)
}

func (g *CustomGenerator) generateContains(f *File, typeName string) {
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
			Id("equalerSl").Op(":=").Id("r.sliceToEqualers").Call(),
			Return(
				Qual("github.com/doctornick42/gosli/lib", "Contains").
					Call(Id("equalerSl"), Id("el")),
			),
		)
}

func (g *CustomGenerator) generateProcessSliceOperation(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("processSliceOperation").
		Params(
			Id("sl2").Id(getStructName(typeName)),
			Id("f").Func().Params(
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
			).Params(
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
				Error(),
			),
		).
		Params(
			Id("[]"+typeName),
			Error(),
		).
		Block(
			Id("equalerSl1").Op(":=").Id("r.sliceToEqualers").Call(),
			Id("equalerSl2").Op(":=").Id("sl2.sliceToEqualers").Call(),
			Id("untypedRes, err").Op(":=").Id("f").Call(Id("equalerSl1"), Id("equalerSl2")),

			If(
				Id("err").Op("!=").Nil(),
			).Block(
				Return(Nil(), Id("err")),
			),

			Id("res").Op(":=").Make(Index().Id(typeName), Len(Id("untypedRes"))),

			For(
				Id("i").Op(":=").Range().Id("untypedRes"),
			).Block(
				Id("res[i]").Op("=").Id("untypedRes[i]").Dot(fmt.Sprintf("(%s)", typeName)),
			),

			Return(Id("res"), Nil()),
		)
}

func (g *CustomGenerator) generateGetUnion(f *File, typeName string) {
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
			Return(
				Id("r.processSliceOperation").Call(
					Id("sl2"),
					Qual("github.com/doctornick42/gosli/lib", "GetUnion"),
				),
			),
		)
}

func (g *CustomGenerator) generateInFirstOnly(f *File, typeName string) {
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
			Return(
				Id("r.processSliceOperation").Call(
					Id("sl2"),
					Qual("github.com/doctornick42/gosli/lib", "InFirstOnly"),
				),
			),
		)
}
