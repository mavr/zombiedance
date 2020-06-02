package cmdgp

import (
	"fmt"

	"github.com/mavr/totalwar/pkg/gameplay"
)

type CommandGamePlay struct{}

func New() *CommandGamePlay {
	return &CommandGamePlay{}
}

func (gp *CommandGamePlay) Check(players []gameplay.Player) bool {
	// fmt.Println("check by gameplay")
	m := make(map[gameplay.PlayerType]int)
	for _, p := range players {
		m[p.Clan()]++
	}

	if len(m) <= 1 {
		return true
	}

	fmt.Printf("Opposition : %d\tzombie : %d\n", m[gameplay.Opposition], m[gameplay.Zombie])

	return false
}

func (gp *CommandGamePlay) Result(all, alive []gameplay.Player) {
	fmt.Println("====================================")
	if len(alive) == 0 {
		fmt.Println("all deads")
		return
	}

	m := make(map[gameplay.PlayerType]int)

	for _, p := range all {
		m[p.Clan()]++
	}

	fmt.Printf("winner : %s; alive - %d; dead - %d\n", alive[0].Clan(), len(alive), m[alive[0].Clan()]-len(alive))
	for t := range m {
		if t == alive[0].Clan() {
			continue
		}

		fmt.Printf("%s dead %d\n", t, m[t])
	}
}
