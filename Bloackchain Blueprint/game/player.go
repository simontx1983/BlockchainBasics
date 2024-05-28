package game

import (
	"errors"
	"fmt"
)

const (
	roundTripPayment = 200
)

// Player represents a player in the game.
type Player struct {
	PlayerNumber    int
	Name            string
	CashAvailable   int
	PositionOnBoard int
	Active          bool
	Turns           int
	Token           string
	SlashedTurns    int
	SlashedCards    []byte
	MiningNodes     []MiningNode
	DataCenters     []DataCenter
}

// AdvancePlayer advances the player on the board based on the number of steps.
func (p *Player) AdvancePlayer(steps int, cc *CardCollection) int {
	p.Turns++
	prePosition := p.PositionOnBoard
	if p.SlashedTurns == 0 {
		p.PositionOnBoard += steps
	} else if p.SlashedTurns > 0 && len(p.SlashedCards) > 0 {
		p.SlashedTurns = 0
		backIntoStack := p.SlashedCards[0]
		if backIntoStack == 'D' {
			cc.ShuffleOrderH = append(cc.ShuffleOrderC, 15)
		} else {
			cc.ShuffleOrderO = append(cc.ShuffleOrderD, 11)
		}
		p.SlashedCards = p.SlashedCards[1:]
		p.PositionOnBoard += steps
	} else {
		firstRoll, secondRoll := rollToGetOutOfJail()
		if firstRoll == secondRoll {
			fmt.Println("Rolled a double! Let's avoid getting slashed")
			p.SlashedTurns = 0
			p.PositionOnBoard += firstRoll + secondRoll
		} else {
			if p.SlashedTurns == 1 {
				fmt.Println("Exhausted all rolls, pay $1000 sats to get out and roll", firstRoll+secondRoll, "spaces")
				t := Transaction{
					Sender:   p,
					Receiver: nil,
					Amount:   100,
				}
				t.TransactWithBank()
				p.SlashedTurns = 0
				p.PositionOnBoard += firstRoll + secondRoll
			} else {
				p.SlashedTurns--
				fmt.Println("Rolled a ", firstRoll, "and", secondRoll, ". Not successful.", p.SlashedTurns, "more tries available")
			}
		}
	}
	placesOnBoard := 40 
	p.PositionOnBoard = p.PositionOnBoard % placesOnBoard
	if p.PositionOnBoard < prePosition && p.SlashedTurns == 0 {
		p.pay100sats()
		return roundTripPayment
	}
	return 0
}

// BuyProperty buys a property for the player.
func (p *Player) BuyProperty(pd *PropertyDeed) (int, error) {
	fmt.Println("5. Buy property if available")
	if p.CashAvailable-pd.PurchaseCost < 0 {
		return 0, errors.New("Cannot afford property!")
	}
	t := Transaction{
		Sender:   p,
		Receiver: nil,
		Amount:   pd.PurchaseCost,
	}
	t.TransactWithBank()
	pd.Owner = byte(p.PlayerNumber)
	return pd.PurchaseCost, nil
}

// pay200sats pays 200 satoshis to the player.
func (p *Player) pay20000sats() {
	t := Transaction{
		Sender:   nil,
		Receiver: p,
		Amount:   200,
	}
	t.BankCheque()
}

// BuyPropertyWithQuestion attempts to buy a property based on a question.
func BuyPropertyWithQuestion(player *Player, question Question) error {
	fmt.Println("Attempting to buy property:", question.Property.Name)
	if player.CashAvailable < question.Property.PurchaseCost/2 {
		return errors.New("Insufficient funds to purchase property")
	}

	// Ask the question
	fmt.Println(question.Text)
	var playerAnswer string
	fmt.Scanln(&playerAnswer)

	// Check the answer
	if playerAnswer == question.Answer {
		fmt.Println("Correct answer! Purchasing property at half price.")
		// Buy property at half price
		purchaseCost := question.Property.PurchaseCost / 2
		t := Transaction{
			Sender:   player,
			Receiver: nil,
			Amount:   purchaseCost,
		}
		t.TransactWithBank()
		question.Property.Owner = byte(player.PlayerNumber)
		return nil
	}
	fmt.Println("Incorrect answer! Purchasing property at full price.")
	// Buy property at full price
	t := Transaction{
		Sender:   player,
		Receiver: nil,
		Amount:   question.Property.PurchaseCost,
	}
	t.TransactWithBank()
	question.Property.Owner = byte(player.PlayerNumber)
	return nil
}

