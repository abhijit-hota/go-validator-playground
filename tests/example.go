package main

import (
	"fmt"

	"abhijithota.me/go-validator-playground/pkg"
)

// WIP
func Demo() {
	demoType := `
        type Demo struct {
            Foo int32 ` + "`json:\"foo\" validate:\"min=1,max=15\"`" + `
            Bar string ` + "`json:\"bar\"`" + `
        }`

	demoJSON := `{"foo": 42, "bar": "hello"}`

	val, err := pkg.ValidateJSONAgainstStructRepr(demoType, []byte(demoJSON))
	if err != nil {
		panic(err)
	}

	if val != nil {
		fmt.Printf("Invalid fields: %v\n", val.Fields)
		fmt.Printf("Invalid tags: %v\n", val.Tags)
		return
	}

	fmt.Println("Struct valid")
}
