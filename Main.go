package main

import "fmt"

func main() {
	movie := NewMovie(LOTR{})
	movie.PerformQuote()
	movie.SetMovie(HarryPotter{})
	movie.PerformQuote()
}

type QuoteReciter interface {
	Recite()
}

type LOTR struct{}

func (spm LOTR) Recite() {
	fmt.Println("Not all those who wander are lost.” — Bilbo Baggins")
}

type HarryPotter struct{}

func (sum HarryPotter) Recite() {
	fmt.Println("If you can do one thing, hone it to perfection. Hone it to the utmost limit - Zenitsu")
}

type movie struct {
	QuoteReciter QuoteReciter
}

func NewMovie(qr QuoteReciter) *movie {
	return &movie{
		QuoteReciter: qr,
	}
}

func (t *movie) PerformQuote() {
	t.QuoteReciter.Recite()
}

func (t *movie) SetMovie(qr QuoteReciter) {
	t.QuoteReciter = qr
}
