package main
import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func main() {
	wg1.Add(2)
	go foo1()
	go bar1()
	wg1.Wait()
}

func foo1() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg1.Done()
}

func bar1() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg1.Done()
}
