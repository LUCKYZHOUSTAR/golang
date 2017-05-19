package main

import "fmt"


func main(){

	var myGreeting map[string]string
	fmt.Print(myGreeting)
	fmt.Print(myGreeting==nil)
	test1()

	test6()

	test8()
}


func test1(){
	var myGreeting=make(map[string]string)

	myGreeting["Time"]="good moring"
	myGreeting["Jenny"]="BonJour"

	fmt.Println(myGreeting)
}


func test2(){
	myGreeting:=make(map[string]string)
	myGreeting["Time"]="good moring"
	myGreeting["Jenny"]="BonJour"

	fmt.Println(myGreeting)
}


func test3(){
	mygreeting:=map[string]string{
		"Tim":"goodmoring",
		"jenny":"bonjoun",
	}

	fmt.Println(mygreeting)
}

func test4(){
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
}


func test5(){
	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)

	delete(myGreeting,"three")
}


func test6(){
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)
}


func test7(){
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}


func test8(){

	letter:=rune("A"[0])
	fmt.Println(letter)
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}