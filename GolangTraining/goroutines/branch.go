package main

import (
	"time"
	"math/rand"
	"fmt"
)

/***
每一个chan存数据的时候，都应该是goroutine来执行的
 */
// 一个比较耗时的事情，比如计算
func do_stuff(x int) int {
	time.Sleep(time.Duration(rand.Intn(10)))
	return 100 - x
}

// 每个分支开出一个goroutine做计算并把计算结果流入各自信道
func branch(x int) chan int {
	ch := make(chan int)

	go func() {
		ch <- do_stuff(x)
	}()

	return ch
}

//流出符号:<-
//流入符号ch<-值
//ch<-<-c
func fanIn(chs ...chan int) chan int {
	ch := make(chan int)
	for _, c := range chs {
		// 注意此处明确传值
		go func(c chan int) { ch <- <-c }(c) // 复合
	}

	return ch
}
func main() {
	//先进先出
	result := fanIn(branch(1), branch(2), branch(3))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
