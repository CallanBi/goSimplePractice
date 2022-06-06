package slice_and_arr_practice

import (
	"fmt"
)

func printLenCap[T string | int](msg string, sli []T) {
	fmt.Printf("%v len=%d cap=%d %v\n", msg, len(sli), cap(sli), sli)
}

func SliceAndArrPractice() {
	fmt.Println("\n\nslice_and_arr_practice 练习：")

	// 在 go 中，数组属于值类型，当一个数组被赋值或传递时，实际上会创建一个新的数组，并将原数组的内容复制到新数组中。
	a := [...]int{1, 2, 3}

	b := a

	a[0] = 100

	fmt.Println(a, b) // [100 2 3] [1 2 3]

	// 为了避免复制数组，一般会传递指向数组的指针

	square := func(arr *[3]int) {
		for idx, val := range *arr {
			(*arr)[idx] = val * val
		}
	}

	square(&a)

	fmt.Println(a) // [10000 4 9]

	// 切片可以指定长度
	sliWizLen := make([]int, 3)        // 初始化为 [0 0 0]
	anotherWizLen := make([]int, 0, 5) // 初始化为 []，初始化的 cap 为 5
	langs := []string{"typescript", "javascript", "python", "go"}
	// 当超出切片初始化长度时，会自动分配内存
	// 容量在比较小的时候，一般是以 2 的倍数扩大的，例如 2 4 8 16 …，
	// 当达到 2048 时，会采取新的策略，避免申请内存过大，导致浪费
	sliWizLen = append(sliWizLen, 1, 2, 3)

	fmt.Println(sliWizLen, anotherWizLen, langs)

	// type MySlice[T int | string] []T

	// 切片本质是一个数组片段的描述，包括了数组的指针，这个片段的长度和容量(不改变内存分配情况下的最大长度)
	// 切片操作并不复制切片指向的元素，创建一个新的切片会复用原来切片的底层数组，因此切片操作是非常高效的

	// 任何泛型类型都必须传入类型实参实例化才可以使用, 比如下面的 func(sli MySlice[string]) 不能写成 func(sli MySlice)
	// printLenCap := func(msg string, sli MySlice[string]) {
	// 	fmt.Printf("%v len=%d cap=%d %v\n", msg, len(sli), cap(sli), sli)
	// }

	langs2 := langs[2:4] // 截取尾部可以超过 len(langs) 的值, langs2 和 langs 指向同一个数组
	printLenCap("langs2: ", langs2)
	langs2[0] = "c++"
	printLenCap("langs2:", langs2)
	printLenCap("langs: ", langs) // len=4 cap=4 [typescript javascript c++ go]

	langs = append(langs, "rust") // langs 被 append 重新赋值后，langs 和 langs2 指向不同的数组
	langs[2] = "python"
	printLenCap("langs: ", langs)  // len=5 cap=8 [typescript javascript python go rust]
	printLenCap("langs2:", langs2) // len=2 cap=2 [c++ go]

	// copy 切片的几种操作
	// copy 操作
	// 1
	slice1 := []int{1, 2, 3, 4, 5}
	printLenCap("slice1: ", slice1)
	slice2 := make([]int, 3)
	copy(slice2, slice1) // 将 slice1 的前 3 个元素复制到 slice2 中
	fmt.Println(slice2)  // [1 2 3]

	// 2
	slice2 = append([]int(nil), slice1...) // 将 slice1 的所有元素复制到 slice2 中
	printLenCap("slice2: ", slice2)        // len=5 cap=6 [1 2 3 4 5]

	// 3
	slice2 = append(slice2[:0:0], slice1...) // 将 slice1 的所有元素复制到 slice2 中
	// s = s[low : high : max]
	// 切片的三个参数的切片截取的意义为 low 为截取的起始下标（含），
	// high 为窃取的结束下标（不含 high），
	// max 为切片保留的原切片的最大下标（不含 max）；
	// 即新切片从老切片的 low 下标元素开始，len = high - low, cap = max - low；
	// high 和 max 一旦超出在老切片中越界，就会发生 runtime err，slice out of range。
	printLenCap("slice2: ", slice2) // len=5 cap=6 [1 2 3 4 5]

	// append 操作
	slice2 = append(slice2, slice1...) // 将 slice1 的所有元素插入到 slice2 中
	printLenCap("slice2: ", slice2)    // len=10 cap=12 [1 2 3 4 5 1 2 3 4 5]， cap 扩大两倍

	// delete 操作 (每次删除时间复杂度为O(n))
	// 1
	// s[:2] 区间截取，左闭右开
	slice2 = append(slice2[:2], slice2[3:]...) // 删除 slice2 的第三个元素
	// 删除元素不会缩小 cap，只会缩小 len
	printLenCap("slice2: ", slice2) // len=9 cap=12 [1 2 4 5 1 2 3 4 5]

	// 2
	slice2 = slice2[:2+copy(slice2[2:], slice2[3:])] // 删除 slice2 的第三个元素
	printLenCap("slice2: ", slice2)                  // len=8 cap=12 [1 2 5 1 2 3 4 5]

	// delete 操作（GC）
	slice3 := deleteEleGC(slice2, 2) // 删除 slice2 的第三个元素
	printLenCap("slice3: ", slice3)  // len=7 cap=8 [1 2 1 2 3 4 5]
	printLenCap("slice2: ", slice2)  // len=8 cap=12 [1 2 5 1 2 3 4 5]

	// insert 操作
	slice2 = append(slice2[:2], append([]int{100, 200}, slice2[2:]...)...)
	printLenCap("slice2: ", slice2) // len=10 cap=12 [1 2 100 200 5 1 2 3 4 5]

	// filter 操作
	slice3 = filter(slice2, func(v int) bool {
		return v > 2
	})
	printLenCap("slice3: ", slice3) // len=6 cap=8 [100 200 5 3 4 5]

	// push 操作
	slice2 = append(slice2, 1000)
	printLenCap("slice2: ", slice2) // len=11 cap=12 [1 2 100 200 5 1 2 3 4 5 1000]

	// pop 操作
	poppedEle, slice2 := slice2[len(slice2)-1], slice2[:len(slice2)-1]
	fmt.Println("poppedEle: ", poppedEle) // 1000
	printLenCap("slice2: ", slice2)       // len=10 cap=12 [1 2 100 200 5 1 2 3 4 5]

	// shift 操作
	shiftedEle, slice2 := slice2[0], slice2[1:]
	fmt.Println("shiftedEle: ", shiftedEle) // 1
	printLenCap("slice2: ", slice2)         // len=9 cap=12 [2 100 200 5 1 2 3 4 5]

}

func deleteEleGC(sli []int, idx int) []int {
	sliCopied := append([]int(nil), sli...)
	if idx < len(sli)-1 {
		copy(sliCopied[idx:], sliCopied[idx+1:])
	}
	sliCopied[len(sliCopied)-1] = 0 // nil or the zero value for the slice's element type
	sliCopied = sliCopied[:len(sliCopied)-1]
	return sliCopied
}

func filter[T int | string | rune | float32 | float64](s []T, f func(T) bool) []T {
	var ret []T
	for _, v := range s {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
