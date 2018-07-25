package crumplib

import (
	"math/rand"
)

// 0NNN Calls RCA 1802 program at address NNN. Not necessary for most ROMs.
func (c *Crump8) op0NNN() {

}

// 00E0 Clears the screen
func (c *Crump8) op00E0() {
	c.Gfx = [32][64]bool{}
	c.pc += 2
}

// 00EE Returns from a subroutine
func (c *Crump8) op00EE() {
	c.stackPointer--
	c.pc = c.stack[c.stackPointer]
	c.pc += 2
}

// 1NNN jumps to address at NNN
func (c *Crump8) op1NNN() {
	c.pc = c.opcode & 0x0FFF
}

// 2NNN Calls subroutine at NNN
func (c *Crump8) op2NNN() {
	c.stack[c.stackPointer] = c.pc
	c.stackPointer++
	c.pc = c.opcode & 0x0FFF
}

// 3XNN Skips the next instruction if VX = NN
func (c *Crump8) op3XNN() {
	vx, _ := getXY(c.opcode)
	nn := byte(c.opcode) & 0x00FF
	if nn == c.V[vx] {
		c.pc += 4
	} else {
		c.pc += 2
	}
}

// 4XNN Skips the next instruction if VX != NN
func (c *Crump8) op4XNN() {
	vx, _ := getXY(c.opcode)
	nn := byte(c.opcode) & 0x00FF
	if nn != c.V[vx] {
		c.pc += 4
	} else {
		c.pc += 2
	}
}

// 5XY0 Skips the next instruction if VX = VY
func (c *Crump8) op5XY0() {
	vx, vy := getXY(c.opcode)
	if c.V[vx] == c.V[vy] {
		c.pc += 4
	} else {
		c.pc += 2
	}
}

// 6XNN sets VX to NN
func (c *Crump8) op6XNN() {
	vx, _ := getXY(c.opcode)
	nn := byte(c.opcode) & 0x00FF
	c.V[vx] = nn
	c.pc += 2
}

// 7XNN adds NN to VX
func (c *Crump8) op7XNN() {
	vx, _ := getXY(c.opcode)
	nn := byte(c.opcode) & 0x00FF
	c.V[vx] += nn
	c.pc += 2
}

// 8XY0 Sets VX to VY
func (c *Crump8) op8XY0() {
	vx, vy := getXY(c.opcode)
	c.V[vx] = c.V[vy]
	c.pc += 2
}

// 8XY1 sets VX to (VX | VY)
func (c *Crump8) op8XY1() {
	vx, vy := getXY(c.opcode)
	c.V[vx] |= c.V[vy]
	c.pc += 2
}

// 8XY2 sets VX to (VX & VY)
func (c *Crump8) op8XY2() {
	vx, vy := getXY(c.opcode)
	c.V[vx] &= c.V[vy]
	c.pc += 2
}

// 8XY3 sets VX to (VX ^ VY)
func (c *Crump8) op8XY3() {
	vx, vy := getXY(c.opcode)
	c.V[vx] ^= c.V[vy]
	c.pc += 2
}

// 8XY4 Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
func (c *Crump8) op8XY4() {
	vx, vy := getXY(c.opcode)
	if c.V[vx] > (0xFF - c.V[vy]) {
		c.V[0xf] = 1
	} else {
		c.V[0xf] = 0
	}
	c.V[vx] += c.V[vy]
	// Increment to the next opcode
	// We increment by two as each opcode is two bytes long
	// if we incremented by 1 we'd be on the second half of the same opcode
	c.pc += 2
}

// 8XY5 VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (c *Crump8) op8XY5() {
	vx, vy := getXY(c.opcode)
	if c.V[vy] > c.V[vx] {
		c.V[0xf] = 0
	} else {
		c.V[0xf] = 1
	}
	c.V[vx] -= c.V[vy]
	c.pc += 2
}

// 8XY6 Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift
func (c *Crump8) op8XY6() {
	vx, _ := getXY(c.opcode)
	// (x & -x) == x & (255 - x + 1) == x & ~x + 1
	c.V[0xf] = c.V[vx] & -c.V[vx]
	c.V[vx] = c.V[vx] >> 1
	c.pc += 2
}

// 8XY7 Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (c *Crump8) op8XY7() {
	vx, vy := getXY(c.opcode)
	if c.V[vx] > c.V[vy] {
		c.V[0xf] = 0
	} else {
		c.V[0xf] = 1
	}
	c.V[vx] = c.V[vy] - c.V[vx]
	c.pc += 2
}

// 8XYE Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.
func (c *Crump8) op8XYE() {
	vx, _ := getXY(c.opcode)
	c.V[0xf] = c.V[vx] & 0x80
	c.V[vx] = c.V[vx] << 1
	c.pc += 2
}

// 9XY0 skips the next instructiion if VX != VY
func (c *Crump8) op9XY0() {
	vx, vy := getXY(c.opcode)
	if c.V[vx] != c.V[vy] {
		c.pc += 4
	} else {
		c.pc += 2
	}
}

