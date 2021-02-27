package main

import "fmt"

func foo1(a string,b int)(c ,d int){
	fmt.Println("a is ",a)
	fmt.Println("b is",b)

	c=b+100
	d=b+11
	return
}

func main() {
	var a,b="hello",100
	c,d:=foo1(a,b)
	fmt.Println("c,d is",c,d)
}
