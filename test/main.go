package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	attempts := 10

	letter := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for attempts != 0 {
		fmt.Print("Enter a letter: ")

		scanner.Scan()

		text := scanner.Text()

		fmt.Println(text)
		letter = append(letter, text)
		attempts--
	}

	fmt.Println(letter)
}
