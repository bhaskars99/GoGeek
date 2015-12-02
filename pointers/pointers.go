package main

import "fmt"

func main() {
	var a int = 43
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(*b)
}