package main

import "fmt"

//var ch chan int = make(chan int)
//
//func foo() {
//	ch<-0//// 向ch中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo, 直到main函数把0这个数据拿走
//}
//func main() {
//	go foo()
//	<- ch // 从ch取数据，如果ch中还没放数据，那就挂起main线，直到foo函数中放数据为止
//}

var complete chan int = make(chan int)

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	complete <- 0 // 执行完毕了，发个消息
}


func main() {
	go loop1()
	<- complete // 直到线程跑完, 取到消息. main在此阻塞住
}