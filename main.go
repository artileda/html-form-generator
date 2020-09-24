package main

import (
	"os"
	"strings"

	"github.com/ywnn/html-form-generator/model"
)

func main() {
	var base = model.Form{
		Action:  "",
		Method:  "",
		EncType: "",
		Datum:   []model.Input{},
	}

	if len(os.Args) < 1 {
		panic(1)
	}
	params := os.Args[1:]
	for _, param := range params {
		paramHash := strings.Split(param, ":")
		if len(paramHash) == 2 {
			base.Datum = append(base.Datum, model.Input{
				Name: paramHash[0],
				Type: paramHash[1],
			})
		}
	}
	base.Generate()

}
