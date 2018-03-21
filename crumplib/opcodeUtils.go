package crumplib

func getX(opcode uint16) byte {
	return byte((opcode & 0x0f00) >> 8)
}

func getY(opcode uint16) byte {
	return byte((opcode & 0x00f0) >> 4)
}

func getXY(opcode uint16) (byte, byte) {
	return getX(opcode), getY(opcode)
}

func getMsb(byteval byte) byte {
	var old byte = 0x1
	var new byte = 0x0
	var msb byte = 0x0
	for new < byteval {
		new = old << 1
		if new > byteval {
			msb = old
			break
		}
		old = new
	}
	return msb
}
