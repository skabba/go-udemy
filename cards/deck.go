package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	redCardSuits := []string{
		"Ruiten",
		"Harten",
	}
	blackCardSuits := []string{
		"Schoppen",
		"Klaver",
	}
	cardValues := []string{
		"Aas",
		"Twee",
		"Drie",
		"Vier",
		"Vijf",
		"Zes",
		"Zeven",
		"Acht",
		"Negen",
		"Tien",
		"Boer",
		"Vrouw",
		"Heer",
	}

	for _, suit := range redCardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" "+value+" - (Rood)")
		}
	}

	for _, suit := range blackCardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" "+value+" - (Zwart)")
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
