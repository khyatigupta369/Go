package main

import "fmt"

func main() {
	a := 5
	b := &a
	// & - address of

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// * - value at
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)
	fmt.Println(*&b)

}
