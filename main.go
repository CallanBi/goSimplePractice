package main

import (
	"fmt"

	"goSimplePractice/forcharacter"
	"goSimplePractice/package1"
	"goSimplePractice/runelength"
	"goSimplePractice/structtest"
	"goSimplePractice/varargs"
)

func main() {
	package1.Pkg1ExportFunc()
	forcharacter.PrintCharacter()
	runelength.CalRuneLength()

	fmt.Print("\n\n变长变量练习：\n")

	varargs.PrintVarArgsUsingStruct(&varargs.Options{Par1: "aaa", Par2: 123, Par3: "bbb"})
	varargs.PrintVarArgsUsingEmptyInterface("aaa", 123, "bbb")

	structtest.Test()

}
