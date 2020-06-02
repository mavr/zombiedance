package gameplay

type PlayerType string
type GamePlayType string

var (
	Zombie     PlayerType = "zombie"
	Opposition PlayerType = "opposition"
)

var (
	GamePlayZombiVsOpposition GamePlayType = "zomopmatch"
)

type Player interface {
	UID() string
	Clan() PlayerType
}

type Gameplay interface {
	Check([]Player) bool
	Result(all, alive []Player)
}
