package main

import "fmt"

func main() {

	jack := get_notification("jack")
	joe := get_notification("joe")

	//获取消息
	fmt.Println(<-jack)
	fmt.Println(<-joe)
}

func get_notification(user string) chan string {

	//查询用户的消息提醒
	notification := make(chan string)
	go func() { /*悬挂一个信道出去*/
		notification <- fmt.Sprint("hi", user)
	}()

	return notification
}
