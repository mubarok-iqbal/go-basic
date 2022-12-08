package main

import "fmt"

func main() {
	var a float64 = 10
	var b float64 = 20

	var result1 = a + b
	var result2 = 10 * 20

	fmt.Println(result1, result2)

	//augmented assignment
	a += 30
	b /= 40
	fmt.Println(a)
	fmt.Println(b)

	//unary operator
	var c = 5
	c++
	c--
	fmt.Println(c)
}
