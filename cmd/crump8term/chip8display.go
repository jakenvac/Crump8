package main

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type chip8display struct {
	viewPort *views.ViewPort
	screen   tcell.Screen
	theme    tcell.Style
}

// creates a new instanxce of the chip8display
func newChip8Display(v *views.ViewPort, s tcell.Screen) (c8d *chip8display) {
	c8d = &chip8display{
		viewPort: v,
		screen:   s,
		theme:    tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorRed),
	}
	c8d.viewPort.Fill(' ', c8d.theme)
	return
}

func (c8d *chip8display) SetStyle(style tcell.Style) {
	c8d.theme = style
	c8d.viewPort.Fill(' ', c8d.theme)
}

func (c8d *chip8display) Render(gfx [32][64]bool) {
	for y := range gfx {
		for x := range gfx[y] {
			if gfx[y][x] {
				c8d.viewPort.SetContent(x, y, '█', nil, c8d.theme.Foreground(tcell.ColorBlue))
			} else {
				c8d.viewPort.SetContent(x, y, ' ', nil, c8d.theme)
			}
		}
	}
	c8d.screen.Show()
}
