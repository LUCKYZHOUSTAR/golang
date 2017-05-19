package main

import "fmt"

var ch5 chan int = make(chan int)

func foo2(id int) { //id: 这个routine的标号
	ch5 <- id
}

func main() {
	// 开启5个routine
	for i1 := 0; i1 < 5; i1++ {
		go foo2(i1)
	}

	// 取出信道中的数据--宏观上我们看到的即 无缓冲信道的数据是先到先出
	for i := 0; i < 5; i++ {
		fmt.Print(<- ch5)
	}
}
