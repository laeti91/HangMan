package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {

	word := WordToFind()
	fmt.Println(word)
	wordrune := []rune(word)
	attempts := 10
	tab := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for attempts != 0 {
		fmt.Print("Enter a letter : ")
		scanner.Scan()
		letter := scanner.Text()

		if len(letter) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}

		array := []rune(letter)
		tab = tab[:0]

		for _, i := range wordrune {
			for j := 0; j < len(word); j++ {
				if i == array[0] {
					tab = append(tab, j)
					fmt.Println(tab)
					break
				}
			}
		}
		if len(tab) != 0 {
			fmt.Println("wright answer, ", letter, "is present", len(tab), "times in the word")
		} else if len(tab) == 0 {
			attempts--
			fmt.Println("wrong answer, you still have", attempts, "attempts to discover the word")
		}
	}
}

func WordToFind() string {
	f, err := os.Open("word.txt")
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
	f, err := os.Open("word.txt")
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
