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

func printWordGuessStatus(word string, wordFoundLetters map[rune]bool, lett string) {
	var Green = "\033[32m"
	var Reset = "\033[0m"
	wordPrinted := ""
	for _, characters := range word {
		if wordFoundLetters[characters] {
			if string(characters) == lett {
				wordPrinted += Green + string(characters) + Reset
			} else {
				wordPrinted += string(characters)
			}

		} else {
			wordPrinted += "_"
		}
	}

	fmt.Println("\n" + wordPrinted)
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

func foundAllLetters(word string, wordFoundLetters map[rune]bool) bool {
	for _, char := range word {
		if !wordFoundLetters[char] {
			return false
		}
	}
	return true
}

func main() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"

	fmt.Println("\nWelcome to the hangman game!")
	fmt.Println("You have 10 attempts, good luck!\n")
	allWordsFile := LinesInTxtDoc("words2.txt")
	word := allWordsFile[rand.Intn(len(allWordsFile))]

	wordFoundLetters := make(map[rune]bool)
	for _, li := range nUniqueRandomLetters(word) {
		for _, i := range li.Indices {
			wordFoundLetters[rune(word[i])] = true
		}
	}
	printWordGuessStatus(word, wordFoundLetters, "")

	scanner := bufio.NewScanner(os.Stdin)
	attempts := 10
	for attempts > 0 {
		fmt.Print("Enter a letter: ")
		scanner.Scan()
		letter := scanner.Text()
		fmt.Println()
		if len(letter) > 1 {
			fmt.Println(Red + "Please enter only one letter." + Reset + "\n")
			continue
		} else if len(letter) == 0 {
			fmt.Println(Red + "Please enter a letter." + Reset + "\n")
			continue
		}
		letterGiven := rune(letter[0])
		if wordFoundLetters[letterGiven] {
			fmt.Println(Red + "You already tried that letter" + Reset + "\n")
			continue
		}
		wordFoundLetters[letterGiven] = true
		if strings.ContainsRune(word, letterGiven) {
			fmt.Println(Green+"Correct answer, ", letter, "is present in the word"+Reset)
		} else {
			attempts--
			ensembleLigneHangman := LinesInTxtDoc("hangman.txt")
			for i := (10 - attempts - 1) * 8; i < (10-attempts-1)*8+8; i++ {
				fmt.Println(ensembleLigneHangman[i])
			}
			if attempts > 0 {
				fmt.Println(Red+"Wrong answer, you still have", attempts, "attempts to discover the word"+Reset)
			}
		}
		printWordGuessStatus(word, wordFoundLetters, letter)
		if foundAllLetters(word, wordFoundLetters) {
			fmt.Println(Green+"Congratulations, you found the word:", word+Reset)
			break
		}
		if attempts == 0 {
			fmt.Println(Red+"Your number of attempts reached 0. The word was:", word+Reset)
		}
	}
}
