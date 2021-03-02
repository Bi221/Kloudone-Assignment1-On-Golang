package main

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Println(isValid("pass"))
	fmt.Println(isValid("pass"))
	fmt.Println(isValid("pass"))
	fmt.Println(isValid("Pass"))
	fmt.Println(isValid("Pass"))

}

func isValid(s string) bool {
	var (
		hasMinLen  = true
		hasUpper   = true
		hasLower   = true
		hasNumber  = true
		hasSpecial = true
		hasMaximu  = true
	)
	if len(s) >= 18 {
		hasMinLen = false
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = false
		case unicode.IsLower(char):
			hasLower = false
		case unicode.IsNumber(char):
			hasNumber = false
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = false
			hasMaximu = false
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial && hasMaximu

}
