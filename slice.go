package main

import "fmt"

func main() {

	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst2 := []string{"potato", "potato", "potato"}
	src2 := []string{"watermelon", "pinnaple"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2) // watermelon pinnaple potato
	fmt.Println(src2) // watermelon pinnaple
	fmt.Println(n2)   // 2

}
