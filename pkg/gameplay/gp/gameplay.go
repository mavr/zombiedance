package gp

import (
	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/gameplay/cmdgp"
)

func New(gp gameplay.GamePlayType) gameplay.Gameplay {
	switch gp {
	case gameplay.GamePlayZombiVsOpposition:
		return cmdgp.New()
	default:
		return nil
	}
}
