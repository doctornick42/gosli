package gen

import (
	"fmt"
	"strings"
	"unicode"

	. "github.com/dave/jennifer/jen"
)

func firstRuneToLower(origin string) string {
	return modifyFirstRune(origin, unicode.ToLower)
}

func firstRuneToUpper(origin string) string {
	return modifyFirstRune(origin, unicode.ToUpper)
}

func modifyFirstRune(origin string, f func(rune) rune) string {
	runes := []rune(origin)
	runes[0] = f(runes[0])
	return string(runes)
}

func getStructName(typeName string) string {
	if string(typeName[0]) == "*" {
		typeName = strings.TrimPrefix(typeName, "*") + "Pointer"
	}

	typeName = firstRuneToUpper(typeName)

	return typeName + "Slice"
}

func generateInfrastructure(f *File, typeName string) {
	structName := getStructName(typeName)

	f.Type().Id(structName).Index().Id(typeName)
}

func generateFileName(originFilePath, suffix, typeName string) string {
	splitted := strings.Split(originFilePath, "/")

	shortFileName := splitted[len(splitted)-1]
	generatedName := strings.ToLower(typeName)
	if len(suffix) > 0 {
		suffix = "_" + suffix
	}

	generatedName = generatedName + suffix + ".go"

	return strings.Replace(originFilePath, shortFileName, generatedName, 1)
}

func generateFirst(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("First").
		Params(
			Id("f").Id("func("+typeName+") bool"),
		).
		Params(
			Id(typeName),
			Error(),
		).
		Block(
			For(
				Id("_, slEl").Op(":=").Range().Id("r").Block(
					If(
						Id("f").Call(Id("slEl")),
					).Block(
						Return(
							Id("slEl"),
							Nil(),
						),
					),
				),
			),
			Var().Id("defVal").Id(typeName),
			Return(Id("defVal"), Qual("errors", "New").Call(Lit("Not found"))),
		)
}

func generateFirstOrDefault(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("FirstOrDefault").
		Params(
			Id("f").Id("func("+typeName+") bool"),
		).
		Id(typeName).
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

func generateWhere(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("Where").
		Params(
			Id("f").Id("func("+typeName+") bool"),
		).
		Id("[]"+typeName).
		Block(
			Id("res").Op(":=").Make(Id("[]"+typeName), Lit(0)),

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
			Id("f").Id(fmt.Sprintf("func(%s) interface{}", typeName)),
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
			Id("[]"+typeName),
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
					Index().Id(typeName).Block(), Nil(),
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

func generateAny(f *File, typeName string) {
	f.Func().
		Params(
			Id("r").Id(getStructName(typeName)),
		).
		Id("Any").
		Params(
			Id("f").Id("func("+typeName+") bool"),
		).
		Params(Bool()).
		Block(
			Id("_, err").Op(":=").Id("r.First").Call(Id("f")),

			Return(Id("err").Op("==").Nil()),
		)
}
