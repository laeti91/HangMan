package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	nbr := rand.Intn(chooseWordRandom())
	fmt.Println(scanFile(nbr))
}

func chooseWordRandom() int {
	f, err := os.Open("words2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	nbrMots := 0

	for scanner.Scan() {
		nbrMots++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nbrMots
}

func scanFile(nbr int) string {
	f, err := os.Open("words2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	nbrMots := 0
	mot := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nbrMots++
		if nbrMots == nbr {
			mot = scanner.Text()
		}
	}
	return mot
}
