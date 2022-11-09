package main

import "fmt"

type nums []int

func main() {
	n := nums{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, num := range n {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}