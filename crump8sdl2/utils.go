package main

import (
	"bufio"
	"os"
)

func readKey() rune {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return []rune(input)[0]
}
