package main

import "fmt"

func main() {
	result := 2
	bir := 1
	iki := 2
	uc := 3
	fmt.Print("1 2 ")
	for uc <= 4000000 {
		bir += iki
		iki += uc
		uc += iki
		fmt.Print(bir, iki, " ")
		if bir%2 == 0 {
			result += bir
		}
		if iki%2 == 0 {
			result += iki
		}
	}
	fmt.Println(" ")
	fmt.Println(result)
}
