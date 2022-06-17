package channel_practice

import (
	"fmt"
	"math/big"
	"time"

	// "math/rand"
	"crypto/rand"
	"sync"
)

type Sender = chan<- int
type Receiver = <-chan int

func ChanPractice() {
	fmt.Println("Chan Practice")

	noCachedPipeline()

	cachedPipeline()

	iterateChan()

	deadlock()

	RWMutexPractice()
}

func noCachedPipeline() {

	var startWg sync.WaitGroup

	startWg.Add(1)

	fmt.Println("\n\n无缓冲信道练习：")

	startWg.Done()

	startWg.Wait()

	var wg sync.WaitGroup

	wg.Add(2)

	// 无缓冲信道
	noCachedPipeline := make(chan int) // 即 make(chan int, 0)

	// 无缓冲信道，发送端和接收端需要同步运行

	go func() {
		defer wg.Done()
		// 单向发信道
		var sender Sender = noCachedPipeline
		big, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			fmt.Println(err)
		}
		data := big.Int64()
		sender <- int(data)
		fmt.Println("发送的数据是： ", data)
	}()

	go func() {
		defer wg.Done()
		// 单向收信道
		var receiver Receiver = noCachedPipeline
		num := <-receiver
		// 只接收但不获取数据可以這樣寫： <-receiver
		fmt.Println("接收到的数据是： ", num)
	}()

	wg.Wait()
}

func cachedPipeline() {

	var startWg sync.WaitGroup

	startWg.Add(1)

	fmt.Println("\n\n缓冲信道练习：")

	startWg.Done()

	startWg.Wait()

	var wg sync.WaitGroup

	// 有缓冲信道
	// 设置缓冲后，发送端和接收端可以异步
	cachedPipeline := make(chan int, 10)

	wg.Add(2)

	go func() {
		defer wg.Done()
		var sender Sender = cachedPipeline
		for i := 0; i < 10; i++ {
			big, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				fmt.Println(err)
			}
			data := big.Int64()
			sender <- int(data)
			fmt.Printf("发送第%d个数据是： %d\n", i, int(data))
		}
	}()

	go func() {
		defer wg.Done()
		var receiver Receiver = cachedPipeline
		count := 0

		for {
			data := <-receiver
			count++
			fmt.Printf("接收到的第%d个数据是：%d\n", count, data)
			// 直到缓冲区没有数据，才会结束
			if len(receiver) == 0 {
				break
			}
		}
	}()

	wg.Wait()
}

func iterateChan() {
	var startWg sync.WaitGroup

	startWg.Add(1)

	fmt.Println("\n遍历信道练习：")

	startWg.Done()

	startWg.Wait()

	pipeline := make(chan int, 10)

	fibonacci := func(myChan chan int) {
		n := cap(myChan)
		x, y := 1, 1
		for i := 0; i < n; i++ {
			myChan <- x
			x, y = y, x+y
		}
		// 记得 close 信道
		// 不然主函数中遍历完并不会结束，而是会阻塞
		close(myChan)
	}

	go fibonacci(pipeline)

	for i := range pipeline {
		fmt.Println(i)
	}
}

func deadlock() {
	var startWg sync.WaitGroup

	startWg.Add(1)

	fmt.Println("\n死锁练习：")

	startWg.Done()

	startWg.Wait()

	// 第一种情况：接收者在发送者发送数据时没有准备好，造成循环等待
	// pipeline := make(chan string)
	// pipeline <- "hello"
	// fmt.Println(<- pipeline)
	// 修复方法：改为有缓冲信道： pipeline := make(chan string, 1)

	// 第二种情况：信道容量满后，不能往信道发送数据，若消费者速度跟不上生产者速度，就会造成阻塞
	// ch1 := make(chan string, 1)
	// ch1 <- "hello"
	// ch1 <- "world"
	// fmt.Println(<- ch1)

	// 第三种情况：程序循环等待接受数据
	// pipeline := make(chan string)
	// go func() {
	// 		pipeline <- "hello world"
	// 		pipeline <- "hello China"
	// }()
	// go func()  {
	// 	for {
	// 		fmt.Println(<- pipeline)
	// 	}
	// }()
	// for data := range pipeline {
	// 	fmt.Println(data)
	// }
	// 修复方法：生产者在发送完数据后手动关闭信道: close(pipeline)
	pipeline := make(chan string)
	go func() {
		pipeline <- "hello world"
		pipeline <- "hello China"
		close(pipeline)
	}()
	for data := range pipeline {
		fmt.Println(data)
	}

}

func RWMutexPractice() {
	// 读写锁：
	// 读锁占用时，写锁阻塞
	// 多个协程拥有读锁，互相之间不阻塞
	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始... \n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)

	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}
