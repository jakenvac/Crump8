package main

import (
	"github.com/JakeHL/crump8"
	"github.com/gdamore/tcell"
)

type chip8input struct {
	screen  tcell.Screen
	KeyChan chan rune
}

func newChip8Input(s tcell.Screen) *chip8input {

	c8out := &chip8input{
		screen:  s,
		KeyChan: make(chan rune),
	}

	return c8out
}

func (c *chip8input) GetKey() byte {
	select {
	case r := <-c.KeyChan:
		if val, ok := crump8.DefaultKeyMap[r]; ok {
			return byte(val)
		}
	default:
		return 16
	}
	return 16
}
