package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	word := WordToFind()
	fmt.Println(word)
	fmt.Println(nUniqueRandomLetters(word))
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

func nUniqueRandomLetters(word string) string {
	n := len(word)/2 - 1

	rand.Seed(time.Now().UnixNano())
	letters := []rune(word)
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})

	uniqueLetters := make(map[rune]bool)
	result := make([]rune, 0, n)

	for _, letter := range letters {
		if !uniqueLetters[letter] {
			uniqueLetters[letter] = true
			result = append(result, letter)
			if len(result) == n {
				break
			}
		}
	}

	return string(result)
}
