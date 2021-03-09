package main

import "fmt"

func main() {
	p0 := 1000
	percent := 0.02
	aug := 500
	p := 1500

	newPop := p0 + int(float64(p0)*percent) + aug

	for newPop > p {
		newPop = newPop + (p0 + int(float64(p0)*percent) + aug)
	}

	fmt.Println(newPop)
}
