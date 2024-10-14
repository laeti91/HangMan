package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	nbr := rune('n')
	GetAsciiLett(int(nbr))
}

func GetAsciiLett(nbr int) {
	f, err := os.Open("standard.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 0
	actualLine := (nbr - 32) * 9

	for scanner.Scan() {
		if line >= actualLine && line < actualLine+9 {
			fmt.Println(scanner.Text())
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
