package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}

	// for i := 0; i < 10; {
	// 	fmt.Println("masuk")

	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Perulangan ke", i)
	// 	i++
	// }
}
