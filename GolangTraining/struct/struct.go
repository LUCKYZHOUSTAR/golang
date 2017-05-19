package main

import (
	"fmt"
	"encoding/json"
	"os"
)


/**
可以自定义类型
 */
type foo int
func main(){
	var myAge foo
	myAge=44
	fmt.Println(myAge,myAge)

	var yourAge int
	yourAge=29

	fmt.Println(yourAge)

	p1:=person{"jamie","bound",20}
	fmt.Println(p1.first)

	fmt.Println(p1.fullName())

	bs,_:=json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))

	p2 := person2{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p2)
	//p1 := doubleZero{
	//	person: person{
	//		First: "James",
	//		Last:  "Bond",
	//		Age:   20,
	//	},
	//	LicenseToKill: true,
	//}
	//
	//p2 := doubleZero{
	//	person: person{
	//		First: "Miss",
	//		Last:  "MoneyPenny",
	//		Age:   19,
	//	},
	//	LicenseToKill: false,
	//}
	//
	//fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	//fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}


func test1(){
	var myAge foo
	myAge=44
	fmt.Print(myAge)


	var yourAge int
	yourAge=29
	fmt.Print(yourAge)


}


type person struct{
	first string
	last string
	age int
}



/**
定义接口的方法
 */
func (p person) fullName()string{
	return p.first+p.last
}


type doubleZero struct {
	person
	LicenseTokill bool
}


type person2 struct {
	First       string
	Last        string
	Age         int
	notExported int
}