package main

//OCP
//open for extension, closed for modification
//Specification pattern

//Scenario: Online store, selling widgets

/*
Go special keyword: iota
in a block, will increment
small size = iota (0)
medium (1)
large (2)
*/

import (
	"fmt"

	"github.com/q10357/solid-principles-go/solid/data"
)

func main() {
	pear := data.NewProduct(
		"Pear",
		data.GetGreenColor(),
		data.GetSmallSize(),
	)
	flower := data.NewProduct(
		"Flower",
		data.GetRedColor(),
		data.GetSmallSize(),
	)

	products := []data.Product{*pear, *flower}
	fmt.Printf("Green Products (old):\n")
	f := data.Filter{}
	for _, v := range f.FilterByColor(products, data.GetGreenColor()) {
		fmt.Printf(" - %s is green\n", v.GetName())
	}

	fmt.Printf("Green Products (new):\n")
	greenSpec := data.ColorSpecification{}
	greenSpec.Color = data.GetGreenColor()

	bf := data.BetterFilter{}

	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.GetName())
	}

}
