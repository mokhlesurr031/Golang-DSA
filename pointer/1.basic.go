package main

import "fmt"

func main() {
	a := 42
	fmt.Println("--------------First-------------")
	fmt.Println(a)
	fmt.Println(&a) //address of var a or pointer to a

	fmt.Println("--------------Second--------------")
	fmt.Printf("%T\n", a)  //type of var a
	fmt.Printf("%T\n", &a) //type of pointer of a

	fmt.Println("--------------Third--------------")
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) //dereference to the address

	fmt.Printf("------------Fourth-----------------\n")
	c := &a
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(*&a)
}
