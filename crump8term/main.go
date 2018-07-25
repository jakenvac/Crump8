package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/JakeHL/crump8/crumplib"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

func main() {
	//First 5 lines are just for a vs code debugging hack
	fmt.Println("Press enter to continue debugging") // 1
	fmt.Scanln()                                     // 2
	cmd := exec.Command("clear")                     // 3
	cmd.Stdout = os.Stdout                           // 4
	cmd.Run()                                        // 5

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = screen.Init(); err != nil {
		panic(err)
	}

	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()
	screen.Sync()

	romPath := "../roms/tetris"

	rom, err := ioutil.ReadFile(romPath)
	if err != nil {
		panic(err)
	}

	c8 := crumplib.NewCrump8(rom)

	quit := make(chan struct{})
	keys := make(chan tcell.Key)

	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					close(quit)
					return
				case tcell.KeyCtrlL:
					screen.Sync()
				default:
					keys <- ev.Key()
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	gameStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlue)
	gameView := views.NewViewPort(screen, 1, 1, 64, 32)

	debugStyle := tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tcell.ColorBlack)
	debugView := views.NewViewPort(screen, 66, 1, 30, 32)
	debugView.Fill(' ', debugStyle)
	debugList := views.NewBoxLayout(views.Vertical)
	debugList.SetView(debugView)

loop:
	for {
		timeToWait := 1000/60 - c8.Cycle()
		if timeToWait < 0 {
			timeToWait = 0
		}

		if c8.ShouldDraw() {
			for y := range c8.Gfx {
				for x := range c8.Gfx[y] {
					if c8.Gfx[y][x] {
						gameView.SetContent(x, y, '█', nil, gameStyle)
					} else {
						gameView.SetContent(x, y, ' ', nil, gameStyle)
					}
				}
			}
		}

		opcode := c8.GetOpCode()
		opCodeHex := fmt.Sprintf("%X", opcode)
		opCodeInt := fmt.Sprintf("%v", opcode)
		debugLine("opcode: ", opCodeInt, opCodeHex, debugList, debugStyle)

		screen.Show()

		select {
		case <-quit:
			break loop
		case k := <-keys:
			if k == tcell.KeyBackspace {
				var key tcell.Key
				for key != tcell.KeyBackspace {
					ev := screen.PollEvent()
					switch ev := ev.(type) {
					case *tcell.EventKey:
						key = ev.Key()
					}
				}

			}
		default:
		}

		time.Sleep(time.Duration(timeToWait) * time.Millisecond)
	}
}

var debugLineCount int

func debugLine(left, center, right string, b *views.BoxLayout, style tcell.Style) {

	widgets := b.Widgets()
	widgetLength := len(widgets)
	if widgetLength >= 32 {
		toRemove := &widgets[0]
		b.RemoveWidget(*toRemove)
		b.Draw()
	}

	// vp := views.NewViewPort(v, 0, debugLineCount, 10, 1)
	tb := views.NewTextBar()
	tb.SetStyle(style)
	tb.SetLeft(left, style)
	tb.SetCenter(center, style)
	tb.SetRight(right, style)
	b.InsertWidget(32, tb, 0)
	tb.Draw()

	debugLineCount++
}