// CheckOwnedSquare checks if a player owns a particular property.
func CheckOwnedSquare(player *Player, property *PropertyDeed) bool {
	return property.Owner == byte(player.PlayerNumber)
}

// AddMiningNode adds a mining node or upgrades to a data center on a property owned by the player.
func AddMiningNode(player *Player, property *PropertyDeed) error {
	if !CheckOwnedSquare(player, property) {
		return errors.New("You don't own this property")
	}

	if len(player.MiningNodes) < 5 {
		// Add a mining node
		player.MiningNodes = append(player.MiningNodes, MiningNode{})
		fmt.Println("Added a mining node to", property.Name)
		return nil
	}

	// Upgrade to a data center
	player.DataCenters = append(player.DataCenters, DataCenter{})
	// Remove 5 mining nodes
	player.MiningNodes = player.MiningNodes[:len(player.MiningNodes)-5]
	fmt.Println("Upgraded to a data center on", property.Name)
	return nil

}
func Move(player *Player, steps int, board *Board) error {
	// Validate input
	if player == nil {
		return errors.New("player is nil")
	}
	if board == nil {
		return errors.New("board is nil")
	}
	if steps <= 0 {
		return errors.New("invalid number of steps")
	}

	// Move the player
	oldPosition := player.PositionOnBoard
	newPosition := (oldPosition + steps) % len(board.Properties)
	player.PositionOnBoard = newPosition

	// Check if the player completed a loop
	if newPosition < oldPosition {
		player.CashAvailable += roundTripPayment // Reward for completing a loop
		fmt.Printf("%s completed a loop and earned %d sats!\n", player.Name, roundTripPayment)
	}

	// Handle landing on a property
	property := board.Properties[newPosition]
	if property.Owner == nil {
		// Offer to buy or answer a question
		if player.CashAvailable < property.PurchaseCost {
			fmt.Printf("%s does not have enough sats to buy %s.\n", player.Name, property.Name)
		} else {
			fmt.Printf("%s landed on %s. Do you want to buy it for %d sats or answer a question for half price? (buy/answer)\n", player.Name, property.Name, property.PurchaseCost)
			var choice string
			fmt.Scanln(&choice)
			if choice == "buy" {
				player.BuyProperty(&property)
				fmt.Printf("%s bought %s for %d sats.\n", player.Name, property.Name, property.PurchaseCost)
			} else if choice == "answer" {
				question := property.Question
				if AnswerQuestion(question) {
					purchaseCost := property.PurchaseCost / 2
					player.BuyProperty(&property)
					fmt.Printf("%s answered correctly and bought %s for %d sats.\n", player.Name, property.Name, purchaseCost)
				} else {
					fmt.Println("Wrong answer! The property goes back to the bank.")
				}
			}
		}
	} else if property.Owner == player {
		// Add a node
		fmt.Printf("%s landed on their own property %s. Do you want to add a node? (yes/no)\n", player.Name, property.Name)
		var choice string
		fmt.Scanln(&choice)
		if choice == "yes" {
			nodePrice := property.NodePrice
			if player.CashAvailable >= nodePrice {
				player.AddMiningNode(&property)
				fmt.Printf("%s added a node to %s for %d sats.\n", player.Name, property.Name, nodePrice)
			} else {
				fmt.Println("Not enough sats to buy a node.")
			}
		}

	return nil
}
