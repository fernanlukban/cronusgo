package person

import (
    "github.com/fernanlukban/cronus/group"
    "github.com/fernanlukban/cronus/common"
)

type Person struct {
    common.Ider
    Username string
    FirstName string
    LastName string
    Groups map[int]Grouper
}

type Personer interface {
    IsPartOfGroup(g Grouper) bool
}