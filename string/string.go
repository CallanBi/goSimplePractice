package string_practice

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func StringPractice() {
	fmt.Println("\n\nstring_practice 练习：")

	aString := "asSASA ddd dsjkdsjs dk 中文字符"

	// strings.Clone
	b := strings.Clone(aString)

	fmt.Println("aString: ", aString)

	fmt.Println("b: ", b)

	// strings.Contains
	fmt.Println(strings.Contains(aString, "SA")) // true

	// strings.ContainsAny
	fmt.Println(strings.ContainsAny(aString, "SAW")) // 在 s 中有没有找到 SAW 的任何一个字符 , trues

	// strings.ContainsRune
	// Finds whether a string contains a particular Unicode code point.
	// The code point for the lowercase letter "a", for example, is 97.
	fmt.Println(strings.ContainsRune("aardvark", 97)) // true

	// strings.Count
	fmt.Println(strings.Count(aString, "s")) // 4

	// strings.Cut
	countPractice()

	// strings.Fields
	fmt.Printf("Fields are: %q/n", strings.Fields("  foo bar  baz   ")) // ["foo" "bar" "baz"]

	// strings.FieldsFunc
	fieldsFuncPractice()

	// strings.HasPrefix
	fmt.Println(strings.HasPrefix(aString, "as")) // true

	// strings.HasSuffix
	fmt.Println(strings.HasSuffix(aString, "字符")) // true

	// strings.Index
	fmt.Println(strings.Index(aString, "中文")) // 23

	// strings.IndexAny
	fmt.Println(strings.IndexAny(aString, "SAW")) // 2

	// strings.IndexFunc
	fmt.Println(strings.IndexFunc(aString, func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	})) // 23

	// strings.Join
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-")) // a-b-c

	// strings.LastIndex
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3

	// strings.LastIndexAny
	fmt.Println(strings.LastIndexAny("go gopher", "god")) // 4

	// strings.Map
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "gopher123")) // hpqifs234

	// strings.Repeat
	fmt.Println(strings.Repeat("go", 3)) // gogogo

	// strings.Replace
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) // oinky oinky oink

	// strings.ReplaceAll
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) // moo moo moo

	// strings.spilt
	fmt.Println(strings.Split("a,b,c", ",")) // ["a" "b" "c"]

	// strings.SplitAfter
	fmt.Println(strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]

	// strings.SplitAfterN
	fmt.Println(strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]

	// strings.SplitN
	fmt.Println(strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]

	// strings.ToLower
	fmt.Println(strings.ToLower("ABC")) // abc

	// strings.ToLowerSpecial
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "ÖÜÄ")) // öüä

	// strings.ToTitle
	fmt.Println(strings.ToTitle("хлеб")) // ХЛЕБ

	// strings.ToTitleSpecial
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "öüä")) // ÖÜÄ

	// strings.ToUpper
	fmt.Println(strings.ToUpper("abc")) // ABC

	// strings.Trim
	fmt.Println(strings.Trim("!!!Abe!!!", "!")) // Abe
	fmt.Println(strings.Trim("  水电费    ", " ")) // 水电费

	// strings.TrimFunc
	fmt.Println(strings.TrimFunc(" !!!Abe!!!  ", func(r rune) bool {
		return !unicode.IsLetter(r)
	})) // Abe

	// strings.TrimPrefix
	fmt.Println(strings.TrimPrefix("!!!Abe!!!  ", "!")) // !!Abe!!!

	// strings.TrimSpace
	fmt.Println(strings.TrimSpace("  \t\n a lone gopher \n\t\r\n")) // a lone gopher

	// strings.Builder
	builderPractice()

	fmt.Print("\n\nstring_practice 练习结束\n\n")

}

func countPractice() {
	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")
}

func fieldsFuncPractice() {
	// f := func(c rune) bool {
	// 	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	// }

	fmt.Printf("Fields are: %q/n", strings.FieldsFunc("  foo1;bar2,az3...", func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})) // ["foo1" "bar2" "az3"]
}

func builderPractice() {
	// 创建一个 strings.Builder
	var b strings.Builder
	b.Grow(100)
	fmt.Printf("len: %d, cap: %d\n", b.Len(), b.Cap())

	// 写入字符串
	b.WriteString("hello, ")
	fmt.Printf("len: %d, cap: %d\n", b.Len(), b.Cap())

	// 写入字节
	b.Write([]byte("world!"))
	fmt.Printf("len: %d, cap: %d\n", b.Len(), b.Cap())

	// 写入字符
	b.WriteRune('!')
	fmt.Printf("len: %d, cap: %d\n", b.Len(), b.Cap())

	// 输出
	fmt.Println(b.String())

	var anotherB strings.Builder

	anotherB.WriteString("another builder test")

	fmt.Println(anotherB.String())

	// strconv.Atoi
	i, _ := strconv.Atoi("-42") // -42
	s := strconv.Itoa(i)        // "-42"

	boolVal, _ := strconv.ParseBool("true")    // true
	f, _ := strconv.ParseFloat("3.1415", 64)   // 3.1415
	iVal, _ := strconv.ParseInt("-42", 10, 64) // -42
	u, _ := strconv.ParseUint("42", 10, 64)    // 42

	bs := strconv.FormatBool(true)                 // "true"
	fs := strconv.FormatFloat(3.1415, 'f', -1, 64) // "3.1415"
	is := strconv.FormatInt(42, 2)                 // "101010"
	us := strconv.FormatUint(42, 16)               // "2a"

	q := strconv.Quote(`Hello,
	世界`) // "Hello,\n\t世界"

	fmt.Println(i, s, boolVal, f, iVal, u, bs, fs, is, us, q)

}
