package main

import (
	"fmt"
)
var messages chan string=make(chan string)

func main(){
	var messages chan string = make(chan string)
	go func(message string) {
		messages <- message // 存消息
	}("Ping!")

	fmt.Println(<-messages) // 取消息

}
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

/***
信道的概念：
信道是什么？简单说，是goroutine之间互相通讯的东西。类似我们Unix上的管道（可以在进程间传递消息）， 用来goroutine之间发消息和接收消息。其实，就是在做goroutine之间的内存共享。
var channel chan int = make(chan int)
// 或
channel := make(chan int)
 */

func chan1(message string){
	messages<-message//存消息

}