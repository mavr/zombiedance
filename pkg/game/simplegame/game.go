package simplegame

import (
	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/matrix"
)

type Game struct {
	matrix matrix.Matrix
}

func NewGame(matrix matrix.Matrix) *Game {
	return &Game{
		matrix: matrix,
	}
}

func (g *Game) Run() {
	for i := 0; i < 1500; i++ {
		g.matrix.CreateUnit(gameplay.Zombie)
	}

	for i := 0; i < 700; i++ {
		g.matrix.CreateUnit(gameplay.Opposition)
	}

	g.matrix.Run()
}
