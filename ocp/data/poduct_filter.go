package data

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

type Filter struct {
	//
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.Color
}

type SizeSpecification struct {
	Size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.Size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
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

func (f *Filter) FilterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(
	products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}
