package main

import "fmt"

func main() {
	fmt.Println(enBuyukAsal(16))
}

func enBuyukAsal(x int) int {
	var result int
	for i := x; i > 1; i-- {
		if x%i == 0 && asalMi(i) {
			result = i
			break
		}
	}
	return result
}

func asalMi(sayi int) bool {
	if sayi < 2 {
		return false
	}
	if sayi == 2 {
		return true
	}
	if sayi%2 == 0 {
		return false
	}
	for i := 3; i < sayi; i += 2 {
		if sayi%i == 0 {
			return false
		}
	}
	return true
}

/*
func asalMi(sayi int) bool {
	if sayi < 2 {
		return false
	}

	if sayi <= 3 {
		return true
	}

	if sayi%2 == 0 || sayi%3 == 0 {
		return false
	}

	for i := 5; i*i <= sayi; i += 6 {
		if sayi%i == 0 || sayi%(i+2) == 0 {
			return false
		}
	}
	return true
}
*/
/* func abcd(metin string) string {
	result := ""
	for _, char := range metin {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32)
		} else {
			result += string(char)
		}
	}
	return result
}

func ciftmi(sayi int) bool {
	if sayi%2 == 0 {
		return true
	} else {
		return false
	}
}
*/
