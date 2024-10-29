package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

/* LetterIndices holds a letter and its indices in the word. */
type LetterIndices struct {
	Letter  string
	Indices []int
}

/* LinesInTxtDoc reads a file put in parameters and returns its lines as an array of string. */
func LinesInTxtDoc(name string) []string {

	var allLines []string

	f, err := os.Open(name)
	scanner := bufio.NewScanner(f) // bufio.NewScanner is used to create a scanner to read the file

	if err != nil {
		log.Fatal(err) // it handles file opening error
	}

	for scanner.Scan() { // it scans through the file
		allLines = append(allLines, scanner.Text()) // it adds each line to the allLines array
	}
	return allLines
}

/* printWordGuessStatus prints the current state of the word with found letters and underscores for missing ones. */
func printWordGuessStatus(word string, wordFoundLetters map[rune]bool) string {
	wordPrinted := ""
	for _, characters := range word {
		if wordFoundLetters[characters] { // it verifies if the letter is present in the word
			wordPrinted += string(characters) // if it is, it's add to the string
		} else {
			wordPrinted += "_" // Otherwise, it adds an underscore
		}
	}
	return wordPrinted
}

/* nUniqueRandomLetters returns an array of unique letters from the word along with their indices. */
func nUniqueRandomLetters(word string) []LetterIndices {
	n := len(word)/2 - 1 // it takes the half of the length of the word
	var tab []LetterIndices
	totalIndices := 0

	for totalIndices <= n { //it continues until half the number of letters is reached
		letterInd := rand.Intn(len(word)) // it takes a random index to get a random letter
		letter := string(word[letterInd])
		found := false
		for i := range tab { // it checks if the letter is already in tab
			if tab[i].Letter == letter {
				found = true
				break
			}
		}
		if !found { // if the letter is not in tab, it adds it with it's index
			var indices []int
			for i, char := range word { // it finds all indices of this letter in the word
				if string(char) == letter {
					indices = append(indices, i)
				}
			}
			tab = append(tab, LetterIndices{Letter: letter, Indices: indices})
			totalIndices += len(indices) // it increments totalIndices by the number of occurrences
		}
	}
	return tab
}

/* GetAsciiLett retrieves the ASCII art representation of a letter based on its ASCII code. */
func GetAsciiLett(nbr int, newLett bool) []string {

	var Reset = "\033[0m"
	var Green = "\033[32m"
	list := LinesInTxtDoc("standard.txt")
	/* it calculates the position of the caracter in the ASCII file. Subtract 32 because the first character
	in the file is at position 32 and multiply by 9 because each characters takes 9 lines in the file. */
	actualLine := (nbr - 32) * 9
	var asciiLetter []string

	for i := actualLine; i < actualLine+9; i++ {
		if newLett {
			asciiLetter = append(asciiLetter, Green+list[i]+Reset) // it colors the newly found letters in green
		} else {
			asciiLetter = append(asciiLetter, list[i]) // Otherwise, it adds the line
		}
	}
	return asciiLetter
}

/* PrintAsciiHugeLett prints the ASCII art of the letters in the input string. */
func PrintAsciiHugeLett(input, lett string) {
	theLetter := false
	var asciiHugeLett [][]string

	for _, char := range input {
		if string(char) == lett {
			theLetter = true // it marks the letter as found
		} else {
			theLetter = false
		}
		asciiHugeLett = append(asciiHugeLett, GetAsciiLett(int(char), theLetter)) // it gets the ASCII art for the letter
	}
	// it prints the ASCII art line by line
	for i := 0; i < 9; i++ {
		for _, letter := range asciiHugeLett {
			fmt.Print(letter[i])
		}
		fmt.Println()
	}
}

func main() {

	var Reset = "\033[0m"  // it puts back the default color
	var Red = "\033[31m"   // it sets the text color to red
	var Green = "\033[32m" // it sets the text color to green

	emojiSadFaces := [3]string{"\U0001F622", "\U0001F61E", "\U0001F62D"} // to print sad emojis when the guesses are wrong

	/* smiling face emojis is printed att the beginning of the welcome sentence */
	fmt.Println("\n\U0001F60A " + "Welcome to the hangman game !")

	fmt.Println("\nGood luck you have 10 attemps !")

	allWordsFile := LinesInTxtDoc("word.txt")          // it makes an array with all the words from the file
	word := allWordsFile[rand.Intn(len(allWordsFile))] // it randomly select a word

	wordFoundLetters := make(map[rune]bool)
	for _, li := range nUniqueRandomLetters(word) { //// it gets unique letters and their indices
		for _, i := range li.Indices {
			wordFoundLetters[rune(word[i])] = true // it marks the found letters
		}
	}

	// it displays the current status of the word
	PrintAsciiHugeLett(printWordGuessStatus(word, wordFoundLetters), "")

	scanner := bufio.NewScanner(os.Stdin)

	// Main game loop
	for attempts := 10; attempts > 0; {
		fmt.Print("\nEnter a letter : ")
		scanner.Scan()
		letter := scanner.Text()

		if len(letter) > 1 {
			fmt.Println(Red + "Please enter only one letter." + Reset) // it prints the text in red
			continue
		} else if len(letter) == 0 {
			fmt.Println(Red + "Please enter a letter." + Reset)
			continue
		}

		letterGiven := rune(letter[0])

		if wordFoundLetters[letterGiven] { // it checks if the letter has already been tried
			fmt.Println(Red + "You already tried that letter" + Reset)
			continue
		}
		wordFoundLetters[letterGiven] = true

		if strings.ContainsRune(word, letterGiven) { //it checks if the letter is in the word
			fmt.Println("\n\U0001F600 "+Green+"Wright answer, ", letter, "is present in the word"+Reset)
		} else { // otherwise, if the letter is not in the word
			attempts--
			ensembleLigneHangman := LinesInTxtDoc("hangman.txt")
			/* it displays the hangman based on remaining attempts. One hangman position takes 8 lines in the file*/
			for i := (10 - attempts - 1) * 8; i < (10-attempts-1)*8+8; i++ {
				fmt.Println(ensembleLigneHangman[i])
			}
			if attempts > 0 {
				fmt.Println(emojiSadFaces[rand.Intn(3)]+Red+" Wrong answer, you still have", attempts, "attempts to discover the word"+Reset)
			}
		}

		PrintAsciiHugeLett(printWordGuessStatus(word, wordFoundLetters), letter)

		// it checks if all letters have been found
		foundAllLetters := true
		for _, characters := range word {
			if !wordFoundLetters[characters] {
				foundAllLetters = false
			}
		}

		if foundAllLetters {
			fmt.Println("\n\U0001F601 "+Green+"Congratulation, you found the word :", word+Reset)
			break
		}

		if attempts == 0 {
			fmt.Println("\n\U0001F62D "+Red+"Your number of attempts reached 0. The word was :", word+Reset)
		}
	}
}
