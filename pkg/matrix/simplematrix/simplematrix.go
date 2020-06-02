package simplematrix

import (
	"math/rand"
	"sync"
	"time"

	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/matrix"
	"github.com/mavr/totalwar/pkg/unit"
)

type liveUnit struct {
	u   matrix.Enemy
	tch chan struct{}
}

type SimpleMatrix struct {
	unitsRunTime map[string]liveUnit
	mUnitRT      sync.RWMutex
	closeChan    chan struct{}

	units map[string]matrix.Enemy

	gamePlay gameplay.Gameplay
}

func NewMatrix(gameplay gameplay.Gameplay) *SimpleMatrix {
	return &SimpleMatrix{
		units:    make(map[string]matrix.Enemy),
		gamePlay: gameplay,
	}
}

func (m *SimpleMatrix) getRuntimeUnitByUID(uid string) matrix.Enemy {
	u, ok := m.units[uid]

	if ok {
		return u
	}

	return nil
}
func (m *SimpleMatrix) Move(unitUID, aimUID string) {
	var enemy matrix.Enemy
	var aim matrix.Enemy

	if enemy = m.getRuntimeUnitByUID(unitUID); enemy == nil {
		return
	}

	if aim = m.getRuntimeUnitByUID(aimUID); aim == nil {
		return
	}

	uSpace := enemy.GetCoordinates()
	aimSpace := aim.GetCoordinates()

	uSpace.Move(aimSpace, enemy.GetSpeed())
}

func (m *SimpleMatrix) Attack(unitUID, aimUID string) bool {
	var enemy matrix.Enemy
	var aim matrix.Enemy

	if enemy = m.getRuntimeUnitByUID(unitUID); enemy == nil || !enemy.Alive() {
		return false
	}

	if aim = m.getRuntimeUnitByUID(aimUID); aim == nil {
		return false
	}

	dead := aim.Protect(enemy.GetDamage())

	return dead
}

func (m *SimpleMatrix) removeRuntimeUnit() {
	t := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-t.C:
			m.mUnitRT.Lock()

			g := make([]gameplay.Player, 0, len(m.unitsRunTime))
			for i := range m.unitsRunTime {
				if m.unitsRunTime[i].u.Alive() {
					g = append(g, m.unitsRunTime[i].u)
				} else {
					delete(m.unitsRunTime, m.unitsRunTime[i].u.UID())
				}
			}

			m.mUnitRT.Unlock()
			if m.gamePlay.Check(g) {
				t.Stop()
				m.stop()

				return
			}
		case <-m.closeChan:
			t.Stop()
			return
		}
	}
}

func (m *SimpleMatrix) GetEnemies(unitUID string) []matrix.Enemy {
	m.mUnitRT.RLock()

	e := []matrix.Enemy{}
	for i := range m.unitsRunTime {
		if i != unitUID {
			e = append(e, m.units[i])
		}
	}

	m.mUnitRT.RUnlock()
	return e
}

func (m *SimpleMatrix) CreateUnit(clan gameplay.PlayerType) matrix.Enemy {
	t := unit.New(clan, NewWorld(m))
	m.units[t.UID()] = t

	return t
}

func (m *SimpleMatrix) GetStartSpace() matrix.WorldSpace {
	return NewSpace(
		matrix.WorldSpacePiece(rand.Float32()*100),
		matrix.WorldSpacePiece(rand.Float32()*100),
	)
}

func (m *SimpleMatrix) IsAccessibility(unitUID, aimUID string) bool {
	var enemy matrix.Enemy
	var aim matrix.Enemy

	if enemy = m.getRuntimeUnitByUID(unitUID); enemy == nil {
		return false
	}

	if aim = m.getRuntimeUnitByUID(aimUID); aim == nil {
		return false
	}

	dist := enemy.GetCoordinates().Distance(aim.GetCoordinates())

	if dist < enemy.GetAttackDistance() {
		return true
	}

	return false
}

func (m *SimpleMatrix) Run() {
	m.mUnitRT = sync.RWMutex{}

	m.closeChan = make(chan struct{})

	m.unitsRunTime = make(map[string]liveUnit)
	for i := range m.units {
		ch := make(chan struct{})
		m.unitsRunTime[i] = liveUnit{
			u:   m.units[i],
			tch: ch,
		}

		go m.unitsRunTime[i].u.Run(ch, nil)
	}

	go m.removeRuntimeUnit()

	t := time.NewTicker(100 * time.Millisecond)

	chArr := make([]chan struct{}, len(m.units))
	for {
		select {
		case <-t.C:
			i := 0
			m.mUnitRT.RLock()
			for u := range m.unitsRunTime {
				if m.unitsRunTime[u].u.Alive() {
					chArr[i] = m.unitsRunTime[u].tch
					i++
				}
			}
			m.mUnitRT.RUnlock()

			for _, ch := range chArr {
				select {
				case ch <- struct{}{}:
				default:
				}
			}

		case <-m.closeChan:
			t.Stop()

			for _, u := range m.unitsRunTime {
				close(u.tch)
			}

			coll := make([]gameplay.Player, 0, len(m.units))
			for i := range m.units {
				coll = append(coll, m.units[i])
			}
			g := make([]gameplay.Player, 0, len(m.unitsRunTime))
			for i := range m.unitsRunTime {
				if m.unitsRunTime[i].u.Alive() {
					g = append(g, m.unitsRunTime[i].u)
				}
			}

			m.gamePlay.Result(coll, g)

			return
		}
	}
}

func (m *SimpleMatrix) stop() {
	go func() {
		m.closeChan <- struct{}{}
	}()
}
