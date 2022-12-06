package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Iqbal Mubarok"
	fmt.Println(name)

	var age int8 = 12
	fmt.Println(age)

	var city = "Bandung"
	fmt.Println(city)

	country := "Indonesia" // var country string = "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Muhammad Iqbal"
		lastName  = "Mubarok"
	)

	fmt.Println(firstName, lastName)
}
