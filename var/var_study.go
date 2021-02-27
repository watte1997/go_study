package main

import "fmt"

const (
	BEIJING=iota*100
	SHANGHAI
	SHENZHEN
)

func main() {
	fmt.Println("hello world")

	var h int = 1
	fmt.Println(h)

	var c=100
	fmt.Println(c)
	fmt.Printf("type of c is %T\n",c)

	e:="100"
	fmt.Printf("type 0f e is %T\n",e)

	fmt.Println(SHENZHEN)
}
