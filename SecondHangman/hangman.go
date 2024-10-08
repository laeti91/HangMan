package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := "H"
	input1 := "ello"
	input2 := input + input1
	PrintAsciiHugeLett(input2)

}

func GetAsciiLett(nbr int) []string {

	f, err := os.Open("standard.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 0
	actualLine := (nbr - 32) * 9
	var asciiLetter []string

	for scanner.Scan() {
		if line >= actualLine && line < actualLine+9 {
			asciiLetter = append(asciiLetter, scanner.Text())
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return asciiLetter
}

func PrintAsciiHugeLett(input string) {
	var asciiHugeLett [][]string

	for _, char := range input {
		asciiHugeLett = append(asciiHugeLett, GetAsciiLett(int(char)))
	}

	for i := 0; i < 9; i++ {
		for _, letter := range asciiHugeLett {
			fmt.Print(letter[i])
		}
		fmt.Println()
	}
}

func GetHangman(nbr int) {
	f, err := os.Open("hangman.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	ligne := 0

	for scanner.Scan() {
		ligne++
		if ligne >= nbr && ligne <= nbr+8 {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
