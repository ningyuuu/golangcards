package main

func main() {
	cards := newDeckFromFile("mycards.txt")
	cards.shuffle()
	cards.print()
}
