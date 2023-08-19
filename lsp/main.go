package main

import "github.com/q10357/solid-principles-go/solid/lsp/data/shape"

//Liskov Substitution Principle

func main() {
	rc := &shape.Rectangle{}
	rc.SetHeight(2)
	rc.SetWidth(3)

	shape.UseIt(rc)
}
