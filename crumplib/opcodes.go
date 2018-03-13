package crumplib

// 0NNN Calls RCA 1802 program at address NNN. Not necessary for most ROMs.
func (c *crump8) op0NNN() {

}

// 00E0 Clears the screen
func (c *crump8) op00E0() {

}

// 00EE Returns from a subroutine
func (c *crump8) op00EE() {

}

// 1NNN jumpts to address at NNN
func (c *crump8) op1NNN() {

}

// 2NNN Calls subroutine at NNN
func (c *crump8) op2NNN() {

}

// 3XNN Skips the next instruction if VX = NN
func (c *crump8) op3XNN() {

}

// 4XNN Skips the next instruction if VX != NN
func (c *crump8) op4XNN() {

}

// 5XY0 Skips the next instruction if VX = VY
func (c *crump8) op5XY0() {

}

// 6XNN sets VX to NN
func (c *crump8) op6XNN() {

}

// 7XNN adds NN to VX
func (c *crump8) op7XNN() {

}

// 8XY0 Sets VX to VY
func (c *crump8) op8XY0() {

}

// 8XY1 sets VX to (VX | VY)
func (c *crump8) op8XY1() {

}

// 8XY2 sets VX to (VX & VY)
func (c *crump8) op8XY2() {

}

// 8XY3 sets VX to (VX ^ VY)
func (c *crump8) op8XY3() {

}

// 8XY4 Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
func (c *crump8) op8XY4() {

}

// 8XY5 VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (c *crump8) op8XY5() {

}

// 8XY6 Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift
func (c *crump8) op8XY6() {

}

// 8XY7 Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func (c *crump8) op8XY7() {

}

// 8XYE Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.
func (c *crump8) op8XYE() {

}

// ANNN Sets I to the address NNN
func (c *crump8) opANNN() {

}
