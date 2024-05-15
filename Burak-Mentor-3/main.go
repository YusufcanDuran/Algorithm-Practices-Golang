package main

import "fmt"

func main() {
	for i := 1; i <= 250; i++ {
		if i%15 == 0 && i%2 != 0 && i != 135 {
			fmt.Print(i, ", ")
		}
	}
	fmt.Println("")
}
