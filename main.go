package main

import (
	"github.com/ywnn/html-form-generator/model"
)

func main() {
	//var b []Input
	var a = model.Form{
		Action:  "",
		Method:  "POST",
		EncType: "",
		Datum:   []model.Input{{Name: "password", Type: "password"}},
	}

	a.Generate()
}
