package main

import "fmt"

/***
下面是一个制作自增整数生成器的例子，直到主线向信道索要数据，我们才添加数据到信道
 */
//xrange用来生成自增的整数
func xrange() chan int {
	var ch chan int = make(chan int)
	go func() { //开出一个goroutine
		for i := 0; ; i++ {
			ch <- i; //直到信道索要数据，才把i添加到信道
		}

	}()
	return ch
}

func main() {
	generator := xrange()
	for i := 0; i < 1000; i++ { // 我们生成1000个自增的整数！
		fmt.Println(<-generator)
	}
}
