package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type LetterIndices struct {
	Letter  string
	Indices []int
}

func readFile(name string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(name)
	scanner := bufio.NewScanner(f)

	if err != nil {
		log.Fatal(err)
	}
	return scanner, f
}

func WordToFind() string {
	var ensembleMots []string
	scanner, _ := readFile("word.txt")

	for scanner.Scan() {
		ensembleMots = append(ensembleMots, scanner.Text())
	}
	return ensembleMots[rand.Intn(len(ensembleMots))]
}

func getStatus(word string, wordFoundLetters map[rune]bool) {
	letterIndices := nUniqueRandomLetters(word)

	for _, li := range letterIndices {
		for _, i := range li.Indices {
			wordFoundLetters[rune(word[i])] = true
		}
	}
}

func printWordGuessStatus(word string, wordFoundLetters map[rune]bool) string {
	wordPrinted := ""
	for _, characters := range word {
		if wordFoundLetters[characters] {
			wordPrinted += string(characters)
		} else {
			wordPrinted += "_"
		}
	}
	return wordPrinted
}

func nUniqueRandomLetters(word string) []LetterIndices {
	n := len(word)/2 - 1
	var tab []LetterIndices
	totalIndices := 0

	for totalIndices < n {
		letterInd := rand.Intn(len(word))
		letter := string(word[letterInd])
		found := false
		for i := range tab {
			if tab[i].Letter == letter {
				tab[i].Indices = append(tab[i].Indices, letterInd)
				found = true
				break
			}
		}
		if !found {
			var indices []int
			for i, char := range word {
				if string(char) == letter {
					indices = append(indices, i)
				}
			}
			tab = append(tab, LetterIndices{Letter: letter, Indices: indices})
			totalIndices += len(indices)
		} else {
			totalIndices++
		}
	}
	return tab
}

func GetHangman(nbr int) {
	scanner, f := readFile("hangman.txt")

	defer f.Close()

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

func GetAsciiLett(nbr int) []string {

	scanner, f := readFile("standard.txt")

	defer f.Close()

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

func main() {

	fmt.Println("\nWelcome to the hangman game !")

	fmt.Println("\nGood luck you have 10 attemps !")

	word := WordToFind()

	wordFoundLetters := make(map[rune]bool)
	getStatus(word, wordFoundLetters)
	word1 := printWordGuessStatus(word, wordFoundLetters)
	PrintAsciiHugeLett(word1)

	scanner := bufio.NewScanner(os.Stdin)

	for attempts := 10; attempts > 0; {
		fmt.Print("\nEnter a letter : ")
		scanner.Scan()
		letter := scanner.Text()

		if len(letter) != 1 {
			fmt.Println("Please enter only one letter.\n")
			continue
		}

		letterGiven := rune(letter[0])

		if wordFoundLetters[letterGiven] {
			fmt.Println("You already tried that letter\n")
			continue
		}
		wordFoundLetters[letterGiven] = true

		if strings.ContainsRune(word, letterGiven) {
			fmt.Println("wright answer, ", letter, "is present in the word\n")
		} else {
			attempts--
			nbr := (10 - attempts - 1) * 8
			GetHangman(nbr)
			if attempts > 0 {
				fmt.Println("wrong answer, you still have", attempts, "attempts to discover the word\n")
			}
		}

		word2 := printWordGuessStatus(word, wordFoundLetters)
		PrintAsciiHugeLett(word2)

		foundAllLetters := true
		for _, characters := range word {
			if !wordFoundLetters[characters] {
				foundAllLetters = false
			}
		}

		if foundAllLetters {
			fmt.Println("Congratulation, you found the word :", word, "\n")
			break
		}

		if attempts == 0 {
			fmt.Println("Your number of attempts reached 0. The word was : ", word, "\n")
		}
	}
}
