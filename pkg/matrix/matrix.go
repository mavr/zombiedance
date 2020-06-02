package matrix

import (
	"time"

	"github.com/mavr/totalwar/pkg/gameplay"
)

type HealthPoint int

type WorldSpacePiece float32

type WorldSpace interface {
	X() WorldSpacePiece
	Y() WorldSpacePiece
	Distance(aim WorldSpace) WorldSpacePiece
	Move(aim WorldSpace, speed WorldSpacePiece)
}

type WorldObject interface {
	UID() string
	Name() string
	GetCoordinates() WorldSpace
}

type Mover interface {
	GetSpeed() WorldSpacePiece
}

type Enemy interface {
	WorldObject
	Mover
	Run(chStart <-chan struct{}, t *time.Ticker)
	Clan() gameplay.PlayerType
	GetAttackDistance() WorldSpacePiece
	GetDamage() HealthPoint
	Protect(damage HealthPoint) bool
	GetHealth() HealthPoint
	Alive() bool
}

type World interface {
	WorldObject
	Move(obj WorldObject)
	IsAccessibility(uid string) bool
	Attack(aim WorldObject) bool
	GetEnemy() []Enemy
}

type Matrix interface {
	Run()
	GetStartSpace() WorldSpace
	CreateUnit(clan gameplay.PlayerType) Enemy
	GetEnemies(unitUID string) []Enemy
	IsAccessibility(unitUID, aimUID string) bool
	Move(unitUID, aimUID string)
	Attack(unitUID, aimUID string) bool
}
