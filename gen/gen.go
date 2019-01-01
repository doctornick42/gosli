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
	fmt.Printf("%#v\r\n", f)

	genFileName := getGeneratedFileName(originFilePath)
	log.Printf("Generated filename: %s", genFileName)
	return f.Save(genFileName)
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

func getGeneratedFileName(originFilePath string) string {
	splitted := strings.Split(originFilePath, "/")

	shortFileName := splitted[len(splitted)-1]
	withoutExtension := strings.Split(shortFileName, ".")[0]

	generatedName := withoutExtension + "_generated.go"

	return strings.Replace(originFilePath, shortFileName, generatedName, 1)
}

func generateFirstOrDefault(f *File, typeName string) {
	f.Func().Id("FirstOrDefault").
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
