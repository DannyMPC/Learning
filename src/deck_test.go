package main

import (
	"testing"
)

func TestnewDeck(t *testing.T) {
	//Test deck func Testnewdeck is capitel letter

	d := newDeck()

	if len(d) != 5 {
		t.Errorf("Error Expected length of 5 but got", len(d))
	}

}
