package main

import (
	"fmt"
	"os"
	"flag"
)


var str string
var num int

var help bool
func main(){

	// command-line args are stored as slice in os.Args
	// first argument in list is program itself
	num_args := len(os.Args)
	// check if received any command line arguments
	if num_args < 2 {
		fmt.Println(">> No args passed in")
	}


	// flag package provides parsing of command-line parameters
	// this example we create global variables and then pass them in
	// as pointers which BoolVar, StringVar and IntVar set as values
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "help", false, "Display Help")
	flag.Parse()
	// check if help was called explicitly
	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	fmt.Println(">> String:", str)
	fmt.Println(">> Number:", num)

	// last command-line argument
	fmt.Println(">> Last Item:", os.Args[num_args-1])

	// the os.Args will include flags for example
	// go run command-line-args.go --str=Foo filename
	// os.Args[1] = "--str=Foo"

	// If you have flags and want just the arguments
	// then after calling flag.Parse() you can call
	// flag.Args which store only the args
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(">> Flag Arg:", args[0])
	}
}
