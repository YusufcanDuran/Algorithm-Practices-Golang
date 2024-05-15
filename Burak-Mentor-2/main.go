package main

import "fmt"

func main() {
	result := make([]int, 8)
	x := 1
	for i := 0; i < 8; i++ {
		result[i] = 15 * x
		x += 2
	}
	fmt.Println(result)
}
