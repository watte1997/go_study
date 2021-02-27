package main

import "fmt"

func main() {
	var mymap map[string]string
	if mymap==nil{
		fmt.Println("mymap is empty")
	}

	mymap=make(map[string]string,3)
	mymap["one"]="JAVA"
	mymap["two"]="GO"
	mymap["three"]="cpp"

	fmt.Println(mymap)
}
