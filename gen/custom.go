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

	if typeName[0] == '*' {
		g.generateEqualImplementation(f, typeName)
	}
	//pureTypeName := strings.TrimLeft(typeName, "*")

	genFileName := g.getGeneratedFileName(originFilePath, typeName)

	log.Printf("Generated filename: %s", genFileName)
	err := f.Save(genFileName)
	if err != nil {
		return err
	}

	if typeName[0] == '*' {
		genFileName = g.getEqualGeneratedFileName(originFilePath, typeName)
		if _, err := os.Stat(genFileName); os.IsNotExist(err) {
			f = NewFile(moduleName)
			g.generateEqualToFillManually(f, typeName)

			log.Printf("Generated filename: %s", genFileName)
			return f.Save(genFileName)
		}
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
			If(
				Id("r").Op("==").Nil().Op("&&").
					Id("another").Op("==").Nil(),
			).Block(Return(True())),

			If(
				Parens(
					Id("r").Op("==").Nil().Op("&&").
						Id("another").Op("!=").Nil(),
				).Op("||").
					Parens(
						Id("r").Op("!=").Nil().Op("&&").
							Id("another").Op("==").Nil(),
					),
			).Block(Return(False())),

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
	sliceEl := "r[i]"
	if !g.isTypePointer(typeName) {
		sliceEl = "&" + sliceEl
	}
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
				Id("equalerSl[i]").Op("=").Id(sliceEl),
			),

			Return(Id("equalerSl")),
		)
}

func (g *CustomGenerator) isTypePointer(typeName string) bool {
	return typeName[0] == '*'
}

func (g *CustomGenerator) generateContains(f *File, typeName string) {
	elArg := "el"
	if !g.isTypePointer(typeName) {
		elArg = "&" + elArg
	}
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
					Call(Id("equalerSl"), Id(elArg)),
			),
		)
}

func (g *CustomGenerator) generateProcessSliceOperation(f *File, typeName string) {
	castingType := typeName
	untypedResEl := "untypedRes[i]"
	structName := getStructName(typeName)
	if g.isTypePointer(typeName) {
	} else {
		untypedResEl = "*" + untypedResEl
		castingType = "*" + castingType
	}
	f.Func().
		Params(
			Id("r").Id(structName),
		).
		Id("processSliceOperation").
		Params(
			Id("sl2").Id(structName),
			Id("f").Func().Params(
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
			).Params(
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
				Error(),
			),
		).
		Params(
			Id(structName),
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
				Id("res[i]").Op("=").Id(untypedResEl).Dot(fmt.Sprintf("(%s)", castingType)),
			),

			Return(
				Id(structName).Call(Id("res")),
				Nil(),
			),
		)
}

func (g *CustomGenerator) generateGetUnion(f *File, typeName string) {
	structName := getStructName(typeName)
	f.Func().
		Params(
			Id("r").Id(structName),
		).
		Id("GetUnion").
		Params(
			Id("sl2").Index().Id(typeName),
		).
		Params(
			Id(structName),
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
	structName := getStructName(typeName)
	f.Func().
		Params(
			Id("r").Id(structName),
		).
		Id("InFirstOnly").
		Params(
			Id("sl2").Index().Id(typeName),
		).
		Params(
			Id(structName),
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
