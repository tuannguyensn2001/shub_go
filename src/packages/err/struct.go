// Code generated by gen-error-code tool. DO NOT EDIT.

package errpkg

type general struct {
	Success   *Error `yaml:"success"`
	NotFound  *Error `yaml:"notFound"`
	Forbidden *Error `yaml:"forbidden"`
}

type auth struct {
	NotFound *Error `yaml:"notFound"`
	Invalid  *Error `yaml:"invalid"`
}

type rootErr struct {
	General *general `yaml:"general"`
	Auth    *auth    `yaml:"auth"`
}
