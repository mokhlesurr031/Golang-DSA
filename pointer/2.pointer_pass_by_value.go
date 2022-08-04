package main

import "fmt"

func foo(x int) {
	fmt.Println(x)
	x = 43
	fmt.Println(x)

}

func bar(y *int) {
	fmt.Println("y address before", y)
	fmt.Println("y value", *y)
	*y = 43
	fmt.Println("y address after", y)
	fmt.Println("y value", *y)
}

func main() {
	fmt.Println("Learning Pointers\n")

	n := 0
	foo(n)
	fmt.Println(n)

	fmt.Println("--------------------")

	m := 0
	fmt.Println("m address before:", &m)
	bar(&m)
	fmt.Println("m value:", m)
	fmt.Println("m address after:", &m)
}
