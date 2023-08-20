package main

import "github.com/q10357/solid-principles-go/solid/lsp/data/shape"

//Liskov Substitution Principle
//This is a difficult subject in go, but trying to demonstrate it

func main() {
	rc := &shape.Rectangle{}
	rc.SetHeight(2)
	rc.SetWidth(3)
	shape.UseIt(rc)

	sq := shape.NewSquare(5)
	shape.UseIt(sq)

	sqRight := shape.Square2{}
	sqRight.SetSize(10)
	squareRect := sqRight.Rectangle()
	shape.UseIt(squareRect)
}
