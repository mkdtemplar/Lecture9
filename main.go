package main

import (
	cardDraw "Lecture9/CardDraw"
	"fmt"
)

func main() {

	//dc := make([]cardGame.DeckOfCards, 52)
	var dc2 cardDraw.Dealer = &cardDraw.Cards{}

	cardDraw.DrawAllCards(dc2)
	fmt.Println()

	fmt.Println(cardDraw.DealOneCard(dc2))

	fmt.Println(dc2.Done())

}
