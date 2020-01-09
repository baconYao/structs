package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastNamse string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastNamse: "Anderson",
		contactInfo: contactInfo{
			email: "qwe@gamil.com",
			zipCode: 12345,
		},
	}

	alexPointer := &alex
	alexPointer.updateName("Yao")
	alex.print()


	peter := person{
		firstName: "Peter",
		lastNamse: "Chang",
		contactInfo: contactInfo{
			email: "peter@mail.com",
			zipCode: 9876,
		},
	}
	// 在這裡演示了不用像上面的alex一樣宣告成指向記憶體變數的參數，go也會自動幫你處理掉(只要你的receiver的type是指標)
	peter.updateName("Gary")
	peter.print()
}

// *person 代表的是這個receiver type要的是指標
func (pointerToPerson *person) updateName(newFirstName string) {
	// pointerToPerson 指向記憶體位址，(*pointerToPerson) 根據記憶體位址，取出值
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}