package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastNamse string
}

func main() {
	alex := person{firstName: "Alex", lastNamse: "Anderson"}
	fmt.Println(alex)

	var yao person
	yao.firstName = "yao"
	yao.lastNamse = "Chang"
	fmt.Println(yao)
	fmt.Printf("%+v", yao)
}