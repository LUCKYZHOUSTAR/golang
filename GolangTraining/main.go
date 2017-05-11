package main

import (
	"fmt"
)



/*
go语言的数据类型操作
 */
func dataType() {
	var b bool = true;
	fmt.Println(b)

	//数字类型：int和float
	/**
		uint8
无符号 8 位整型 (0 到 255)
2	uint16
无符号 16 位整型 (0 到 65535)
3	uint32
无符号 32 位整型 (0 到 4294967295)
4	uint64
无符号 64 位整型 (0 到 18446744073709551615)
5	int8
有符号 8 位整型 (-128 到 127)
6	int16
有符号 16 位整型 (-32768 到 32767)
7	int32
有符号 32 位整型 (-2147483648 到 2147483647)
8	int64
有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	*/

	/**
		float32
IEEE-754 32位浮点型数
2	float64
IEEE-754 64位浮点型数
3	complex64
32 位实数和虚数
4	complex128
64 位实数和虚数*/
	/*
		byte
类似 uint8
2	rune
类似 int32
3	uint
32 或 64 位
4	int
与 uint 一样大小
5	uintptr
无符号整型，用于存放一个指针*/
	var tempCount uint8;
	tempCount = 0;
	fmt.Println(tempCount)
	fmt.Println(tempCount)
	//字符串类型
	//派生类型
	/*
	(a) 指针类型（Pointer）
	(b) 数组类型
	(c) 结构化类型(struct)
	(d) Channel 类型
	(e) 函数类型
	(f) 切片类型
	(g) 接口类型（interface）
	(h) Map 类型
	 */


}

var x, y int
var (
	// 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
//变量定义信息
func varaible() {
	var v_name string = "hello"
	fmt.Println(v_name)
	var name = "hello"//类型自动判断
	fmt.Println(name)

	//多变量声明信息
	var name1, name2, name3 string
	name1, name2, name3 = "hi", "hello", "haa"
	fmt.Println(name3, name2, name1)
	//g, h := 123, "hello"
	//println(x, y, a, b, c, d, e, f, g, h)
}

const c_name1, c_name2 = "hi", "hello"

func constgrammer() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	//const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	//const (
	//	a = "abc"
	//	b = len(a)
	//	c = unsafe.Sizeof(a)
	//)

	/***
	*iota，特殊常量，可以认为是一个可以被编译器修改的常量。
在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
	 */
	const (
		a = iota
		b = iota
		c = iota
	)

}

func operator() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - c 的值为 %d\n", a)
	a--
	fmt.Printf("第七行 - c 的值为 %d\n", a)

	//var a3 int = 21
	//var b3 int = 10

	if ( a == b ) {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于  b\n")
	}
	if ( b >= a ) {
		fmt.Printf("第五行 - b 大于等于 b\n")
	}

	//var a bool = true
	//var b bool = false
	//if ( a && b ) {
	//	fmt.Printf("第一行 - 条件为 true\n")
	//}
	//if ( a || b ) {
	//	fmt.Printf("第二行 - 条件为 true\n")
	//}
	/* 修改 a 和 b 的值 */
	//a = false
	//b = true
	//if ( a && b ) {
	//	fmt.Printf("第三行 - 条件为 true\n")
	//} else {
	//	fmt.Printf("第三行 - 条件为 false\n")
	//}
	//if ( !(a && b) ) {
	//	fmt.Printf("第四行 - 条件为 true\n")
	//}


	//位运算符
	/***
	*p	q	p & q	p | q	p ^ q
	0	0	0	0	0
	0	1	0	1	1
	1	1	1	1	0
	1	0	0	1	1
	*
	*
	 */


	/*****
	*
	  var a int = 21
var c int

c =  a
fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

c +=  a
fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

c -=  a
fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

c *=  a
fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

c /=  a
fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

c  = 200;

c <<=  2
fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

c >>=  2
fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

c &=  2
fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

c ^=  2
fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

c |=  2
fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )

	 */


	/**
	其他运算符
	&:返回变量存储地址
	*：指针变量
	 */

}



//条件语句信息
func condition() {
	var a int = 10
	if a < 20 {
		fmt.Print("a 小于20\n")
	} else if (a > 20) {

	}

	var b int = 15
	numbers := [6]int{1, 2, 3, 5}
	for a := 0; a < 10; a++ {
		fmt.Print(a)
	}

	for a < b {
		a++
		fmt.Print(b)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

}



//函数定义
func function_name() {
	/**
	func function_name( [parameter list] ) [return_types] {
   函数体
}
	*/


}

func max(num1, num2 int) int {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}


//函数返回多个值信息
func swap(x, y string) (string, string) {
	return y, x
}


//值传递
func swap2(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

func swap3(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
func main() {
	fmt.Println("hello world!");
	var b int = 100
	var a int = 50

	fmt.Printf("交换前，a 的值 : %d\n", a)
	fmt.Printf("交换前，b 的值 : %d\n", b)

	/* 调用 swap() 函数
	* &a 指向 a 指针，a 变量的地址
	* &b 指向 b 指针，b 变量的地址
	*/
	swap3(&a, &b)

	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)
}

//变量作用域范围信息
func varableBoundary() {
	/***
	函数内定义的变量称为局部变量
	函数外定义的变量称为全局变量
	函数定义中的变量称为形式参数
	 */
}


func array(){
	var arr [12] int
	var balance=[12]int{1,2,3,4}
	balance[4]=0
	fmt.Println(arr)
}


func point(){
	var a int=10
	fmt.Printf("变量的地址: %x\n", &a  )
	/**一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。*/
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	//空指针信息
	var ptr *int
	fmt.Println("空指针信息为",ptr)
	if(ptr != nil) {

	}    /* ptr 不是空指针 */
}



//定义结构体信息
type Book struct {
	title string
	author string
	subject string
	bookId int
}

func calcuteBook(){
	var book1 Book
	book1.author="张是哪"
	fmt.Println(book1.author)
}

/**
相当于动态的数组
 */
func slice(){

}