package pkg

func ValidateJSONAgainstStructRepr(structTypeRep string, jsonData []byte) (*InvalidResult, error) {
	structTypeSpec, err := ParseStructReprToAST(structTypeRep)
	if err != nil {
		return nil, err
	}

	structType, err := BuildStructTypeFromAST(structTypeSpec)
	if err != nil {
		return nil, err
	}

	newStruct, err := BuildStructFromJSON(structType, jsonData)
	if err != nil {
		return nil, err
	}

	return ValidateStruct(newStruct)
}
