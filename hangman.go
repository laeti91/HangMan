package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {

	fmt.Println("Welcome to the hangman game !")
	fmt.Print("Are you ready to guess the word and have fun. (Y/N) : ")

	scanner1 := bufio.NewScanner(os.Stdin) // Pour faire entrer une valeur (oui/non)
	scanner1.Scan()                        //pour afficher la valeur(oui/non)
	cont := scanner1.Text()                // affecter la valeur entrée dans la variable cont(continuer)

	if cont == "Y" || cont == "y" { // si cont = oui alors

		fmt.Println("\nGood luck you have 10 attempts!")

		word := WordToFind() // on affecte le mot choisit de façon random a la variable word pour pouvoir réutiliser le mot

		letterIndices := nUniqueRandomLetters(word) // on appelle la fonction pour afficher les lettres indices et on l'affecte a une variable pour la réutilisée
		wordFoundLetters := make(map[rune]bool)     // on déclare un tableau sous forme de carte qui permet de crée un tableau de rune avec le [rune] et on ajoute un booleen qui permet de verifier si la valeur a déjà était ajoutée dedans

		for _, li := range letterIndices { // on parcours les valeurs de letterIndices
			for _, i := range li.Indices { // on parcours les valeurs de la structure letterIndices (uniquement l'indice ici ( le mot avec ----- et les lettres indices au milieu))
				wordFoundLetters[rune(word[i])] = true // on dit que la lettre est dans le wordfoundletter pour les test plus tard
			}
		}
		printWordGuessStatus(word, wordFoundLetters) // on affiche le status du mot avec des ---- et les lettres indices en appelant la fonction printwordguessstatus

		attempts := 10 // on a dix chances pour trouver le mot

		scanner := bufio.NewScanner(os.Stdin) // on fait saisir une lettre

		for attempts != 0 { // tant que le compteur de chances n'est pas a 0 la boucle continue
			fmt.Print("Enter a letter : ")
			scanner.Scan()           // on affiche la lettre entrée
			letter := scanner.Text() // on affecte la lettre dans une variable

			if len(letter) != 1 { // si il y a plus d'une lettre donc si letter a plus de 1 caractère alors
				fmt.Println("Please enter only one letter.\n") // on demande d'entrer une seule lettre
				continue                                       // la boucle reprend à la saisie de la lettre
			}

			letterGiven := rune(letter[0]) // on affecte la première rune de la string letter (la seule lettre donc) à une variable pour garder uniquement cette rune ci

			if wordFoundLetters[letterGiven] { // on teste si la lettre qui a était donnée à déjà était donnée
				fmt.Println("You already tried that letter\n") // si elle est déjà présente on dit que la lettre a déjà était entrée
				continue                                       // la boucle reprend à la saisie de la lettre
			}
			wordFoundLetters[letterGiven] = true // on fait en sorte que la lettre donnée ne puisse pas être de nouveau entrée en gardant la lettre

			if strings.ContainsRune(word, letterGiven) { // on teste si la lettre est présente dans le mot avec la fonction string.ContainsRune
				fmt.Println("wright answer, ", letter, "is present in the word\n") // la lettre est présente dans le mot
			} else { // sinon
				attempts--                     // on réduit le nombre d'essais de 1
				nbr := (10 - attempts - 1) * 8 // les actions pour afficher le hangman correctement
				GetHangman(nbr)                // on appelle la fontion pour afficher le hangman
				if attempts > 0 {              // si il reste des chances
					fmt.Println("wrong answer, you still have", attempts, "attempts to discover the word\n") // on dit que c'est faux et on donne le nombre de chances restantes
				}
			}

			printWordGuessStatus(word, wordFoundLetters) // on affiche le nouveau status du mot avec les lettre trouvée ou non

			if allLettersFounds(word, wordFoundLetters) { // si toutes les lettres ont étaient trouvée alors
				fmt.Println("Congratulation, you found the word :", word) // le jeu est gagné on affiche le mot final
				break                                                     // on met fin au programme
			}

			if attempts == 0 { // si le nombre de chances est arrivé à 0 alors
				fmt.Println("Your number of attempts reached 0. The word was :", word) // on dit que c'est un echec et on donne le mot
			}
		}
	} else { // si la personne ne veux pas jouer alors
		fmt.Println("See you next time !") // on lui dit au revoir et on finit le programme
	}
}

func printWordGuessStatus(word string, wordFoundLetters map[rune]bool) { // fonction pour ajouter les lettres dans le tableau
	wordPrinted := ""                 // on initialise une string
	for _, characters := range word { // on parcours le mot
		if wordFoundLetters[characters] { // si le caractère est dans le mot alors
			wordPrinted += string(characters) // on ajoute le caractère dans la string
		} else { // sinon
			wordPrinted += "_" // on met des tirets pour les lettres non devinées
		}
	}
	fmt.Println(wordPrinted) // on imprime le status du mot avec les lettres ajouter
}

func allLettersFounds(word string, wordFoundLetters map[rune]bool) bool { // la fonction pour tester si le mot est diviné
	for _, characters := range word { // on parcours le mot de nouveau
		if !wordFoundLetters[characters] { // si il y a des caractères non ajouter dans le mot alors
			return false // on dit que le mot n'est pas trouvé
		}
	}
	return true // on dit que le mot est trouvé
}

func WordToFind() string { // fonction pour choisir le nombre random et afficher le mot avec
	f, err := os.Open("word.txt")  // on ouvre word.txt
	scanner := bufio.NewScanner(f) // on crée un nouveau scanner de f (word.txt)
	nbrMots := 0                   // on initialise nbrmot a 0

	for scanner.Scan() { // on scan le document
		nbrMots++ // on ajoute le nombre de mot présent dans le document
	}

	if err != nil { // si le fichier ne peut pas être ouvert on donne une erreur
		log.Fatal(err)
	}

	randomNumber := rand.Intn(nbrMots) // on génère un nombre random sur le nombre de mot présent dans le fichier

	return Scan(randomNumber) // on appelle la fonction scan et on lui donne le nombre random
}

func Scan(nbr int) string { // fonction pour prendre un mot dans le texte
	mot := ""                      // string mot
	nbrMots2 := 0                  // int de 0
	f, err := os.Open("word.txt")  // on ouvre word.txt de nouveau
	scanner := bufio.NewScanner(f) // on fait un nouveau scanner

	if err != nil { // on gère les erreurs
		log.Fatal(err)
	}

	for scanner.Scan() { // on scanne le fichier
		nbrMots2++           // le nombre augmente pour chaque mot
		if nbrMots2 == nbr { // quand le nombre de mot atteint le nombre passer en parametre alors
			mot = scanner.Text() // on ajoute le mot à cette ligne dans la variable mot
		}
	}
	return mot // on print le mot
}

type LetterIndices struct { // structure pour crée l'indice
	Letter  string // on initialise un string pour les lettre
	Indices []int  // on fait un tableau de int pour avoir la position
}

func nUniqueRandomLetters(word string) []LetterIndices { // fonction pour crée l'indice
	n := len(word)/2 - 1    // n est la moitier du mot
	var tab []LetterIndices // on appelle la structure pour entrer les lettres
	totalIndices := 0       // on met total indice a 0

	for totalIndices < n { // tant que la moitié du mot n'est pas égale au nombre d'indices
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
func GetHangman(nbr int) { // fonction pour afficher le hangman
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
