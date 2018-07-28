package main

import (
	"io/ioutil"

	"github.com/JakeHL/crump8"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

func main() {

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = screen.Init(); err != nil {
		panic(err)
	}

	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tcell.ColorGreen))
	screen.Clear()

	gameStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlue)
	gameView := views.NewViewPort(screen, 1, 1, 64, 32)
	gameView.Fill(' ', gameStyle)

	display := newChip8Display(gameView, screen)
	display.SetStyle(gameStyle)

	input := newChip8Input(screen)

	evtMgr := crump8.NewEventManager()

	romPath := "../../roms/brix"
	rom, err := ioutil.ReadFile(romPath)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					close(evtMgr.Stop)
					return
				default:
					input.KeyChan <- ev.Rune()
				}
			case *tcell.EventResize:
				gameView.Resize(1, 1, 64, 32)
			}
		}
	}()

	c8 := crump8.NewChip8(rom)
	c8.SetGraphics(display)
	c8.SetInput(input)
	c8.SetEventManager(evtMgr)
	c8.Run()
}
