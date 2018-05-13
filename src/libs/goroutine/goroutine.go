package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

/*
 goroutine 理解

 Go调度器内部有三个重要的结构：M P G

 M：(work thread)代表真正的内核OS线程，和POSIX里的thread差不多
 G：(goroutine)代表一个goroutine，用于被调度，它有自己的栈，指令，程序计数器和其他信息
 P：(processor)代表调度的上下文，它使go代码在一个线程上执行

 P的数量可以通过GOMAXPROCS来设置，它其实代表了真正的并发度，即有多少个goroutine可以同时运行，
*/

func fn0() {
	runtime.GOMAXPROCS(1)

	var total = 3
	var wg = &sync.WaitGroup{}
	wg.Add(total)

	//for i := 0; i < total; i++ {
	//	go func(i int, wg *sync.WaitGroup) {
	//		defer func() {
	//			wg.Add(-1)
	//		}()
	//		fmt.Println(i)
	//	}(i, wg)
	//}

	go func(i int, wg *sync.WaitGroup) {
		defer wg.Add(-1)
		fmt.Println(i)
	}(0, wg)

	go func(i int, wg *sync.WaitGroup) {
		defer wg.Add(-1)
		fmt.Println(i)
	}(1, wg)

	go func(i int, wg *sync.WaitGroup) {
		defer wg.Add(-1)
		fmt.Println(i)
	}(2, wg)

	wg.Wait()

	fmt.Println("===========")

	wg.Add(2)
	var chn1 = make(chan int)
	var chn2 = make(chan int)

	go func(wg *sync.WaitGroup) {
		defer wg.Add(-1)
		for i := 1; i < 10; i += 2 {
			fmt.Println(<-chn2)
			chn1 <- i
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Add(-1)
		for i := 0; i < 9; i += 2 {
			chn2 <- i
			fmt.Println(<-chn1)
		}
	}(wg)

	wg.Wait()
}
