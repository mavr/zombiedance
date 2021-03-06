package zombie

import (
	"math/rand"
	"time"

	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/matrix"
	"github.com/mavr/totalwar/pkg/unit/unit"
)

const (
	Damage         matrix.HealthPoint     = 15
	Speed          matrix.WorldSpacePiece = 3
	AttackDistance matrix.WorldSpacePiece = 5
	DefaultHealth  matrix.HealthPoint     = 200
)

type Zombie struct {
	*unit.Unit
	aim matrix.Enemy
}

func New(world matrix.World, chStart chan struct{}) *Zombie {
	t := &Zombie{
		Unit: unit.New(world, gameplay.Zombie, unit.UnitSkills{
			AttackDistance: AttackDistance,
			Damage:         Damage,
			DefaultHealt:   DefaultHealth,
			Speed:          Speed,
		}),
	}

	return t
}

func (u *Zombie) Run(chTick <-chan struct{}, t *time.Ticker) {
	for range chTick {
		if u.GetHealth() <= 0 {
			return
		}
		u.Action()
	}
}

func (u *Zombie) Action() {
	if u.aim == nil || !u.aim.Alive() {
		enemies := u.GetEnemy()
		if len(enemies) == 0 {
			return
		}

		lenem := make([]matrix.Enemy, 0, len(enemies))
		for _, e := range enemies {
			if e.Alive() && e.Clan() != u.Clan() {
				lenem = append(lenem, e)
			}
		}

		if len(lenem) == 0 {
			return
		}

		u.aim = lenem[rand.Intn(len(lenem))]
	}

	if !u.IsAccessibility(u.aim.UID()) {
		u.Move(u.aim)
		return
	}

	if u.Attack(u.aim) {
		u.aim = nil
	}
}
