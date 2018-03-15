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
