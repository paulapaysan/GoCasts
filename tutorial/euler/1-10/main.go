package main

import (
	"fmt"
)

func main() {

}

func prob1() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("The sum of all the multiples of 3 or 5 below 1000 is", sum)
}

func prob2() {

}
