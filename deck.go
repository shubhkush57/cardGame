package main

import "fmt"

type deck []string

func (cards deck) print() {
	// this is reciever function
	// every variable of type deck can use this function 
	for i, card := range cards {
		fmt.Println(i,card)
	}
}