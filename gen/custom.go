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

	f := NewFile(moduleName)
	g.generateInfrastructure(f, typeName)
	g.generateFirstOrDefault(f, typeName)
	g.generateFirst(f, typeName)
	g.generateWhere(f, typeName)
	g.generateSelect(f, typeName)
	g.generatePage(f, typeName)
	g.generateAny(f, typeName)
	g.generateSliceToEqualers(f, typeName)
	g.generateContains(f, typeName)
	g.generateProcessSliceOperation(f, typeName)
	g.generateGetUnion(f, typeName)
	g.generateInFirstOnly(f, typeName)
	g.generateEqualImplementation(f, typeName)

	genFileName := g.getGeneratedFileName(originFilePath, typeName)

	log.Printf("Generated filename: %s", genFileName)
	err = f.Save(genFileName)
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
	return g.generateFileName(originFilePath, "generated", typeName)
}

func (g *CustomGenerator) getEqualGeneratedFileName(originFilePath, typeName string) string {
	return g.generateFileName(originFilePath, "equal", typeName)
}

func (g *CustomGenerator) generateFileName(originFilePath, suffix, typeName string) string {
	splitted := strings.Split(originFilePath, "/")

	shortFileName := splitted[len(splitted)-1]
	generatedName := strings.ToLower(typeName) + "_" + suffix + ".go"

	return strings.Replace(originFilePath, shortFileName, generatedName, 1)
}

func (g *CustomGenerator) generateFirstOrDefault(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("FirstOrDefault").
		Params(
			Id("f").Id("func(*"+typeName+") bool"),
		).
		Id("*"+typeName).
		Block(
			For(
				Id("_, slEl").Op(":=").Range().Id("r").Block(
					If(
						Id("f").Call(Id("slEl")),
					).Block(
						Return(Id("slEl")),
					),
				),
			),
			Return(Nil()),
		)
}

func (g *CustomGenerator) generateFirst(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("First").
		Params(
			Id("f").Id("func(*"+typeName+") bool"),
		).
		Params(Id("*"+typeName), Error()).
		Block(
			Id("first").Op(":=").Id("r.FirstOrDefault").Call(Id("f")),
			If(
				Id("first").Op("==").Nil(),
			).Block(
				Return(Nil(), Qual("errors", "New").Call(Lit("Not found"))),
			),
			Return(Id("first"), Nil()),
		)
}

func (g *CustomGenerator) generateWhere(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("Where").
		Params(
			Id("f").Id("func(*"+typeName+") bool"),
		).
		Id("[]*"+typeName).
		Block(
			Id("res").Op(":=").Make(Id("[]*"+typeName), Lit(0)),

			For(
				Id("_, slEl").Op(":=").Range().Id("r").Block(
					If(
						Id("f").Call(Id("slEl")),
					).Block(
						Id("res").Op("=").Append(Id("res"), Id("slEl")),
					),
				),
			),
			Return(Id("res")),
		)
}

func (g *CustomGenerator) generateSelect(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("Select").
		Params(
			Id("f").Id(fmt.Sprintf("func(*%s) interface{}", typeName)),
		).
		Id("[]interface{}").
		Block(
			Id("res").Op(":=").Make(Id("[]interface{}"), Len(Id("r"))),

			For(
				Id("i").Op(":=").Range().Id("r").Block(
					Id("res").Index(Id("i")).Op("=").
						Id("f").Call(Id("r").Index(Id("i"))),
				),
			),
			Return(Id("res")),
		)
}

func (g *CustomGenerator) generatePage(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("Page").
		Params(
			Id("number").Int64(),
			Id("perPage").Int64(),
		).
		Params(
			Id("[]*"+typeName),
			Error(),
		).
		Block(
			If(
				Id("number").Op("<=").Lit(0),
			).Block(
				Return(Nil(), Qual("errors", "New").Params(Lit("Page number should start with 1"))),
			),

			Id("number").Op("--"),

			Id("first").Op(":=").Id("number").Op("*").Id("perPage"),

			If(
				Id("first").Op(">").Int64().Params(Len(Id("r"))),
			).Block(
				Return(
					Index().Id("*"+typeName).Block(), Nil(),
				),
			),

			Id("last").Op(":=").Id("first").Op("+").Id("perPage"),

			If(
				Id("last").Op(">").Int64().Params(Len(Id("r"))),
			).Block(
				Id("last").Op("=").Int64().Params(Len(Id("r"))),
			),

			Return(Id("r").Index(Id("first").Op(":").Id("last")), Nil()),
		)
}

func (g *CustomGenerator) generateAny(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("Any").
		Params(
			Id("f").Id("func(*"+typeName+") bool"),
		).
		Params(Bool()).
		Block(
			Id("first").Op(":=").Id("r.FirstOrDefault").Call(Id("f")),

			Return(Id("first").Op("!=").Nil()),
		)
}

func (g *CustomGenerator) generateEqualToFillManually(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id("*" + typeName),
		).Id("equal").
		Params(
			Id("another").Id("*" + typeName),
		).
		Bool().
		Block(
			Comment("`equal` method has to be implemented manually"),
		)
}

func (g *CustomGenerator) generateEqualImplementation(f *File, typeName string) {
	f.Func().
		Params(Id("r").Id("*"+typeName)).
		Id("Equal").
		Params(
			Id("another").Qual("github.com/doctornick42/gosli/lib", "Equaler"),
		).
		Params(
			Bool(),
			Error(),
		).
		Block(
			Id("anotherCasted, ok").Op(":=").Id("another").Dot(fmt.Sprintf("(*%s)", typeName)),

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
			Id("r").Id(g.getStructName(typeName)),
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
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("Contains").
		Params(
			Id("el").Id("*"+typeName),
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
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("processSliceOperation").
		Params(
			Id("sl2").Id(g.getStructName(typeName)),
			Id("f").Func().Params(
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
			).Params(
				Index().Qual("github.com/doctornick42/gosli/lib", "Equaler"),
				Error(),
			),
		).
		Params(
			Id("[]*"+typeName),
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

			Id("res").Op(":=").Make(Index().Id("*"+typeName), Len(Id("untypedRes"))),

			For(
				Id("i").Op(":=").Range().Id("untypedRes"),
			).Block(
				Id("res[i]").Op("=").Id("untypedRes[i]").Dot(fmt.Sprintf("(*%s)", typeName)),
			),

			Return(Id("res"), Nil()),
		)
}

func (g *CustomGenerator) generateGetUnion(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("GetUnion").
		Params(
			Id("sl2").Index().Id("*"+typeName),
		).
		Params(
			Index().Id("*"+typeName),
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
			Id("r").Id(g.getStructName(typeName)),
		).
		Id("InFirstOnly").
		Params(
			Id("sl2").Index().Id("*"+typeName),
		).
		Params(
			Index().Id("*"+typeName),
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

func (g *CustomGenerator) getStructName(typeName string) string {
	return typeName + "Slice"
}

func (g *CustomGenerator) generateInfrastructure(f *File, typeName string) {
	structName := g.getStructName(typeName)

	f.Type().Id(structName).Index().Id("*" + typeName)
}
