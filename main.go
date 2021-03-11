package main

import (
	"fmt"
	"strings"
)

func main() {
	High("a dsblje edqu ysorpqal trlal nnmq qcfdzgwhng pzpfxc gjbrmcd arad boyxe nudzgakz igevikyn quhadm wgotw oarhsyw ixebfjzd dtmvsuieha offm")
}

func High(s string) {
	values := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}

	stringSlice := strings.Split(s, " ")
	stringValue := make(map[string]int)

	for s := range stringSlice {
		sumValue := 0
		for key, value := range values {
			letterCount := strings.Count(stringSlice[s], key)
			if letterCount > 1 {
				sumValue += value * letterCount
			}
			if letterCount == 1 {
				sumValue += value
			}
		}
		stringValue[stringSlice[s]] = sumValue
	}

	highestValue := 0
	highestKey := ""

	for key, value := range stringValue {
		if value > highestValue {
			highestValue = value
			highestKey = key
		}
	}
	fmt.Println(highestKey)
}
