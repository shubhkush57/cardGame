package main

import "fmt"

func main() {
	cards := newDec()
	// cards.print()

	// hand, remainingDec := deal(cards, 1)
	// fmt.Println(hand,remainingDec)
	fmt.Println(cards.toString())
	cards.saveToFile("myCards")

}