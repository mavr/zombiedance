package unit

import (
	"sync"

	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/matrix"
)

type UnitSkills struct {
	AttackDistance matrix.WorldSpacePiece
	Speed          matrix.WorldSpacePiece
	Damage         matrix.HealthPoint

	DefaultHealt matrix.HealthPoint
}

type Unit struct {
	health  matrix.HealthPoint
	mHealth sync.Mutex

	clan gameplay.PlayerType

	matrix.World

	skills UnitSkills
}

func New(world matrix.World, clan gameplay.PlayerType, skills UnitSkills) *Unit {
	u := &Unit{
		health:  skills.DefaultHealt,
		mHealth: sync.Mutex{},
		World:   world,
		clan:    clan,
		skills:  skills,
	}

	return u
}

func (u *Unit) Protect(points matrix.HealthPoint) bool {
	u.mHealth.Lock()
	defer u.mHealth.Unlock()

	u.health -= points
	if u.health < 0 {
		u.health = 0
	}

	return u.health == 0
}

func (u *Unit) Alive() bool {
	return u.GetHealth() > 0
}

func (u *Unit) GetHealth() matrix.HealthPoint {
	return u.health
}

func (u *Unit) Clan() gameplay.PlayerType {
	return u.clan
}

func (u *Unit) GetAttackDistance() matrix.WorldSpacePiece {
	return u.skills.AttackDistance
}

func (u *Unit) GetDamage() matrix.HealthPoint {
	return u.skills.Damage
}

func (u *Unit) GetSpeed() matrix.WorldSpacePiece {
	return u.skills.Speed
}
