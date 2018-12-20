package person

import (
    "group"
)

type Person struct {
    Id int
    Username string
    FirstName string
    LastName string
    Groups map[int]Grouper
}

type Personer interface {
    IsPartOfGroup(g Grouper) bool
}