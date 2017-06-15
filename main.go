package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("Crump8 - Chip8 Emulator (C) JakeHL 2017 - https://github.com/jakehl/crump8")

	rom, err := ioutil.ReadFile("./roms/TETRIS")
	if err != nil {
		panic(err)
	}

	chip := newChip8(rom)

	// TODO remove this to enter the loop
	os.Exit(1)
	// Chip 8 loop
	for {
		chip.cycle()

		if chip.draw {
			// TODO drawGraphics()
		}

		chip.keys()
	}

}
