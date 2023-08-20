package shape

/*
Code below will work
*/
type Square2 struct {
	size int
}

func (s *Square2) Rectangle() *Rectangle {
	r := Rectangle{}
	r.SetHeight(s.size)
	r.SetWidth(s.size)
	return &r
}

func (s *Square2) SetSize(size int) {
	s.size = size
}

/*
This square violates LSP, width and height is same, but when setting width to another
value (or height), they may differ, resulting in wrong calculations
*/
type Square struct {
	Rectangle
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// Enforce width == height in square
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}
