package main
import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop2() {
	for i := 0; i < 100; i++ { //为了观察，跑多些
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

/***
从输出结果中，可以看出，共同一个线程，抢夺式的请求
默认地， Go所有的goroutines只能在一个线程里跑 。

也就是说， 以上两个代码都不是并行的，但是都是是并发的。

如果当前goroutine不发生阻塞，它是不会让出CPU给其他goroutine的, 所以例子一中的输出会是一个一个goroutine进行的，而sleep函数则阻塞掉了 当前goroutine, 当前goroutine主动让其他goroutine执行, 所以形成了逻辑上的并行, 也就是并发。
0 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 1 17 18 19 2 3 4 5 6 7 8 9 20 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 21 22 23 24 25 26 27 28 29 30 31 56 57 58 59 60 61 62 63 64 32 33 34 35 36 37 38 39 40 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 41 42 43 44 45 81 82 83 84 85 86 87 88 46 47 48 49 50 51 52 53 54 55 56 89 90 91 92 93 94 95 96 97 98 99 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99
 */
func main() {
	runtime.GOMAXPROCS(2) // 最多使用2个核

	go loop2()
	go loop2()

	for i := 0; i < 2; i++ {
		<- quit
	}
}