package coroutine

import (
	"fmt"
	"sync"
)

func worker(workerId int, wg *sync.WaitGroup) {
	defer wg.Done() // 当协程执行完毕后，调用 Done() 方法，通知调度器，协程已经结束, 计数器减1
	for i := 0; i < 10; i++ {
		fmt.Printf("worker %d: %d\n", workerId, i)
	}
}

func Coroutine() {
	var startWg sync.WaitGroup

	// 控制「协程练习」在 worker 之前输出
	startWg.Add(1)
	print("\n\n协程练习：\n")
	startWg.Done()
	startWg.Wait()

	// 协程是一种轻量级的线程，它可以被调度器调度，并且可以被调度器挂起，
	// 并且可以被调度器恢复，并且可以被调度器终止。

	var wg sync.WaitGroup

	wg.Add(2) // 增加计数器的值，计数器的值为2

	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait() // 阻塞当前程序，等待计数器的值为0后，才会执行下面的语句

}
