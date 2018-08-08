package main

import "fmt"

func main() {
	ints := [11]int{}
	for i := range ints {
		ints[i] = i
		fmt.Print(i)
		if i%2 == 0 {
			fmt.Println(" is even")
		} else {
			fmt.Println(" is odd")
		}
	}
}
