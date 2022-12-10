package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Muhammad", "Iqbal", "Mubarok"}
	printMessage("halo", names)
}

func printMessage(message string, slc []string) {
	var nameString = strings.Join(slc, " ")
	fmt.Println(message, nameString)

	total := sumAll(10, 90, 30, 50, 40)
	fmt.Println(total)

	slice := []int{10, 20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}

//variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
