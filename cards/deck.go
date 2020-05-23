package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	//convert the deck to a byte array to write to disk
	ioutil.WriteFile(filename, []byte(d))

}
