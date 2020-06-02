package simplecli

import (
	"fmt"

	"github.com/mavr/totalwar/pkg/game/simplegame"
	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/gameplay/gp"
	"github.com/mavr/totalwar/pkg/matrix/gamematrix"
	"github.com/nsf/termbox-go"
)

type SimpleCliUI struct {
}

func NewSimpleCliUI() *SimpleCliUI {
	return &SimpleCliUI{}
}

func (ui *SimpleCliUI) Start() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	for {
		fmt.Println("Tap space key for start or ESC to exit.")

		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyEsc {
			return
		}

		if ev.Key == termbox.KeySpace {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()

			ui.startGame()
		}
	}
}

func (ui *SimpleCliUI) startGame() {
	matrix := gamematrix.New(
		gp.New(gameplay.GamePlayZombiVsOpposition),
	)
	game := simplegame.NewGame(matrix)

	game.Run()
}
