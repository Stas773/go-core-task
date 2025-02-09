package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var (
	numDecimal, numOctal, numHexadecimal int
	pi                                   float64
	name                                 string
	isActive                             bool
	complexNum                           complex64
)

func main() {
	fmt.Println(printType(numDecimal))
	fmt.Println(printType(numOctal))
	fmt.Println(printType(numHexadecimal))
	fmt.Println(printType(pi))
	fmt.Println(printType(name))
	fmt.Println(printType(isActive))
	fmt.Println(printType(complexNum))

	str := toString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	strRune := stringToRune(str)
	salt := "go-2024"
	fmt.Println(hashRunes(strRune, salt))
}

func printType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func toString(v ...interface{}) string {
	result := ""
	for _, v := range v {
		result += fmt.Sprint(v)
	}
	return result
}

func stringToRune(str string) []rune {
	return []rune(str)
}

func hashRunes(sliceRune []rune, salt string) string {
	saltRune := []rune(salt)
	res := append(sliceRune[:len(sliceRune)/2], append(saltRune, sliceRune[len(sliceRune)/2:]...)...)

	hash := sha256.Sum256([]byte(string(res)))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}
