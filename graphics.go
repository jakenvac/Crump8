package crump8

// GraphicsRenderer determins the output method of the chip8
type GraphicsRenderer interface {
	Render(display [32][64]bool)
}
