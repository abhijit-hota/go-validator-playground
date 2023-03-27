package pkg

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"reflect"
	"time"
)

func BuildStructTypeFromAST(structTypeSpec *ast.TypeSpec) (reflect.Type, error) {
	structType, ok := structTypeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("%v is not a struct type", structTypeSpec.Name.Name)
	}

	var structFields []reflect.StructField
	for _, field := range structType.Fields.List {
		fieldName := field.Names[0].Name
		fieldType, err := convertExprToType(field.Type)
		if err != nil {
			return nil, fmt.Errorf("error creating field %v: %v", fieldName, err)
		}

		structField := reflect.StructField{
			Name: fieldName,
			Type: fieldType,
			Tag:  reflect.StructTag(field.Tag.Value),
		}
		structFields = append(structFields, structField)
	}

	return reflect.StructOf(structFields), nil
}

func convertExprToType(expr ast.Expr) (reflect.Type, error) {
	switch t := expr.(type) {

	case *ast.Ident:
		return getTypeFromIdent(t)

	case *ast.SelectorExpr:
		return getTypeFromSelectorExpr(t)

	case *ast.ArrayType:
		elemType, err := convertExprToType(t.Elt)
		if err != nil {
			return nil, err
		}
		return reflect.SliceOf(elemType), nil

	case *ast.StarExpr:
		innerType, err := convertExprToType(t.X)
		if err != nil {
			return nil, err
		}
		return reflect.PtrTo(innerType), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", expr)
	}
}

func getTypeFromIdent(ident *ast.Ident) (reflect.Type, error) {
	switch ident.Name {
	case "bool":
		return reflect.TypeOf(bool(false)), nil
	case "int":
		return reflect.TypeOf(int(0)), nil
	case "int8":
		return reflect.TypeOf(int8(0)), nil
	case "int16":
		return reflect.TypeOf(int16(0)), nil
	case "int32":
		return reflect.TypeOf(int32(0)), nil
	case "int64":
		return reflect.TypeOf(int64(0)), nil
	case "uint":
		return reflect.TypeOf(uint(0)), nil
	case "uint8":
		return reflect.TypeOf(uint8(0)), nil
	case "uint16":
		return reflect.TypeOf(uint16(0)), nil
	case "uint32":
		return reflect.TypeOf(uint32(0)), nil
	case "uint64":
		return reflect.TypeOf(uint64(0)), nil
	case "uintptr":
		return reflect.TypeOf(uintptr(0)), nil
	case "float32":
		return reflect.TypeOf(float32(0)), nil
	case "float64":
		return reflect.TypeOf(float64(0)), nil
	case "string":
		return reflect.TypeOf(""), nil
	default:
		return nil, fmt.Errorf("unknown type: %v", ident.Name)
	}
}

func getTypeFromSelectorExpr(selector *ast.SelectorExpr) (reflect.Type, error) {
	pkgName, typeName := selector.X.(*ast.Ident).Name, selector.Sel.Name
	if pkgName != "time" {
		return nil, fmt.Errorf("unsupported package: %v", pkgName)
	}
	switch typeName {
	case "Time":
		return reflect.TypeOf(time.Now()), nil
	case "Duration":
		return reflect.TypeOf(time.Duration(0)), nil
	default:
		return nil, fmt.Errorf("unsupported time type: %v", typeName)
	}
}

func BuildStructFromJSON(structType reflect.Type, jsonData []byte) (any, error) {
	newStruct := reflect.New(structType).Elem()

	err := json.Unmarshal(jsonData, newStruct.Addr().Interface())
	if err != nil {
		return nil, err
	}

	return newStruct.Interface(), nil
}
