package main

type Player struct {
	number int
	name   string
	hand   []Card
}

type Card struct {
	value  string
	symbol string
}
