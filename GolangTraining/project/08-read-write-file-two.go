package main
/**
 * read-write-files.go
 *
 * The io/ioutil examples previously worked on the whole file at once
 * The majority of the time, that is all you need. However, if for some
 * reason you want to read or write line-by-line
 *
 */

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main(){

	filePath:="./extras/rabbits.txt"
	// read the file by using buffer io
	// The Scan() method scans token by token, default token is ScanLines
	// ScanWords and other methods available, See http://golang.org/pkg/bufio/

	f,_:=os.Open(filePath)
	// defer is a nifty bit of magic which automatically runs
	// the command before exiting in this case close file
	// good practice to defer right after opening
	/**
	执行完了关闭该文件
	 */
	defer f.Close()

	scanner:=bufio.NewScanner(f)

	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Print(line)
	}

	//check if any errors occured
	if err:=scanner.Err();err!=nil{
		log.Fatal(err)
	}


	f,err:=os.Create("output.txt")

	if err!=nil{
		log.Fatal("error create file",err)
	}

	defer f.Close()
	for _,str:=range[]string{"sdf","sdf"}{
		bytes,err:=f.WriteString(str)
		if err!=nil{
			log.Fatal("error writing string",err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}


}