package main

import (
	"fmt"
	"io/ioutil"

	"github.com/JakeHL/crump8/crumplib"
)

func main() {
	fmt.Print("Crump8 - Chip8 Emulator (C) JakeHL 2017 - https://github.com/jakehl/crump8")

	rom, err := ioutil.ReadFile("../roms/TETRIS")
	if err != nil {
		panic(err)
	}

	//crumplib.LogWriter = os.Stdout

	c8 := crumplib.NewCrump8(rom)

	for {

	}
}
