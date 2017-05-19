package main

import "fmt"

/**chan一定要配对*/

func add(temp int) chan int {
	c := make(chan int)

	go func() {
		for{
			i:=1
			fmt.Println("开始放入",i)
			c <- temp
			fmt.Println("执行逻辑")
			i++
		}


	}()

	return c
}

func main() {
	c := add(3)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
