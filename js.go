package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall/js"

	"abhijithota.me/go-validator-playground/pkg"
)

type JSResult struct {
	Error         error              `json:"error"`
	Status        string             `json:"status"`
	InvalidResult *pkg.InvalidResult `json:"invalid_result"`
}

func (r *JSResult) toJSON() js.Value {
	if r.Error != nil {
		r.Status = "error"
	}

	if r.InvalidResult != nil {
		r.Status = "invalid"
	}

	if (r.Error == nil && r.InvalidResult == nil) || r.Status == "" {
		r.Status = "valid"
	}

	jsonBytes, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	obj := make(map[string]interface{})

	if err := json.Unmarshal(jsonBytes, &obj); err != nil {
		panic(err)
	}

	if r.Error != nil {
		obj["error"] = r.Error.Error()
	}

	return js.ValueOf(obj)
}

func validateStruct_JS(_ js.Value, args []js.Value) any {
	res := JSResult{Status: "valid"}

	structTypeRepr, jsonData, err := validateArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't validate JS args: %v", err)
		res.Error = err
		return res.toJSON()
	}

	val, err := pkg.ValidateJSONAgainstStructRepr(structTypeRepr, []byte(jsonData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed: %v", err)
		res.Error = err
		return res.toJSON()
	}

	if val != nil {
		res.InvalidResult = val
		return res.toJSON()
	}

	fmt.Println("Struct valid")

	return res.toJSON()
}

func validateArgs(args []js.Value) (string, string, error) {
	if len(args) != 2 {
		return "", "", fmt.Errorf("expected 2 arguments, got %d", len(args))
	}
	if len(args) != 2 {
		return "", "", fmt.Errorf("expected 2 arguments, got %d", len(args))
	}

	structTypeRepr := args[0]
	jsonData := args[1]

	// check if both args are non-empty strings
	if structTypeRepr.Type() != js.TypeString || jsonData.Type() != js.TypeString {
		return "", "", fmt.Errorf("expected string arguments, got %s and %s", structTypeRepr.Type(), jsonData.Type())
	}

	if structTypeRepr.String() == "" || jsonData.String() == "" {
		return "", "", fmt.Errorf("expected non-empty string arguments, got %s and %s", structTypeRepr.String(), jsonData.String())
	}

	return structTypeRepr.String(), jsonData.String(), nil
}
