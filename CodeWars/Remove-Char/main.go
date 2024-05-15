package main

import (
	"fmt"
)

func main() {
	fmt.Println(RemoveChar("Merhaba"))
}

func RemoveChar(word string) string {
	result := ""
	for i := 1; i < len(word)-1; i++ {
		result += string(word[i])
	}
	return result
}

/*
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
*/

/*
func RemoveChar(word string) string {
	result := make([]byte, len(word))

	for i := 1; i < len(word)-1; i++ {
		result = append(result, word[i])
	}
	return string(result)
}
*/
