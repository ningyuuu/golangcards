package main

import "strings"

type card struct {
	value string
	suit  string
}

func (c card) toString() string {
	return c.value + " of " + c.suit
}

func newCardFromString(s string) card {
	valSuit := strings.Split(s, " of ")
	return card{valSuit[0], valSuit[1]}
}
