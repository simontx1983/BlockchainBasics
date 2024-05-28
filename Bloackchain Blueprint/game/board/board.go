package game

import (
	"fmt"
	"math/rand"

	"github.com/simontx1983/Blockchain_Blueprint/game"
)

const (
	boardSize        = 48
	maxPropertyValue = 1500
)

// Property represents a property on the game board
type Property struct {
	Name        string
	Value       int
	Owner       *game.Player
	QuestionIdx int
}

// Board represents the game board
type Board struct {
	Properties []Property
}

// NewBoard creates a new game board
func NewBoard() *Board {
	board := &Board{
		Properties: make([]Property, boardSize),
	}

	// Initialize board properties
	for i := 0; i < boardSize; i++ {
		board.Properties[i] = Property{
			Name:        fmt.Sprintf("Property %d", i+1),
			Value:       rand.Intn(maxPropertyValue) + 1,
			QuestionIdx: rand.Intn(maxQuestions),
		}
	}

	return board
}

// GetProperty returns the property at the given position
func (b *Board) GetProperty(position int) *Property {
	return &b.Properties[position]
}

// SetPropertyOwner sets the owner of a property
func (b *Board) SetPropertyOwner(position int, owner *game.Player) {
	b.Properties[position].Owner = owner
}
