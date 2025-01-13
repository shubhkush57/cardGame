package main

import "fmt"

type deck []string

func newDec() deck{
	cards := deck{}
    cardSuits := deck {"Spades", "Diamonds","Hearts","Clubs"}
	cardValues := deck {"Ace","Two","Three","Four"}
	for _, suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards, value + " of "+ suit)
		}
	}
	return cards
}

func (cards deck) print() {
	// this is reciever function
	// every variable of type deck can use this function 
	for i, card := range cards {
		fmt.Println(i,card)
	}
}