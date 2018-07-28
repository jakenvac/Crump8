package crump8

// InputReader handles the input to the Chip8
type InputReader interface {
	GetKey() byte
}

// DefaultKeyMap configures a default mapping of keyboard keys to Chip8 Inputs
var DefaultKeyMap = map[rune]int8{
	'1': 0, '2': 1, '3': 2, '4': 3,
	'q': 4, 'w': 5, 'e': 6, 'r': 7,
	'a': 8, 's': 9, 'd': 10, 'f': 11,
	'z': 12, 'x': 13, 'c': 14, 'v': 15,
	'Q': 4, 'W': 5, 'E': 6, 'R': 7,
	'A': 8, 'S': 9, 'D': 10, 'F': 11,
	'Z': 12, 'X': 13, 'C': 14, 'V': 15,
}
