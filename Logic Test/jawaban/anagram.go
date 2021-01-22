package main

import (
	"fmt"
	"sort"
	"strings"
)

type sortRunes []rune

func main() {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	anagram := anagram(input)

	for _, words := range anagram {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}

}

func anagram(words []string) map[string][]string {
	list := make(map[string][]string)

	for _, word := range words {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	return list
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
