package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Year  int
	Month int
	Day   int
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName         FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName:         FullName{FirstName: "Jérôme", LastName: "Golaz"},
		BirthDate:        BirthDate{Year: 2008, Month: 4, Day: 29},
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u2649', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
