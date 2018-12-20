package person

import (
	"github.com/fernanlukban/cronus/common"
	"github.com/fernanlukban/cronus/group"
)

type Person struct {
	common.Ider
	Username  string
	FirstName string
	LastName  string
	Groups    map[int]Grouper
}

type Personer interface {
	IsPartOfGroup(g Grouper) bool
}
