package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string = []string{"kita", "atik", "tika", "aku", "kua"}

// TestAnagram anagram true case
func TestAnagram(t *testing.T) {

	expected := make(map[string][]string)
	expected["aku"] = []string{"aku", "kua"}
	expected["aikt"] = []string{"kita", "atik", "tika"}

	actual := anagram(input)

	assert.Equal(t, actual, expected, "TestAnagram : Anagram Error")
}

func TestSortStr(t *testing.T) {
	expected := "aabc"
	word := "baca"

	actual := sortStr(word)

	if expected != actual {
		t.Error("TestSortStr: Sort string error")
	}
}

func TestSortStrError(t *testing.T) {
	expected := "aambc"
	word := "bacam"

	actual := sortStr(word)

	if expected == actual {
		t.Error("TestSortStrError: Sort string error")
	}
}
