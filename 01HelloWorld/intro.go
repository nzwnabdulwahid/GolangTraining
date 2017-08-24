package main


import "fmt"

const (

	_ = iota //ignoring 0 with blank identifier
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (1 * 10)
	GB = 1 << (iota * 10) // 1 << (1 * 10)
	TB = 1 << (iota * 10) // 1 << (1 * 10)

)

func Explain_IOTA(){
	fmt.Printf("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%v\n", KB)
}