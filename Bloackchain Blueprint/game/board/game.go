package game

type Board struct {
	Properties []Property
	Cards      []Card
}

func NewBoard() *Board {
	// Initialize properties and cards
	properties := LoadProperties()
	cards := LoadCards()
	return &Board{Properties: properties, Cards: cards}
}

func LoadProperties() []Property {
	// Load properties from assets
	// Example: return []Property{...}
	return []Property{}
}

func LoadCards() []Card {
	// Load cards from assets
	// Example: return []Card{...}
	return []Card{}
}
