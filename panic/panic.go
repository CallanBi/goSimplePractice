package panic_practice

import (
	"fmt"
)

type MyCustomError struct {
	code int
	err  error
}

func (err MyCustomError) Error() string {
	return fmt.Sprintln("code: ", err.code, "err: ", err.err)
}

func setData(idx int) {
	// 若 idx 大于10， 则可以故意制造一个 panic 异常
	var arr [10]int
	arr[idx] = 404
}

func throwPanic() {
	panic(MyCustomError{code: 5000001, err: fmt.Errorf("custom error")})
}

func returnErr() (ok bool, err error) {
	return false, MyCustomError{code: 5000001, err: fmt.Errorf("custom error")}
}

func PanicPractice() {
	fmt.Println("\n\npanic_practice 练习：")

	defer func() {
		// recover() 可以捕获到 panic 异常
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
			fmt.Println("recover from panic")
			// debug.PrintStack()
		}
	}()

	_, err := returnErr()

	if err != nil {
		fmt.Println("returned err: ", err)
		fmt.Println(err)
	}

	throwPanic()

	// 如果可以捕获到 panic 异常，则不会执行下面的语句
	// 当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪，最后程序终止。
	setData(20)

	fmt.Println("\n\npanic_practice 练习结束：")
}
