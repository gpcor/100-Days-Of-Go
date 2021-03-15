package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toNato("if you can read this, you can read nato"))
}

func toNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
	newString := strings.Split(words, "")
	// var betterString []string

	for i, w := range newString {
		for _, n := range nato {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(w)) {
				newString[i] = n
			}
		}
	}

	// fmt.Println(betterString)

	return strings.ReplaceAll(strings.Join(newString, " "), "   ", " ")
}
