package gen_code

import (
	"go/format"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	errpkg "shub_go/src/packages/err"
	"strings"
)

func build() (*WrapStruct, *WrapPublicField) {
	var tree map[string]map[string]errpkg.Error

	f, err := os.Open("error.yaml")

	if err != nil {
		log.Fatalln("fail open file", err)
	}

	defer func() {
		f.Close()
	}()

	err = yaml.NewDecoder(f).Decode(&tree)

	if err != nil {
		log.Fatalln("err parse tree data", err)
	}

	var wrapFields []*WrapField
	var rootErrorFields []*Field
	var publicFields []*PublicField

	for key, value := range tree {
		wrapFieldElement := WrapField{
			Name: key,
		}
		for name, _ := range value {
			field := Field{
				Name:   strings.Title(name),
				Type:   "*Error",
				YmlTag: name,
			}
			wrapFieldElement.Fields = append(wrapFieldElement.Fields, &field)
		}
		wrapFields = append(wrapFields, &wrapFieldElement)

		rootErrorField := Field{
			Name:   strings.Title(key),
			Type:   "*" + key,
			YmlTag: key,
		}
		rootErrorFields = append(rootErrorFields, &rootErrorField)

		publicField := PublicField{
			Name: strings.Title(key),
			Type: "*" + key,
		}
		publicFields = append(publicFields, &publicField)
	}

	wrapFields = append(wrapFields, &WrapField{
		Name:   "rootErr",
		Fields: rootErrorFields,
	})
	result := WrapStruct{
		PackageName: "errpkg",
		Elements:    wrapFields,
	}

	return &result, &WrapPublicField{
		Elements: publicFields,
	}
}

func genLoad(loads *WrapPublicField) {
	loadFile := "src/packages/err/load.go"

	f, err := os.Create(loadFile)

	if err != nil {
		panic("err create file")
	}

	templ := template.Must(template.ParseFiles("src/template/error_code/load.tmpl"))
	templ.Execute(f, loads)
	f.Close()

	//goFmt(loadFile)
}

func genStruct(structs *WrapStruct) {
	structFile := "src/packages/err/struct.go"

	f, err := os.Create(structFile)

	if err != nil {
		panic("err create file")
	}

	templ := template.Must(template.ParseFiles("src/template/error_code/struct.tmpl"))
	templ.Execute(f, structs)
	f.Close()

	goFmt(structFile)

}

func goFmt(path string) {
	read, _ := ioutil.ReadFile(path)
	content, _ := format.Source(read)
	ioutil.WriteFile(path, []byte(content), 0)
}

func GenCode() {
	structs, loads := build()

	genStruct(structs)
	genLoad(loads)
}
