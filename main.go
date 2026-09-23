// package main

// import "fmt"

// func main() {

// 	fmt.Println("Hello World")
// 	fmt.Println("Привет от GO")

// }
// package main //начанию делать по второму видосу

// import "fmt"

//	func main() {
//		var name string
//		var age int8
//		fmt.Println("What is your name?")
//		fmt.Scan(&name)
//		fmt.Println("Hello " + name + "!")
//		fmt.Println("How old are you?")
//		fmt.Scan(&age)
//		fmt.Println("You are " + fmt.Sprint(age) + "years!")
//	}
package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Println("РЕШИ КВАДРАТНОЕ УРОВНЕНИЕ")

	fmt.Println("Введи а:")
	fmt.Scan(&a)

	fmt.Println("Введи b:")
	fmt.Scan(&b)

	fmt.Println("Введи c:")
	fmt.Scan(&c)
	D := (b * b) - 4*(a*c)
	if D > 0 {
		var x1 float64
		var x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уровнение имеет 2 корня\nD=" + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println("Ваше уровнение имеет 1 корень \nD = 0")
		fmt.Println("X: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Ваше уровнение не имеет корней \nD < 0 \nD = " + fmt.Sprint(D))
	}

}