// ANNN Sets I to the address NNN
func (c *Crump8) opANNN() {
	c.i = c.opcode & 0x0FFF
	c.pc += 2
}

// BNNNN Jumps to the address NNN plus V0
func (c *Crump8) opBNNN() {
	c.pc = uint16(c.opcode&0x0FFF) + uint16(c.V[0])
}

// CXNN Sets VX to (Rand & NN)
func (c *Crump8) opCXNN() {
	randomVal := byte(rand.Int31n(255)) & byte((c.opcode & 0xFF))
	vx, _ := getXY(c.opcode)
	c.V[vx] = randomVal
	c.pc += 2
}

// DXYN draws a sprite at coordinate VX, VY that has a width of 8 pixels and a height of N pixels. draws the bit coded sprite starting at I. value of I doesn't change. If any pixels are set from on to off, VF is set to 1, 0 if not
var count int

func (c *Crump8) opDXYN() {

	vx, vy := getXY(c.opcode)
	x := c.V[vx]
	y := c.V[vy]

	var width uint16 = 8
	height := c.opcode & 0xF

	spritePixels := c.ram[c.i : c.i+height]

	c.V[0xF] = 0x0

	for i, b := range spritePixels {
		xpos := (uint16(i) % width) + uint16(x)
		ypos := (uint16(i) / width) + uint16(y)
		xoffset := uint16(0x80 >> xpos)

		if int(ypos) < len(c.Gfx) && int(xpos) < len(c.Gfx[ypos]) {
			oldPixelVal := &c.Gfx[ypos][xpos]
			if *oldPixelVal {
				c.V[0xF] = 0x1
			} else {
				c.V[0xF] = 0x0
			}

			*oldPixelVal = (((uint16(b) & xoffset) >> (width - xpos)) ^ 0x1) == 1
		}
	}

	c.draw = true
	c.pc += 2
}

// EX9E Skips the next instruction if the key stored in VX is pressed
func (c *Crump8) opEX9E() {
	vx, _ := getXY(c.opcode)
	if c.Input[c.V[vx]] {
		c.pc += 4
	} else {
		c.pc += 2
	}
}

// EXA1 Skips the next instruction if the key stored in VX is not pressed
func (c *Crump8) opEXA1() {
	vx, _ := getXY(c.opcode)
	if !c.Input[c.V[vx]] {
		c.pc += 4
	} else {
		c.pc += 2
	}
}

// FX07 Sets VX to the value of the delay timer
func (c *Crump8) opFX07() {
	vx, _ := getXY(c.opcode)
	c.V[vx] = c.delayTimer
	c.pc += 2
}

// waits for a keypress
func (c *Crump8) opFX0A() {
	var keyPressed bool
	var keyIndex uint16

	for !keyPressed {
		for i := range c.Input {
			keyPressed = c.Input[i]
			if keyPressed {
				keyIndex = uint16(i)
				break
			}
		}
	}

	if keyPressed {
		vx, _ := getXY(c.opcode)
		c.V[vx] = byte(keyIndex)
		c.pc += 2
	}
}

// FX15 Sets the timer to the value of VX
func (c *Crump8) opFX15() {
	vx, _ := getXY(c.opcode)
	c.delayTimer = c.V[vx]
	c.pc += 2
}

// FX18 sets the sound timer to VX
func (c *Crump8) opFX18() {
	vx, _ := getXY(c.opcode)
	c.soundTimer = c.V[vx]
	c.pc += 2
}

// FX1E adds VX to I
func (c *Crump8) opFX1E() {
	vx, _ := getXY(c.opcode)
	c.i += uint16(c.V[vx])
	c.pc += 2
}

// FX29 sets I to the location of the sprite for the character in VX
func (c *Crump8) opFX29() {
	vx, _ := getXY(c.opcode)
	c.i = uint16(5 * c.V[vx])
	c.pc += 2
}

// FX33 Stores the binary coded decimal of VX with the 100 in I, 10 in I + 1, 1 in I + 2
func (c *Crump8) opFX33() {
	vx, _ := getXY(c.opcode)
	c.ram[c.i] = c.V[vx] / 100
	c.ram[c.i+1] = (c.V[vx] / 10) % 10
	c.ram[c.i+2] = c.V[vx] % 10
	c.pc += 2
}

// FX55 Stores from V0 to VX (inclusive) in memory starting at address I. I is increased for each value written
func (c *Crump8) opFX55() {
	vx, _ := getXY(c.opcode)
	vs := c.V[0 : vx+1]

	for v := range vs {
		c.ram[c.i] = vs[v]
		c.i++
	}

	c.pc += 2
}

// FX65 Fills V0 to VX with the values from memory starting at address I. I is increased with each value
func (c *Crump8) opFX65() {
	vx, _ := getXY(c.opcode)
	for count := 0; count <= int(vx); count++ {
		c.V[count] = c.ram[c.i]
		c.i++
	}
	c.pc += 2
}
