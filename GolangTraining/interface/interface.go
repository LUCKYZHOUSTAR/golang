package main

import "fmt"



/***
只要struct里面有interfac里面的方法，就算实现了该接口
 */
type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}


type shape interface {
	area() float64
}

func info(z shape){
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	fmt.Println("Area: ", s.area())

	s1 := square{10}
	fmt.Printf("%T\n",s1)
	info(s1)
}