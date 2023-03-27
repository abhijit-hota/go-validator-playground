package main

import (
	"syscall/js"
)

func main() {
	ch := make(chan bool)
	js.Global().Set("validateStruct", js.FuncOf(validateStruct_JS))
	<-ch
}
