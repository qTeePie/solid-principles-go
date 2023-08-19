package ocp

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
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}
