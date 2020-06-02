package simplematrix

import (
	"github.com/google/uuid"
	"github.com/mavr/totalwar/pkg/matrix"
)

type Simple2DWorld struct {
	uid    uuid.UUID
	matrix matrix.Matrix
	space  matrix.WorldSpace
}

func NewWorld(m matrix.Matrix) *Simple2DWorld {
	uid, _ := uuid.NewUUID()
	return &Simple2DWorld{
		uid:    uid,
		matrix: m,
		space:  m.GetStartSpace(),
	}
}

func (w *Simple2DWorld) UID() string {
	return w.uid.String()
}

func (w *Simple2DWorld) Name() string {
	return w.uid.String()
}

func (w *Simple2DWorld) Move(obj matrix.WorldObject) {
	w.matrix.Move(w.UID(), obj.UID())
}

func (w *Simple2DWorld) GetEnemy() []matrix.Enemy {
	return w.matrix.GetEnemies(w.UID())
}

func (w *Simple2DWorld) IsAccessibility(aimUID string) bool {
	return w.matrix.IsAccessibility(w.UID(), aimUID)
}

func (w *Simple2DWorld) GetCoordinates() matrix.WorldSpace {
	return w.space
}

func (w *Simple2DWorld) Attack(aim matrix.WorldObject) bool {
	return w.matrix.Attack(w.UID(), aim.UID())
}
