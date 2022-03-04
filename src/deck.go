package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

type s string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "Jack", "Q", "K"}

	for _, suit := range cardValues {
		for _, value := range cardSuits {

			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	fmt.Println(fmt.Println(d[:handSize]))
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) savetoFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//error handling
		//option 1 log the error and return call to a new deck
		//option 2 log error and quit program
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") //Two of hearts,...,..
	return deck(s)

}
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	//This is a seed that will always give a random set of numbers

	r := rand.New(source)
	//passing the seed into the random number generator

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
		//swap i and new postiion around
		//doesnt randomize after everys run

	}
}
