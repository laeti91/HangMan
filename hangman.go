package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(MotATrouver())
}
func MotATrouver() string {
	file, err := os.ReadFile("word.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
