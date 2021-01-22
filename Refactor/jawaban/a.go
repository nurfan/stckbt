package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Nurfan Ramadhandi Backend Engineer (Golang)"
	res := findFirstStringInBracket(str)

	fmt.Println(res)
}

func findFirstStringInBracket(str string) string {

	indexFirstBracketFound := strings.Index(str, "(")

	if indexFirstBracketFound >= 0 {
		runes := []rune(str)
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
		indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")

		if indexClosingBracketFound >= 0 {
			runes := []rune(wordsAfterFirstBracket)
			return string(runes[1:indexClosingBracketFound])
		}
	}

	return ""
}
