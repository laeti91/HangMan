package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func readFile(name string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(name)
	scanner := bufio.NewScanner(f)

	if err != nil {
		log.Fatal(err)
	}
	return scanner, f
}

func wordToFind() string {
	nbrWord := 0
	scanner, _ := readFile("words2.txt")

	for scanner.Scan() {
		nbrWord++
	}

	randomNumber := rand.Intn(nbrWord)

	return scanWord(randomNumber)
}

func scanWord(nbr int) string {
	word := ""
	nbrWord2 := 0

	scanner, _ := readFile("words2.txt")

	for scanner.Scan() {
		nbrWord2++
		if nbrWord2 == nbr {
			word = scanner.Text()
		}
	}
	return word
}

func getStatus(word string, wordFoundLetters map[rune]bool) {
	letterIndices := nUniqueRandomLetters(word)

	for _, li := range letterIndices {
		for _, i := range li.Indices {
			wordFoundLetters[rune(word[i])] = true
		}
	}
	printWordGuessStatus(word, wordFoundLetters)
}

func printWordGuessStatus(word string, wordFoundLetters map[rune]bool) {
	wordPrinted := ""
	for _, characters := range word {
		if wordFoundLetters[characters] {
			wordPrinted += string(characters)
		} else {
			wordPrinted += "_"
		}
	}
	fmt.Println(wordPrinted)
}

func allLettersFound(word string, wordFoundLetters map[rune]bool) bool {
	for _, characters := range word {
		if !wordFoundLetters[characters] {
			return false
		}
	}
	return true
}

type LetterIndices struct {
	Letter  string
	Indices []int
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

func getHangman(nbr int) {
	scanner, f := readFile("hangman.txt")

	defer f.Close()

	line := 0

	for scanner.Scan() {
		line++
		if line >= nbr && line <= nbr+8 {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Println("\nWelcome to the hangman game !")
	fmt.Println("You have 10 attemps, good luck !\n")

	word := wordToFind()

	wordFoundLetters := make(map[rune]bool)
	getStatus(word, wordFoundLetters)

	attempts := 10

	scanner := bufio.NewScanner(os.Stdin)

	for attempts != 0 {
		fmt.Print("Enter a letter : ")
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
			getHangman(nbr)
			if attempts > 0 {
				fmt.Println("wrong answer, you still have", attempts, "attempts to discover the word\n")
			}
		}

		printWordGuessStatus(word, wordFoundLetters)

		if allLettersFound(word, wordFoundLetters) {
			fmt.Println("Congratulation, you found the word :", word, "\n")
			break
		}

		if attempts == 0 {
			fmt.Println("Your number of attempts reached 0. The word was : ", word, "\n")
		}
	}
}
