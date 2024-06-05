package drawcards

import (
	"fmt"
	"math/rand"
	"time"
)

type DrawCard struct {
	Id           int
	Designator   byte // C = Community Node D = DAO
	Content      string
	PlayerToBank *Transaction
	BankToPlayer *Transaction
	MoveToSpace  *int
	RelativeMove *int
	NearestType  *int
}

type CardCollection struct {
	AllDrawCards  [32]DrawCard
	ShuffleOrderC []int
	ShuffleOrderD []int
	CurrentCardC  []int
	CurrentCardD  []int
}

func GenerateOrderForCommunityNodeCards() []int {
	cardsToDeal := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(cardsToDeal)
	rand.Seed(time.Now().Unix())
	time.Sleep(200 * time.Millisecond) // Ensuring it sleeps for 200 milliseconds
	rand.Shuffle(len(cardsToDeal), func(i, j int) { cardsToDeal[i], cardsToDeal[j] = cardsToDeal[j], cardsToDeal[i] })
	fmt.Println(cardsToDeal)
	return cardsToDeal
}

func drawCardType(cardType byte) string {
	var ctype string
	switch cardType {
	case 'C':
		ctype = "CommunityNode"
	case 'D':
		ctype = "DAO"
	default:
		ctype = "Unknown"
	}
	return ctype
}
