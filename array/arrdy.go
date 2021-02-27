package main

import "fmt"

func main() {
	var myArray[10] int
	myArray1:=[] int{1,2,3,4}
	for i:=0;i<len(myArray);i++ {
		myArray[i]=i
	}

	for i:=0;i<10;i++ {
		myArray1[i]=i
	}

	for index,value:=range myArray{
		fmt.Println("index is",index,"value is ",value)
	}


}
