package crump8

// This file contains the definition of the 'hardware' of the Chip8

import (
	"fmt"
	"math/rand"
	"time"
)

// Chip8 defines the format of the chip8
type Chip8 struct {
	/* HARDWARE */
	// An array to hold all 4096 bytes of ram
	ram [4096]byte
	// gfx is a 2D array of pixels that can be either on or off
	gfx [32][64]bool
	// the graphics renderer
	graphics GraphicsRenderer
	// input: chip8 has 16 inputs that are either on or offl
	input InputReader
	// The currently executing opcode which is 16Bits
	opcode uint16
	// V is the array of 16 general purpose registers V0 - VE with a 16th Carry register
	V [16]byte
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
	// unix time of last cycle
	lastCycle int64
	// clock speed in Hz
	clockSpeed int64

	/* UTILITY */
	// manages execution events such as pause, resume, stop
	eventManager *EventManager
	// pauses execution if true
	shouldPause bool
}

// NewChip8 Creates a new instance of the Crump8 Emulator
func NewChip8(rom []byte) (c *Chip8) {

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	LogWrite("Initializing Chip8")
	c = &Chip8{
		pc:           0x200,
		opcode:       0,
		i:            0,
		stackPointer: 0,
		clockSpeed:   120,
		input:        nil,
	}

	LogWrite("Loading font...")
	// load the font into memory
	for i := 0; i < 80; i++ {
		c.ram[i] = fontSet[i]
	}

	LogWrite("Loading ROM...")
	// load the rom into memory from byte 512 onwards
	for i := 0; i < len(rom); i++ {
		c.ram[i+512] = rom[i]
	}

	return
}

// SetInput Sets the input device on the Chip8
func (c *Chip8) SetInput(input InputReader) {
	c.input = input
}

// SetGraphics sets the graphics output device on the Chip8
func (c *Chip8) SetGraphics(gfx GraphicsRenderer) {
	c.graphics = gfx
}

// SetEventManager sets the event manager for the chip8
func (c *Chip8) SetEventManager(evtMgr *EventManager) {
	c.eventManager = evtMgr
}

// Cycle emulates one cycle of the cpu and returns the delta time since the last cycle
func (c *Chip8) Cycle() int64 {

	// TODO might have to come back and rethink this
	if c.shouldPause {
		return 0
	}

	// As opcodes are two bytes long we fetch two bytes from memory and merge them by shifting the first byte left 8 bits and or-ing it with the next byte
	c.opcode = uint16(c.ram[c.pc])<<8 | uint16(c.ram[c.pc+1])

	// fetch execute
	switch c.opcode & 0xF000 {
	case 0x1000:
		c.op1NNN()
	case 0x2000:
		c.op2NNN()
	case 0x3000:
		c.op3XNN()
	case 0x4000:
		c.op4XNN()
	case 0x5000:
		c.op5XY0()
	case 0x6000:
		c.op6XNN()
	case 0x7000:
		c.op7XNN()
	case 0x9000:
		c.op9XY0()
	case 0xA000:
		c.opANNN()
	case 0xB000:
		c.opBNNN()
	case 0xC000:
		c.opCXNN()
	case 0xD000:
		c.opDXYN()
	case 0x8000: // 8 series of opcodes
		switch c.opcode & 0x000F {
		case 0x0000:
			c.op8XY0()
		case 0x0001:
			c.op8XY1()
		case 0x0002:
			c.op8XY2()
		case 0x0003:
			c.op8XY3()
		case 0x0004:
			c.op8XY4()
		case 0x0005:
			c.op8XY5()
		case 0x0006:
			c.op8XY6()
		case 0x0007:
			c.op8XY7()
		case 0x000E:
			c.op8XYE()
		}
	case 0xE000: // E series of opcodes
		switch c.opcode & 0x000F {
		case 0x000E:
			c.opEX9E()
		case 0x0001:
			c.opEXA1()
		}
	case 0xF000: // F series of opcodes
		switch c.opcode & 0x00FF {
		case 0x0007:
			c.opFX07()
		case 0x000A:
			c.opFX0A()
		case 0x0015:
			c.opFX15()
		case 0x0018:
			c.opFX18()
		case 0x001E:
			c.opFX1E()
		case 0x0029:
			c.opFX29()
		case 0x0033:
			c.opFX33()
		case 0x0055:
			c.opFX55()
		case 0x0065:
			c.opFX65()
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

	if c.delayTimer > 0 {
		c.delayTimer--
	}
	if c.soundTimer > 0 {
		c.soundTimer--
	}

	if c.draw {
		c.graphics.Render(c.gfx)
		c.draw = false
	}

	delta := time.Now().UnixNano() - c.lastCycle
	c.lastCycle = time.Now().UnixNano()
	return delta / int64(time.Millisecond)
}

// Run starts the chip8
func (c *Chip8) Run() {
	for {
		select {
		case <-c.eventManager.Pause:
			// TODO add pause
		case <-c.eventManager.Resume:
			// TODO add resume
		case <-c.eventManager.Stop:
			return
		default:
			loopTime := 1000/c.clockSpeed - c.Cycle()
			if loopTime < 0 {
				loopTime = 0
			}
			time.Sleep(time.Duration(loopTime) * time.Millisecond)
		}
	}
}
