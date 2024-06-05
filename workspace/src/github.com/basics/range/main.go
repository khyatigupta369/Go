package main

import "fmt"

func main() {

	emails := map[string]string{"Khyati": "Khyati@gmail.com", "Jeff": "Jeff@ama.in"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
