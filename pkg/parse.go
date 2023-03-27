package pkg

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseStructReprToAST(structTypeRep string) (*ast.TypeSpec, error) {
	fset := token.NewFileSet()
	val, err := parser.ParseFile(
		fset, "",
		fmt.Sprintf("package _\n%v", structTypeRep),
		parser.AllErrors,
	)

	if err != nil {
		return nil, err
	}

	// Get the first type declaration in the string
	var obj *ast.Object
	for _, v := range val.Scope.Objects {
		if v.Kind == ast.Typ {
			obj = v
			break
		}
	}
	if obj == nil {
		return nil, fmt.Errorf("no type declaration found")
	}

	structTypeSpec := obj.Decl.(*ast.TypeSpec)

	return structTypeSpec, nil
}
