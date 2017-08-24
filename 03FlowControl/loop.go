package main

import "fmt"

func main() {
	//common for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	printLine()

	//for like while loop
	b := 20
	a := 1
	for a < b {
		fmt.Println(a)
		a *= 2
	}
	printLine()

	//for loop like c (;;) sort of infinite if no break
	c := 5
	for {
		if c < 0 {
			break
		}
		c--
		fmt.Println(c)
	}
}

func printLine() {
	fmt.Println("==========")
}
