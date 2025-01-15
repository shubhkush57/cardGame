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

func newDeckFromFile(filename string) deck{
	bs,err := ioutil.ReadFile(filename)

	if(err != nil) {
		fmt.Println(err)
		os.Exit(1)

	}
	s := strings.Split(string(bs),",")
	return deck(s)
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := range d{
		newPos := r.Intn(len(d)-1) 
		d[i], d[newPos] = d[newPos] ,d[i]

	}
}