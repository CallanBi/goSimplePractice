package runelength

import (
	"fmt"
	"runtime"
	"unicode/utf8"
)

func CalRuneLength() {
	fmt.Print("\ncalRuneLength 练习:\n\n")
	a := `asSASA ddd dsjkdsjs dk`
	b := `asSASA ddd dsjkdsjsこん dk`
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("a 字节数", len(a)) // 统计字节，带空格
	fmt.Println("b 字节数", len(b)) // 统计字节，带空格
	fmt.Println("a 字符数", utf8.RuneCountInString(a))
	fmt.Println("b 字符数", utf8.RuneCountInString(b))

	fmt.Println("运行时", runtime.GOOS)
}
