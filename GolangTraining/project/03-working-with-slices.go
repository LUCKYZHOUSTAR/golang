package main

import "fmt"

/**
 * working-with-slices.go
 *
 * Go has an array type, but most interactions are with slices which are built
 * off arrays. I don't worry too much about the name and just use them
 * The specification for slices are "[]T" where T is the type.
 * So "[]string" is a set of strings, "[]int" is a set of integers, and so on.
 */
func main(){
	var empty []int
	alphas:=[]string{"Abc","cdd","efd","ghs"}

	empty=append(empty,123)
	empty=append(empty,23)

	fmt.Println(empty)

	alphas=append(alphas,"sdff","asdf")
	fmt.Println("length",len(alphas))

	//retrieve a single element from slice
	fmt.Println(alphas[1])

	alphas2:=alphas[1:3]
	fmt.Println(alphas2)

	if eleExists("asdf",alphas){
		fmt.Print("exists")
	}

}


func eleExists(s string,a[]string)bool{
	for _,v:=range a{
		if v==s{
			return true
		}
	}
	return false
}