package structtest

import (
	"fmt"

	"goSimplePractice/varargs"
)

func Test() {
	fmt.Print("\n\nstructTest: ")

	varMadeFromType := &varargs.Options{Par1: "123", Par2: 19980416, Par3: "testTestTest"}
	varMadeFromStruct := &struct {
		Par1 string
		Par2 int64
		Par3 interface{}
	}{Par1: "123", Par2: 19980416, Par3: "testTestTest"}

	varargs.PrintVarArgsUsingStruct(varMadeFromType)

	// varargs.PrintVarArgsUsingStruct(varMadeFromStruct)
	// Output: Cannot use 'varMadeFromStruct' (type *struct {...}) as the type *Options

	varargs.PrintVarArgsUsingStruct((*varargs.Options)(varMadeFromStruct)) // ✅
}
