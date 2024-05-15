package main

import "fmt"

func main() {
	result := []int{}
	for i := 1; i <= 250; i++ {
		if i%3 == 0 && i%5 == 0 && i%2 != 0 {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
