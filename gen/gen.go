package gen

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	. "github.com/dave/jennifer/jen"
)

func Run(args []string) error {
	if len(args) < 2 {
		return errors.New("Wrong amount of arguments")
	}

	originFilePath := args[0]
	typeName := args[1]

	moduleName, err := getModuleName(originFilePath)
	if err != nil {
		return err
	}

	log.Printf("Module name: %s", moduleName)

	f := NewFile(moduleName)
	generateInfrastructure(f, typeName)
	generateFirstOrDefault(f, typeName)
	generateFirst(f, typeName)
	generateWhere(f, typeName)
	generateSelect(f, typeName)
	generatePage(f, typeName)
	generateSliceToEqualers(f, typeName)
	generateContains(f, typeName)
	generateProcessSliceOperation(f, typeName)
	generateGetUnion(f, typeName)
	generateInFirstOnly(f, typeName)
	generateEqualImplementation(f, typeName)

	genFileName := getGeneratedFileName(originFilePath, typeName)

	log.Printf("Generated filename: %s", genFileName)
	err = f.Save(genFileName)
	if err != nil {
		return err
	}

	genFileName = getEqualGeneratedFileName(originFilePath, typeName)
	if _, err := os.Stat(genFileName); os.IsNotExist(err) {
		f = NewFile(moduleName)
		generateEqualToFillManually(f, typeName)

		log.Printf("Generated filename: %s", genFileName)
		return f.Save(genFileName)
	}

	return nil
}

func getModuleName(originFilePath string) (string, error) {
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

func getGeneratedFileName(originFilePath, typeName string) string {
	return generateFileName(originFilePath, "generated", typeName)
}

func getEqualGeneratedFileName(originFilePath, typeName string) string {
	return generateFileName(originFilePath, "equal", typeName)
}

func generateFileName(originFilePath, suffix, typeName string) string {
	splitted := strings.Split(originFilePath, "/")

	shortFileName := splitted[len(splitted)-1]
	generatedName := strings.ToLower(typeName) + "_" + suffix + ".go"

	return strings.Replace(originFilePath, shortFileName, generatedName, 1)
}

func generateFirstOrDefault(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func generateFirst(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func generateWhere(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func generateSelect(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func generatePage(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func generateEqualToFillManually(f *File, typeName string) {
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

func generateEqualImplementation(f *File, typeName string) {
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

func generateSliceToEqualers(f *File, typeName string) {
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

func generateContains(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func generateProcessSliceOperation(f *File, typeName string) {
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

func generateGetUnion(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func generateInFirstOnly(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
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

func firstRuneToLower(origin string) string {
	runes := []rune(origin)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func getStructName(typeName string) string {
	return typeName + "Slice"
}

func generateInfrastructure(f *File, typeName string) {
	structName := getStructName(typeName)

	f.Type().Id(structName).Index().Id("*" + typeName)
}
