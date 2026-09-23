// package main

// import "fmt"

// func main() {

// 	fmt.Println("Hello World")
// 	fmt.Println("Привет от GO")

// }
package main //начанию делать по второму видосу

import "fmt"

func main() {
	var name string
	var age int8
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age) + "years!")
}
