package main

import "fmt"

func main() {
	slice1:= make([]int,10,15)

	fmt.Printf("slice len is%d,slice is %v\n",len(slice1),slice1)

	slice1 = append(slice1, 2)
	fmt.Printf("slice len is%d,slice is %v\n",len(slice1),slice1)
}
