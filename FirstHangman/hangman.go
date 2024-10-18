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

func LinesInTxtDoc(name string) []string {
	var allLines []string

	f, err := os.Open(name)
	scanner := bufio.NewScanner(f)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}
	return allLines
}

func printWordGuessStatus(word string, wordFoundLetters map[rune]bool) {
	wordPrinted := ""
	fmt.Println()
	for _, characters := range word {
		if wordFoundLetters[characters] {
			wordPrinted += string(characters)
		} else {
			wordPrinted += "_"
		}
	}
	fmt.Println(wordPrinted)
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

func main() {

	fmt.Println("\nWelcome to the hangman game !")
	fmt.Println("You have 10 attemps, good luck !")
	fmt.Println()

	allWordsFile := LinesInTxtDoc("words2.txt")
	word := allWordsFile[rand.Intn(len(allWordsFile))]

	wordFoundLetters := make(map[rune]bool)
	letterIndices := nUniqueRandomLetters(word)

	for _, li := range letterIndices {
		for _, i := range li.Indices {
			wordFoundLetters[rune(word[i])] = true
		}
	}
	printWordGuessStatus(word, wordFoundLetters)

	letterScanner := bufio.NewScanner(os.Stdin)

	for attempts := 10; attempts > 0; {
		fmt.Print("Enter a letter : ")
		letterScanner.Scan()
		letter := letterScanner.Text()
		fmt.Println()

		if len(letter) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}

		letterGiven := rune(letter[0])

		if wordFoundLetters[letterGiven] {
			fmt.Println("You already tried that letter")
			continue
		}
		wordFoundLetters[letterGiven] = true

		if strings.ContainsRune(word, letterGiven) {
			fmt.Println("wright answer, ", letter, "is present in the word")
		} else {
			attempts--
			nbr := (10 - attempts - 1) * 8

			ensembleLigneHangman := LinesInTxtDoc("hangman.txt")
			for i := nbr; i < nbr+8; i++ {
				fmt.Println(ensembleLigneHangman[i])
			}

			if attempts > 0 {
				fmt.Println("wrong answer, you still have", attempts, "attempts to discover the word")
			}
		}

		printWordGuessStatus(word, wordFoundLetters)

		foundAllLetters := true
		for _, characters := range word {
			if !wordFoundLetters[characters] {
				foundAllLetters = false
			}
		}

		if foundAllLetters {
			fmt.Println("Congratulation, you found the word :", word, "")
			break
		}

		if attempts == 0 {
			fmt.Println("Your number of attempts reached 0. The word was : ", word)
		}
	}
}
