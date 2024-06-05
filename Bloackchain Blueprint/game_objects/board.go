package board

import "fmt"

// Board represents the game board
type Board struct {
	Tiles []*Tile
}

type Tile struct {
	Name        string
	Type        TileType
	Property    *Property
	Node        *Node
	EventCard   *EventCard
	DAOCard     *DAOCard
	SpecialTile *SpecialTile
}

// TileType represents the type of a tile
type TileType int

const (
	PropertyTile TileType = iota
	NodeTile
	EventTile
	DAOTile
	SpecialTile
)

// Property represents a property that can be owned by a player
type Property struct {
	Name        string
	Price       int
	Owner       *Player
	Nodes       []*Node
	Multiplier  float64
	QuestionSet *QuestionSet
}

// Node represents a node that can be placed on a property
type Node struct {
	Name        string
	Owner       *Player
	Multiplier  float64
	QuestionSet *QuestionSet
}

// EventCard represents an event card that can be drawn
type EventCard struct {
	Name        string
	Description string
	Effect      func(*Player, *Board)
}

// DAOCard represents a DAO card that can be drawn
type DAOCard struct {
	Name        string
	Description string
	Effect      func(*Player, *Board)
}

// SpecialTile represents a special tile on the board
type SpecialTile struct {
	Name        string
	Description string
	Effect      func(*Player, *Board)
}

// NewBoard creates a new game board
func NewBoard() *Board {
	// Initialize the board with tiles
	tiles := make([]*Tile, 0)
	// Add tiles to the board
	return &Board{Tiles: tiles}
}

// RenderBoard prints the current state of the board
func (b *Board) RenderBoard() {
	fmt.Println("Current Board State:")
	for _, tile := range b.Tiles {
		fmt.Printf("%s: %v\n", tile.Name, tile)
	}
}
