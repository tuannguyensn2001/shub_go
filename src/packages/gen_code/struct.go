package gen_code

type Field struct {
	Name   string
	Type   string
	YmlTag string
}

type WrapField struct {
	Name   string
	Fields []*Field
}

type WrapStruct struct {
	PackageName string
	Elements    []*WrapField
}

type PublicField struct {
	Name string
	Type string
}

type WrapPublicField struct {
	Elements []*PublicField
}
