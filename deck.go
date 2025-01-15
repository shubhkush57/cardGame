package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func deal(d deck,handSize int)(deck,deck){
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string{
	// deck to string
	return strings.Join(d,",")

}
// error is a type here.
func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte (d.toString()),0666)

}