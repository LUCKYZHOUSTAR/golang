package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)


func main(){
	res,err:=http.Get("http://www.baidu.com")
	if err!=nil{
		log.Fatal(err)
	}

	page,err:=ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
	http2()
}


func http2(){
	res,_:=http.Get("http://www.baidu.com")
	page,_:=ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("%s",page)
}