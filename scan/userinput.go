package main

import "fmt"

func main() {

	var age int
	fmt.Println("what is your age")
	fmt.Scan(&age)

	fmt.Println("after five years your age is", age+5)
}