package cardDraw

import (
	//"errors"
	//"fmt"
	cardGame "Lecture9/CardGame"
)

type Dealer interface {
	Deal() *cardGame.DeckOfCards
	DealOneCard() *cardGame.DeckOfCards
	Done() bool
}

func DealOneCard(c Dealer) {

	c.DealOneCard()
}

func DrawAllCards(dealer Dealer) {
	dealer.Deal()
}
