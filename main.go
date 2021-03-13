package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://corsini.tech")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%T", resp.Body)

}
