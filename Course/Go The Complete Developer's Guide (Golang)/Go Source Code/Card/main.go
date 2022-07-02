package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// card := newCard()

	card := []string{"asdfas", newCard()}
	card = append(card, "I love you")
	fmt.Println(card)

	for i, car := range card {
		fmt.Println(i, car)
	}
}

// basic data types: bool string int float64

func newCard() string {
	return "hello mike!"
}
