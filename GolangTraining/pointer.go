package main

import "fmt"

func main(){
	a:=43
	fmt.Println(a)
	fmt.Println(&a)
	var b=&a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	point2()


	x:=5
	zero(x)
	fmt.Println(x)
}

func point2(){
	a:=43
	fmt.Println(a)
	fmt.Println(&a)

	var b=&a
	fmt.Println(b)
	fmt.Println(*b)

	*b=56
	fmt.Println(b)

	fmt.Println(a)
}


func zero(z int){
	z=0
}


func zeropoint(z *int){
	*z=0
}




