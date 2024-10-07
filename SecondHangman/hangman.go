package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	nbr := rune(')')
	GetAsciiLett(int(nbr))
}

func GetAsciiLett(nbr int) {
	f, err := os.Open("standard.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	ligne := -9

	for scanner.Scan() {
		ligne = ligne + 9
		if ligne == (nbr-32)*9 {
			for i := ligne; i < ligne+9; i++ {
				print(i)
				fmt.Println(scanner.Text())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
