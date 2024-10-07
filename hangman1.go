package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	word := WordToFind()
	fmt.Println(word)
	letterIndices := nUniqueRandomLetters(word)
	fmt.Println(letterIndices)
	printWordWithHiddenLetters(word, letterIndices)
	//essaye := 0
	//nbr := (10 - essaye - 1) * 8
	//GetHangman(nbr)
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
	rand.Seed(time.Now().UnixNano())
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

func printWordWithHiddenLetters(word string, letterIndices []LetterIndices) {
	hiddenWord := strings.Repeat("_", len(word))
	hiddenWordRunes := []rune(hiddenWord)

	for _, li := range letterIndices {
		for _, index := range li.Indices {
			hiddenWordRunes[index] = rune(word[index])
		}
	}

	fmt.Println(string(hiddenWordRunes))
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
