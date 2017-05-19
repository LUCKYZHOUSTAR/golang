package main
/**
 * functions.go
 * Examples on how to use functions in go
 */

import (
	"fmt"
	"math"
)
// Basic function accepting one parameter
// You must specify the type of variables expected
// The format is function_name(variable type)
func Echo(s string){
	fmt.Println(s)
}
// Function with single return value
// the type of the return value is specifed after function declaration
func say(s string) string{
	phrase:="hello"+s
	return phrase
}

// Function with multiple parameters and return values
func Divide(x,y float64)(float64,float64){
	q:=math.Trunc(x/y)
	r:=math.Mod(x,y)
	return q,r
}


// Function with multiple parameters and named return values
// If the types are the same you can specify the type once at the end
func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return
}


func main(){
	Echo("bonjour tout le monde")
	fmt.Println(say("hello"))

	q,r:=Divide(11,3)

	fmt.Printf("Quotient: %v, Remainder %v \n", q, r)

}

