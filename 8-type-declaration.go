package main

import "fmt"

func main() {
	type NoKTP string
	type Marriage bool

	var noKtpIqbal NoKTP = "12350236202829"
	var marriageStatus Marriage = true

	fmt.Println(noKtpIqbal)
	fmt.Println(marriageStatus)
}
