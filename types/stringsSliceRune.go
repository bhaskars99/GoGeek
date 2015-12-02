package main

import (
	"fmt"
	"reflect"
)

func main(){
	lname :="settipalli"
	fname := "bhaskar"
	hello :="Hello"
	fmt.Println("I'm",fname)
	rune := fname[5]
	fmt.Println("rune is ", rune)
	fmt.Println("Sliced string is ",lname[5:10])

	fmt.Printf("%T\n",lname)
	fmt.Println(reflect.TypeOf(lname))
	fmt.Println(string("a"))
	fmt.Println([]byte{'h', 'e', 'l', 'l', 'o'})
	fmt.Println([]byte(hello))

	var l int = 12
	var m float64 = float64(l)
	fmt.Println(m)

	var x float64 = 12.1230123
	var y int = int(x)

	fmt.Println(y)
}

