package shape

type Square struct {
	Rectangle
}

// Enforce width == height in square
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}
