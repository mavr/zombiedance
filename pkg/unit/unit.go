package unit

import (
	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/matrix"
	"github.com/mavr/totalwar/pkg/unit/trooper"
	"github.com/mavr/totalwar/pkg/unit/zombie"
)

func New(gp gameplay.PlayerType, world matrix.World) matrix.Enemy {
	switch gp {
	case gameplay.Opposition:
		return trooper.New(world, nil)
	case gameplay.Zombie:
		return zombie.New(world, nil)
	default:
		return nil
	}
}
