package group

type Group struct {
	id int
    name string
}

type Grouper interface {
    Id() int
    Name() string
}

func (g Group) Id() int {
    return g.id
}

func (g Group) Name() string {
    return g.name
}