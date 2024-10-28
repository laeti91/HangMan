package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// creates a LetterIndices structure used to have all the letters and their position in the word when the user enters the correct letters into the terminal
type LetterIndices struct {
	Letter  string
	Indices []int
}

func LinesInTxtDoc(name string) []string {
	/*This function takes the name of the txt file we want to open and iterates each line of that file adding them into the allLines array which will then
	be returned.*/
	var allLines []string

	f, err := os.Open(name)
	scanner := bufio.NewScanner(f) // it creates a scanner to read the file and the split function as default is the ScanLines (line by line)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		allLines = append(allLines, scanner.Text()) //it appends each line in allLines to return all the lines in the file as an array
	}
	return allLines
}

func printWordGuessStatus(word string, wordFoundLetters map[rune]bool, lett string) {
	/*This function prints the status of the word to find with the letters already found by the user and a _ for the ones not yet found. The letter given
	by the user will also be green if it is in the word.*/
	var Green = "\033[32m" // creates a variable with the green color
	var Reset = "\033[0m"  // used to set the color back to the default
	wordPrinted := ""
	//this loop looks if each letter in the word is in the structure wordFoundLetters to then add it to wordPrinted, if not it will add a _.
	for _, characters := range word {
		if wordFoundLetters[characters] {
			if string(characters) == lett {
				wordPrinted += Green + string(characters) + Reset //the green color will be added only to the letter given by the user if it's in the word
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
	/*This function puts half of the letters of the word in tab, it also takes into account double letters*/
	n := len(word)/2 - 1
	var tab []LetterIndices
	totalIndices := 0

	for totalIndices <= n { //the loop continues until the number of letters reached half of the word
		letterInd := rand.Intn(len(word)) //takes a random index of the word to then get a random letter
		letter := string(word[letterInd])
		found := false
		for i := range tab { //iterates tab to know if the random letter is already in it
			if tab[i].Letter == letter {
				found = true
				break
			}
		}
		if !found { //if the letter is not already in tab, that letter will be integrated to tab with it's indexes as a value
			var indices []int
			for i, char := range word { //cheks for doubles and appends there indexes in indices
				if string(char) == letter {
					indices = append(indices, i)
				}
			}
			tab = append(tab, LetterIndices{Letter: letter, Indices: indices})
			totalIndices += len(indices) //it increments totalIndices by the number of times that the random letter is present in the word
		}
	}
	return tab
}

func foundAllLetters(word string, wordFoundLetters map[rune]bool) bool {
	/*This function checks if all the letters in the word have been found by the user. It returns a boolean saying true if the word is found
	and false otherwise.*/
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
	allWordsFile := LinesInTxtDoc("words2.txt")        //makes an array with all the words in words2.txt
	word := allWordsFile[rand.Intn(len(allWordsFile))] //takes a random word of the array wich will be the mistery word

	wordFoundLetters := make(map[rune]bool)
	for _, li := range nUniqueRandomLetters(word) { //iterates each letter of the type []LetterIndices returned in nUniqueRandomLetters(word)
		for _, i := range li.Indices { //iterates the idices of thoses letters
			wordFoundLetters[rune(word[i])] = true //puts true to each letter showing at the begining
		}
	}
	//calls the printWordGuessStatus function to show half of the letters in the terminal and _ for the unknown letters
	printWordGuessStatus(word, wordFoundLetters, "")

	scanner := bufio.NewScanner(os.Stdin)
	attempts := 10
	for attempts > 0 { //creates a loop to continue asking the user until the number of attemtps left reached 0
		fmt.Print("Enter a letter: ")
		scanner.Scan()
		letter := scanner.Text() //takes the letter entered by the user
		fmt.Println()
		if len(letter) > 1 {
			fmt.Println(Red + "Please enter only one letter." + Reset + "\n")
			continue
		} else if len(letter) == 0 {
			fmt.Println(Red + "Please enter a letter." + Reset + "\n")
			continue
		}
		letterGiven := rune(letter[0])
		if wordFoundLetters[letterGiven] { //if the word is already in wordFoundLetters it means that the letter has already been tried
			fmt.Println(Red + "You already tried that letter" + Reset + "\n")
			continue
		}
		wordFoundLetters[letterGiven] = true
		if strings.ContainsRune(word, letterGiven) { //if the letter is in the word to find a good message is printed
			fmt.Println(Green+"Correct answer, ", letter, "is present in the word"+Reset)
		} else { //otherwise the number of attempts decreases and the hangman is showed according to the number of attempts made
			attempts--
			ensembleLigneHangman := LinesInTxtDoc("hangman.txt")
			//the loop shows the lines in hangman.txt from the number of attemps multiplied by 8 to that number plus 8 because each hangman take 8 lines
			for i := (10 - attempts - 1) * 8; i < (10-attempts-1)*8+8; i++ {
				fmt.Println(ensembleLigneHangman[i])
			}
			if attempts > 0 {
				fmt.Println(Red+"Wrong answer, you still have", attempts, "attempts to discover the word"+Reset)
			}
		}
		printWordGuessStatus(word, wordFoundLetters, letter)
		if foundAllLetters(word, wordFoundLetters) { //if all the foundAllLetters function returns true it means that the word has been found
			fmt.Println(Green+"Congratulations, you found the word:", word+Reset)
			break
		}
		if attempts == 0 { //if the number of attempts reached 0 the word is showed with a message
			fmt.Println(Red+"Your number of attempts reached 0. The word was:", word+Reset)
		}
	}
}
