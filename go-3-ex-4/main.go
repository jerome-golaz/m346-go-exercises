package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	// TODO: Loop over suits and ranks to output all combinations.
	for iSuits := 0; iSuits < len(suits); iSuits++ {
		for iRanks := 0; iRanks < len(ranks); iRanks++ {
			fmt.Printf("%c:%c\n", suits[iSuits], ranks[iRanks])
		}
	}

	// TODO: delete this line afterwards
	// fmt.Println(suits, ranks)
}
