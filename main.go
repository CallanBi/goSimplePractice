package main

import (
	"fmt"

	"goSimplePractice/forcharacter"
	"goSimplePractice/package1"
	"goSimplePractice/returnfunc"
	"goSimplePractice/runelength"
	"goSimplePractice/structtest"
	"goSimplePractice/varargs"

	chan_practice "goSimplePractice/channel"
	"goSimplePractice/coroutine"
	panic_practice "goSimplePractice/panic"
	slice_and_arr_practice "goSimplePractice/slice_and_arr"
	string_practice "goSimplePractice/string"
)

func main() {
	package1.Pkg1ExportFunc()
	forcharacter.PrintCharacter()
	runelength.CalRuneLength()

	fmt.Print("\n\n变长变量练习：\n")

	varargs.PrintVarArgsUsingStruct(&varargs.Options{Par1: "aaa", Par2: 123, Par3: "bbb"})
	varargs.PrintVarArgsUsingEmptyInterface("aaa", 123, "bbb")

	structtest.Test()

	fmt.Print("\n\n返回 defer 的匿名函数练习：\n")

	fmt.Println(returnfunc.DeferReturnFunc())

	coroutine.Coroutine()

	panic_practice.PanicPractice()

	string_practice.StringPractice()

	slice_and_arr_practice.SliceAndArrPractice()

	chan_practice.ChanPractice()

}
