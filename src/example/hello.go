package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	var x int = 5
	var z int = 14
	fmt.Println(x+z)
	k := 18
	fmt.Println(k)

	if x>12 {
		fmt.Println("Big number")
	}

	var a [5]int
	a[4] = 6
	fmt.Println(a)

	b := [4]int{4,9,8,3}

	fmt.Println(b)

}