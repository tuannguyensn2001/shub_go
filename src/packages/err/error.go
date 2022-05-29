package errpkg

type Error struct {
	Status  int32  `yaml:"status" json:"-"`
	Message string `yaml:"message" json:"message"`
	Code    string `yaml:"code" json:"code"`
}

func (e *Error) Error() string {
	return e.Message
}
