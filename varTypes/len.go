package main
import "fmt"

func main () {
	myName := "bhaskar"
	lastName := "settipalli"
	newLastName := "settipalli"
	fmt.Println ("Currently the length of my name is ", len(myName) + len (lastName))
	fmt.Println ("The new length of my name will be ", len(myName) + len(newLastName))
	fmt.Println ("Hooray! Nothing changed.")
}
