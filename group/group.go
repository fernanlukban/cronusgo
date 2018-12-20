package group

import (
    "github.com/fernanlukban/cronus/common"
)

type Group struct {
	id int
    name string
}

type Grouper interface {
    common.IderAndNamer
}

func (g Group) Id() int {
    return g.id
}

func (g Group) Name() string {
    return g.name
}