package data

// NewProduct creates a new Product
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

func (p *Product) GetName() string {
	return p.name
}

func NewProduct(name string, color Color, size Size) *Product {
	return &Product{name: name, color: color, size: size}
}

// GetGreenColor returns the green color value
func GetGreenColor() Color {
	return green
}

func GetRedColor() Color {
	return red
}

func GetBlueColor() Color {
	return blue
}

// GetSmallSize returns the small size value
func GetSmallSize() Size {
	return small
}
