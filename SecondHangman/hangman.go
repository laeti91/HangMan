package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

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

func allLettersFounds(word string, wordFoundLetters map[rune]bool) bool {
	for _, characters := range word {
		if !wordFoundLetters[characters] {
			return false
		}
	}
	return true
}

func WordToFind() string {
	f, err := os.Open("words2.txt")
	scanner := bufio.NewScanner(f)
	nbrMots := 0

	for scanner.Scan() {
		nbrMots++
	}

	if err != nil {
		log.Fatal(err)
	}

	randomNumber := rand.Intn(nbrMots)

	return Scan(randomNumber)
}

func Scan(nbr int) string {
	mot := ""
	nbrMots2 := 0
	f, err := os.Open("words2.txt")
	scanner := bufio.NewScanner(f)

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		nbrMots2++
		if nbrMots2 == nbr {
			mot = scanner.Text()
		}
	}
	return mot
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

func main() {

	fmt.Println("\nWelcome to the hangman game !")
	fmt.Print("Are you ready to guess the word and have fun. (Y/N) : ")

	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	cont := scanner1.Text()

	if cont == "Y" {

		fmt.Println("\nGood luck you have 10 attemps !")

		word := WordToFind()

		letterIndices := nUniqueRandomLetters(word)
		wordFoundLetters := make(map[rune]bool)

		for _, li := range letterIndices {
			for _, i := range li.Indices {
				wordFoundLetters[rune(word[i])] = true
			}
		}
		word1 := printWordGuessStatus(word, wordFoundLetters)
		PrintAsciiHugeLett(word1)

		attempts := 10

		scanner := bufio.NewScanner(os.Stdin)

		for attempts != 0 {
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

			if allLettersFounds(word, wordFoundLetters) {
				fmt.Println("\nCongratulation, you found the word :", word, "\n")
				break
			}

			if attempts == 0 {
				fmt.Println("Your number of attempts reached 0. The word was : ", word, "\n")
			}
		}
	} else {
		fmt.Println("See you next time !")
	}
}
