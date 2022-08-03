package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	Person
	gun_license bool
}

func main() {
	fmt.Println("We are learning about embedded struct")
	fmt.Println("----------------")

	p1 := secretAgent{
		Person: Person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		gun_license: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.gun_license)
	fmt.Println(p1.Person)
	fmt.Println(p1.Person.first)
	fmt.Println("--------------------")

	p2 := Person{
		first: "Jeffy",
		last:  "Mahin",
		age:   27,
	}
	fmt.Println(p2)
	fmt.Println("--------------")

	p3 := secretAgent{
		Person:      p2,
		gun_license: false,
	}

	fmt.Println(p3)
	fmt.Println(p3.first)
	fmt.Println(p3.last)
	fmt.Println(p3.gun_license)
	fmt.Println(p3.Person)
	fmt.Println(p3.Person.first)
}
