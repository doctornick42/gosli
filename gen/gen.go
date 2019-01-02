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
	generateFirstOrDefault(f, typeName)
	generateFirst(f, typeName)
	generateWhere(f, typeName)
	generateEqualImplementation(f, typeName)
	generateSelect(f, typeName)
	generateSliceToEqualers(f, typeName)
	generateSliceToInterfacesSlice(f, typeName)
	generateContains(f, typeName)
	generateProcessSliceOperation(f, typeName)
	generateGetUnion(f, typeName)
	generateInFirstOnly(f, typeName)

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
	r := bufio.NewReader(f)
	firstLine, err := r.ReadString('\n')
	if err != nil {
		return "", err
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
	f.Func().Id(typeName+"FirstOrDefault").
		Params(
			Id("sl").Id("[]*"+typeName),
			Id("f").Id("func(*"+typeName+") bool"),
		).
		Id("*"+typeName).
		Block(
			For(
				Id("_, slEl").Op(":=").Range().Id("sl").Block(
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
	f.Func().Id(typeName+"First").
		Params(
			Id("sl").Id("[]*"+typeName),
			Id("f").Id("func(*"+typeName+") bool"),
		).
		Params(Id("*"+typeName), Error()).
		Block(
			Id("first").Op(":=").Id(typeName+"FirstOrDefault").Call(Id("sl"), Id("f")),
			If(
				Id("first").Op("==").Nil(),
			).Block(
				Return(Nil(), Qual("errors", "New").Call(Lit("Not found"))),
			),
			Return(Id("first"), Nil()),
		)
}

func generateWhere(f *File, typeName string) {
	f.Func().Id(typeName+"Where").
		Params(
			Id("sl").Id("[]*"+typeName),
			Id("f").Id("func(*"+typeName+") bool"),
		).
		Id("[]*"+typeName).
		Block(
			Id("res").Op(":=").Make(Id("[]*"+typeName), Lit(0)),

			For(
				Id("_, slEl").Op(":=").Range().Id("sl").Block(
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
	f.Func().Id(typeName+"Select").
		Params(
			Id("sl").Id("[]*"+typeName),
			Id("f").Id(fmt.Sprintf("func(*%s) interface{}", typeName)),
		).
		Id("[]interface{}").
		Block(
			Id("res").Op(":=").Make(Id("[]interface{}"), Len(Id("sl"))),

			For(
				Id("i").Op(":=").Range().Id("sl").Block(
					Id("res").Index(Id("i")).Op("=").
						Id("f").Call(Id("sl").Index(Id("i"))),
				),
			),
			Return(Id("res")),
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
		Id(typeName+"SliceToEqualers").
		Params(
			Id("sl").Id("[]*"+typeName),
		).
		Index().Qual("github.com/doctornick42/gosli/lib", "Equaler").
		Block(
			Id("equalerSl").Op(":=").Make(Id("[]lib.Equaler"), Len(Id("sl"))),

			For(
				Id("i").Op(":=").Range().Id("sl"),
			).Block(
				Id("equalerSl[i]").Op("=").Id("sl[i]"),
			),

			Return(Id("equalerSl")),
		)
}

func generateSliceToInterfacesSlice(f *File, typeName string) {
	f.Func().
		Id(typeName+"SliceToInterfacesSlice").
		Params(
			Id("sl").Id("[]*"+typeName),
		).
		Id("[]interface{}").
		Block(
			Id("equalerSl").Op(":=").Make(Id("[]interface{}"), Len(Id("sl"))),

			For(
				Id("i").Op(":=").Range().Id("sl"),
			).Block(
				Id("equalerSl[i]").Op("=").Id("sl[i]"),
			),

			Return(Id("equalerSl")),
		)
}

func generateContains(f *File, typeName string) {
	f.Func().
		Id(typeName+"Contains").
		Params(
			Id("sl").Id("[]*"+typeName),
			Id("el").Id("*"+typeName),
		).
		Params(
			Bool(),
			Error(),
		).
		Block(
			Id("equalerSl").Op(":=").Id(typeName+"SliceToEqualers").Call(Id("sl")),
			Return(
				Qual("github.com/doctornick42/gosli/lib", "Contains").
					Call(Id("equalerSl"), Id("el")),
			),
		)
}

func generateProcessSliceOperation(f *File, typeName string) {
	f.Func().
		Id(typeName+"ProcessSliceOperation").
		Params(
			Id("sl1, sl2").Id("[]*"+typeName),
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
			Id("equalerSl1").Op(":=").Id(typeName+"SliceToEqualers").Call(Id("sl1")),
			Id("equalerSl2").Op(":=").Id(typeName+"SliceToEqualers").Call(Id("sl2")),
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
		Id(typeName+"GetUnion").
		Params(
			Id("sl1, sl2").Index().Id("*"+typeName),
		).
		Params(
			Index().Id("*"+typeName),
			Error(),
		).
		Block(
			Return(
				Id(typeName+"ProcessSliceOperation").Call(
					Id("sl1, sl2"),
					Qual("github.com/doctornick42/gosli/lib", "GetUnion"),
				),
			),
		)
}

func generateInFirstOnly(f *File, typeName string) {
	f.Func().
		Id(typeName+"InFirstOnly").
		Params(
			Id("sl1, sl2").Index().Id("*"+typeName),
		).
		Params(
			Index().Id("*"+typeName),
			Error(),
		).
		Block(
			Return(
				Id(typeName+"ProcessSliceOperation").Call(
					Id("sl1, sl2"),
					Qual("github.com/doctornick42/gosli/lib", "InFirstOnly"),
				),
			),
		)
}
