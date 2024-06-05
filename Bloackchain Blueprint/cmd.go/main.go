package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

const (
	maxDiceLow       int = 1
	maxDiceHigh      int = 6
	placesOnBoard    int = 48
	minThresholdData int = 500
	formattingHashes int = 28
)

type GameState struct {
	CurrentPlayer         *player.player
	CurrentPropertyOfTurn *property.property
	CurrentDiceRoll       int
	GlobalTurnsMade       int
	AllPlayers            []Player
	AllProperties         *Property
	
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Roll simulates rolling two six-sided dice and returns their sum.
func Roll() int {
	// Roll two dice, each with a value between 1 and 6.
	dice1 := rand.Intn(6) + 1
	dice2 := rand.Intn(6) + 1

	return dice1 + dice2
}

func (gs *GameState) Roll() {
	gs.GlobalTurnsMade++
	gs.CurrentDiceRoll = roll()
	fmt.Println("You Rolled an, "gs.CurrentDiceRoll)
}

func GetTheCurrentCard(board int, gs *GameState) (string, *property) {
	for _, card := range gs.AllProperties.AllProperty {
		for n, v := range card.Card {
			if v.PositionOnBoard == board {
				return n, v
			}
		}
	}
	return GetTheOtherPropertyName(board, gs.Others), nil
}

func GetTheOtherPropertyName(board int, oc *OtherPropertyCollection) string {
	for _, card := range oc.AllProperty {
		for n, v := range card.Card {
			if v.PositionOnBoard == board {
				return n
			}
		}
	}
	return "(unknown)"
}

func GetTheCurrentCardName(board int, gs *GameState) string {
	name, _ := GetTheCurrentCard(board, gs)
	return name
}

func (gs *GameState) ProcessNonPropertySquare(CurrentPlayer *Player, sqType int, DAO int, cc *CardCollection) {
	taxCollection := 0
	switch sqType {
	case Tax:
		fmt.Println("Let's Join A DAO!")
		taxCollection += DAO
		t := Transaction{
			Sender:   CurrentPlayer,
			Receiver: nil,
		}
		if CurrentPlayer.PositionOnBoard == 7 {
			t.Amount += DAO.DaoCards
		}
		t.TransactWithBank()
		fmt.Println("Collected Tax: $", t.Amount)
	case Jail:
		if CurrentPlayer.PositionOnBoard == 30 {
			CurrentPlayer.PositionOnBoard = 10
			CurrentPlayer.JailTurns = 3
		}
	case JustVisiting, FreeParking:
		fmt.Println("Have a rest!")
	case Payment:
		fmt.Println("Block Mined!")
	case Chance:
		gs.processDrawCard(0, cc)
		gs.processDrawCard(16, cc)
	default:
		fmt.Println("Unknown or To Be Implemented")
	}
}

func (gs *GameState) processDrawCard(offset int, cc *CardCollection) {
	var card DrawCard
	if offset == 0 {
		cc.CurrentCardH = (cc.CurrentCardH + 1) % len(cc.ShuffleOrderH)
		card = cc.AllDrawCards[cc.ShuffleOrderH[cc.CurrentCardH]]
	} else {
		cc.CurrentCardO = (cc.CurrentCardO + 1) % len(cc.ShuffleOrderO)
		card = cc.AllDrawCards[offset+cc.ShuffleOrderO[cc.CurrentCardO]]
	}
	fmt.Println(drawCardType(card.Designator), "Card", card.Id, "->", card.Content)
	processDrawCardInternal(&card, gs, cc)
}

func (gs *GameState) GoToSquare(space int, paymentCheck bool) {
	if space < 0 || space > 39 {
		panic("We are attempting to move to a board space out of range: " + strconv.Itoa(space))
	}
	prePosition := gs.CurrentPlayer.PositionOnBoard
	gs.CurrentPlayer.PositionOnBoard = space
	fmt.Println("Player has moved to space", space, GetTheCurrentCardName(space, gs))
	if gs.CurrentPlayer.PositionOnBoard < prePosition && paymentCheck == true {
		gs.CurrentPlayer.pay200Dollars()
	}
}

func (gs *GameState) RemoveToken(playerToRemove *Player) {
	fmt.Println("Removing token", playerToRemove.Token, "played by", playerToRemove.Name)
	for _, j := range gs.AllPlayers {
		if j.PlayerNumber == playerToRemove.PlayerNumber {
			j.Active = false
			break
		}
	}
}

func main() {
	board, players := initializeGame()

	for {
		renderBoard(board)
		renderPlayerInfo(players)

		currentPlayer := getCurrentPlayer(players)

		diceRoll := rollDice()
		fmt.Printf("%s rolled %d\n", currentPlayer.Name, diceRoll)

		movePlayer(currentPlayer, diceRoll, board)

		handlePlayerTurn(currentPlayer, board)

		if isGameOver(players) {
			break
		}
		switchToNextPlayer(players)
	}
	fmt.Println("Game over!")
	displayLeaderboard(players)
}

func initializeGame() (*Board, []Player) {
	return nil, nil
}

func renderBoard(board *Board) {
}

func renderPlayerInfo(players []Player) {
}

func getCurrentPlayer(players []Player) *Player {
	return nil
}

func movePlayer(player *Player, diceRoll int, board *Board) {
}

func handlePlayerTurn(player *Player, board *Board) {
}

func switchToNextPlayer(players []Player) {
}

func displayLeaderboard(players []Player) {
}
