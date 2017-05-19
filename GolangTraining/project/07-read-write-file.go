package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	filename:="./project/rabbits.txt"


	// read in file, one command loads file into content variable
	// if an error occurs returns it as the second return value
	// if no error, err will be nil
	constant,err:=ioutil.ReadFile(filename)
	if err!=nil{
		// log is a nifty little utility which can also output
		// in this case, a fatal log will halt the program
		log.Fatal("error reading file",filename,nil)
	}


	// content returned as []byte not string
	// so must cast to string first and then can display
	fmt.Println(string(constant))

	// write back to new file
	// see documentation for which methods take what type
	outfile:="./project/output.txt"
	err=ioutil.WriteFile(outfile,constant,0644)
	if err!=nil{
		log.Fatal("Error writing file",err)
	}


}
