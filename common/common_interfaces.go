package common

type Ider interface {
    Id() int
}

type Namer interface {
    Name() string
}

type IderAndNamer interface {
    Ider
    Namer
}