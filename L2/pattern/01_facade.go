package pattern

import (
	"fmt"
	"strings"
)

// NewMan создает man
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Facade
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Todo возвращает, что мужщина должен сделать
func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

type House struct {
}

// Build реализация.
func (h *House) Build() string {
	return "Build house"
}

type Tree struct {
}

// Grow реализация.
func (t *Tree) Grow() string {
	return "Tree grow"
}

type Child struct {
}

// Born реализация.
func (c *Child) Born() string {
	return "Child born"
}

func Facade() {
	man := NewMan()
	res := man.Todo()
	fmt.Println(res)
}
