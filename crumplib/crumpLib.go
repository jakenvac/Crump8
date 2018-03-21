package crumplib

import (
	"fmt"
	"math/rand"
	"time"
)

// This file contains the definition of the 'hardware' of the Chip8

// Crump8 defines the format of the chip8
type Crump8 struct {
	// An array to hold all 4096 bytes of ram
	ram [4096]byte
	// 2D array of pixels that can be either on or off
	gfx [64][32]bool
	// chip8 has 16 inputs that are either on or off
	input [16]bool
	// The currently executing opcode which is 16Bits
	opcode uint16
	// the array of 16 general purpose registers V0 - VE with a 16th Carry register
	v [16]byte
	// the index register that can have a max value of 0xFFF (12 bits)
	i uint16
	// the program counter register that can have a max value of 0xFFF (12 bits)
	pc uint16
	// two timer registers that count down at 60hz if they are set above 0
	delayTimer, soundTimer byte
	// 16 level stack of 16bit values
	stack [16]uint16
	// points to the last location on the stack
	stackPointer uint16
	// determines if the graphics should be updated
	draw bool
}

// NewCrump8 Creates a new instance of the Crump8 Emulator
func NewCrump8(rom []byte) *Crump8 {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	LogWrite("Initializing Chip8")
	c := &Crump8{
		pc:           0x200,
		opcode:       0,
		i:            0,
		stackPointer: 0,
	}

	LogWrite("Loading font...")
	// load the font into memory
	for i := 0; i < 80; i++ {
		c.ram[i] = fontSet[i]
	}

	LogWrite("Loading ROM")
	// load the rom into memory from byte 512 onwards
	for i := 0; i < len(rom); i++ {
		c.ram[i+512] = rom[i]
	}

	LogWrite("Chip8 Initialized. Printing memory dump:")
	//LogWrite(string(c.ram[:len(c.ram)]))

	return c
}

// Cycle emulates one cycle of the cpu and returns the delta time between each cycle
func (c *Crump8) Cycle() float32 {
	// As opcodes are two bytes long we fetch two bytes from memory and merge them by shifting the first byte left 8 bits and or-ing it with the next byte
	c.opcode = uint16(c.ram[c.pc])<<8 | uint16(c.ram[c.pc+1])

	// TODO Remove
	c.v[2] = 0xFF
	c.v[3] = 0x01
	c.opcode = 0x8234

	// fetch execute
	switch c.opcode & 0xF000 {
	case 0xA000:
		c.opANNN()
	case 0x2000:
		c.op2NNN()
	case 0x1000:
		c.op1NNN()
	case 0x8000:
		switch c.opcode & 0x000F {
		case 0x0000:
			c.op8XY0()
		case 0x0004:
			c.op8XY4()
		}
	case 0x0000: // For more opcodes
		switch c.opcode & 0x000F {
		case 0x0000: // Clear the screen
			c.op00E0()
		case 0x000E:
			c.op00EE()
		}
	default:
		msg := fmt.Sprintf("Invalid opcode: %x", c.opcode)
		LogWrite(msg)
	}

	return 0
}

// Keys will detect key presses on the Chip 8
func (c *Crump8) Keys() {
	//fmt.Print("Not Implemented")
}
