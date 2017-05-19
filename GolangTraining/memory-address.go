package memoryaddress

import "fmt"

const metersToYards float64 = 1.09361

func main(){
	a:=43
	fmt.Println("a-",a)
	fmt.Println("a memory address",&a)
	scan()
}


func scan(){
	var meters float64
	fmt.Println("Enter meters swam:")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
