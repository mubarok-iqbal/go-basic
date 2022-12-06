package main

import "fmt"

func main() {
	var nilai32 int32 = 10000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32) // 10000
	fmt.Println(nilai64) // 10000
	fmt.Println(nilai8)  // 16 => karena maksimum di -128 s.d 127 , looping sampai berhenti di 10000

	var name = "Muhammad Iqbal Mubarok"
	var m = name[0]
	var mString string = string(m)

	fmt.Println(m)       // 77
	fmt.Println(mString) // M

}
