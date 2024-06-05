package bank

import (
	"fmt"
)

type Bank struct {
	CashReserves int
}

func NewBank(initialCashReserves int) *Bank {
	return &Bank{
		CashReserves: initialCashReserves,
	}
}

func (b *Bank) Deposit(player *Player, amount int) error {
	// Implement deposit logic
	return nil
}

func (b *Bank) Withdraw(player *Player, amount int) error {
	// Implement withdrawal logic
	return nil
}

func (b *Bank) BuyProperty(player *Player, property *Property) error {
	// Implement property buying logic
	return nil
}

// SaveState saves the bank's state to a file or database.
func (b *Bank) SaveState() error {
	// Implement state saving logic
	return nil
}

// LoadState loads the bank's state from a file or database.
func (b *Bank) LoadState() error {
	// Implement state loading logic
	return nil
}

// BuyPropertyWithDiscount handles the transaction when a player buys a property from the bank with a discount after answering a question correctly.
func (b *Bank) BuyPropertyWithDiscount(player *Player, property *Property, question *Question) error {
	// Display the question and options
	fmt.Println("Question:", question.Text)
	for i, option := range question.Options {
		fmt.Printf("%d: %s\n", i+1, option)
	}
	fmt.Printf("Reward for correct answer: %d%% discount\n", question.Reward)

	// Get the user's answer
	var userAnswer int
	fmt.Print("Enter your answer: ")
	_, err := fmt.Scan(&userAnswer)
	if err != nil {
		return err
	}

	// Check if the answer is correct
	if userAnswer-1 == question.CorrectIndex {
		// Apply the discount
		discountedPrice := float64(property.Price) * (1 - float64(question.Reward)/100)
		if player.Balance >= int(discountedPrice) {
			player.Balance -= int(discountedPrice)
			player.Properties = append(player.Properties, property)
			fmt.Printf("Congratulations! You answered correctly and bought %s for %d.\n", property.Name, int(discountedPrice))
		} else {
			fmt.Println("You don't have enough balance to buy this property, even with the discount.")
		}
	} else {
		fmt.Println("Sorry, your answer is incorrect. You cannot buy the property with a discount.")
	}

	return nil
}
