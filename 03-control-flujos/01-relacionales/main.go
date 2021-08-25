package main

import "fmt"

func main() {
	a := 4
	b := 5

	var r bool

	//Igualdad
	r = a == b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)

	//Distito
	r = a != b
	fmt.Printf("%d es distinto que %d? %t \n", a, b, r)

	//Mayor que
	r = a > b
	fmt.Printf("%d es mayor que %d? %t \n", a, b, r)

	//Menor que
	r = a < b
	fmt.Printf("%d es menor que %d? %t \n", a, b, r)

	//Mayor o igual que
	r = a >= b
	fmt.Printf("%d es mayor o igual que %d? %t \n", a, b, r)

	//Menor o igual que
	r = a <= b
	fmt.Printf("%d es menor o igual que %d? %t \n", a, b, r)
}
