package main

import "fmt"

func FindMultiples(integer, limit int) []int {
	result := []int{}
	for i := integer; i <= limit; i += integer {
		result = append(result, i)
	}
	return result
}

func main() {
	fmt.Println(FindMultiples(2, 8))
}
