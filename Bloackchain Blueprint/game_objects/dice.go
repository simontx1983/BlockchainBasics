package dice

import (
	"math/rand"
	"time"
)

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
