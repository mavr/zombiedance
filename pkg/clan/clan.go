package clan

import "github.com/google/uuid"

const (
	ClanOpposition ClanType = "opposition"
	ClanZombie     ClanType = "zombie"
)

var ZombieClan *CommonClan

func init() {
	ZombieClan = New("zombie band", ClanZombie)
}

type ClanType string

type Clan interface {
	UID() uuid.UUID
	Type() ClanType
	Name() string
}

type CommonClan struct {
	uid  uuid.UUID
	t    ClanType
	name string
}

func New(name string, t ClanType) *CommonClan {
	uid, _ := uuid.NewUUID()
	return &CommonClan{
		uid:  uid,
		name: name,
		t:    t,
	}
}

func (c *CommonClan) UID() uuid.UUID {
	return c.uid
}

func (c *CommonClan) Name() string {
	return c.name
}

func (c *CommonClan) Type() ClanType {
	return c.t
}
