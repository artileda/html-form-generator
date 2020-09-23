package model

type Input struct {
	Name string
	Type string
}

type Form struct {
	Action  *string
	Method  *string
	EncType *string
	Datum   []Input
}

func (f Form) Generate() error {
	return nil
}
