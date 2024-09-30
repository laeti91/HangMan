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
	word := WordToFind() //gives a random word in words2.txt
	fmt.Println(word)
	letterIndices := nUniqueRandomLetters(word) //creates a dictionnary where the keys are the the random letters in word and the values are the position of each letter in word
	fmt.Println(letterIndices)
	printWordWithHiddenLetters(word, letterIndices) //uses the dictionnary to show the word with only the letters in the dictionnary and _ to hide the others
}

func WordToFind() string {
	f, err := os.Open("words2.txt") //opens the file words2.txt to be able to read it
	scanner := bufio.NewScanner(f)  //scan each line of the file
	nbrMots := 0

	for scanner.Scan() { //the loop will count the numbers of lines in the file
		nbrMots++
	}

	if err != nil { //this loop checks if there are any problems scanning the file
		log.Fatal(err)
	}

	randomNumber := rand.Intn(nbrMots) //choosing a random line in the file

	return Scan(randomNumber) //calls the func Scan() to return the word in that line
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
		if nbrMots2 == nbr { //it will scan the file words2.txt line by line until the word is in the nbr line
			mot = scanner.Text()
		}
	}
	return mot //it will return the word
}

type LetterIndices struct { //creates a dictionnary used in nUniqueRandomLetters()
	Letter  string
	Indices []int
}

func nUniqueRandomLetters(word string) []LetterIndices {
	rand.Seed(time.Now().UnixNano()) //ensure different random numbers are generated each time the function is called
	n := len(word)/2 - 1             //n is the half of the word
	var tab []LetterIndices          //creates an empty slice of LetterIndices structure named tab
	totalIndices := 0

	for totalIndices < n { //the loop continues until the total number of indices collected is bigger or equal to n
		letterInd := rand.Intn(len(word)) //choosing a random index in the word
		letter := string(word[letterInd]) //takes the letter in index letterInd
		found := false
		for i := range tab { //if the letter is already in tab then  it appends the new index to the existing Indices slice
			if tab[i].Letter == letter {
				tab[i].Indices = append(tab[i].Indices, letterInd)
				found = true
				break
			}
		}
		if !found { //if the letter is not in tab then it creates a new indices, finds all indices of the letter in the word, and appends a new LetterIndices structure to the tab slice
			var indices []int
			for i, char := range word {
				if string(char) == letter {
					indices = append(indices, i)
				}
			}
			tab = append(tab, LetterIndices{Letter: letter, Indices: indices})
			totalIndices += len(indices) //updates the totalIndices counter with the number of indices added
		} else {
			totalIndices++ //if the letter was found the totalIndices increases by one
		}
	}
	return tab
}

func printWordWithHiddenLetters(word string, letterIndices []LetterIndices) {
	hiddenWord := strings.Repeat("_", len(word)) //creates a string with only _ the len() of word
	hiddenWordRunes := []rune(hiddenWord)        //convert that string in slices of runes

	for _, li := range letterIndices { //this loop iterates over each keys in letterIndices
		for _, index := range li.Indices { //this loop iterates over each value of the key li
			hiddenWordRunes[index] = rune(word[index]) //puts the letter in the right index of the word
		}
	}

	fmt.Println(string(hiddenWordRunes))
}
