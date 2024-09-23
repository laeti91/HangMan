package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	fmt.Println(WordToFind())
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

