package services

import (
	"fmt"
)

type Transaction struct {
	Sender   *game.Player
	Receiver *game.Player
	Amount   int
}

const (
	half       = 0.5
	tenPercent = 0.1
)

var BankGameState *GameState

type Bank struct {
	CashReservesInSat int
	TransactionLedger []Transaction
	TotalMiningRigs   int
	TotalDataCenter   int
}

// Player Gives money to the Bank
func (txn *Transaction) TransactionWithBank() {
	if txn.Sender.Active {
		if txn.Sender.CashAvailable < txn.Amount {
			fmt.Println("Not Enough SATs. Do you want a loan?")
			txn.Amount = txn.Sender.CashAvailable
			txn.Sender.Active = false
			BankGameState.RemoveToken(txn.Sender)
		}
		txn.Sender.CashAvailable -= txn.Amount
		TheBank.CashReservesInSat += txn.Amount
		TheBank.TransactionLedger = append(TheBank.TransactionLedger, *txn)
	}
}

// Bank Gives Players Sats
func (txn *Transaction) BankCheque() {
	TheBank.CashReservesInSat -= txn.Amount
	txn.Receiver.SatsAvailable += txn.Amount
	if txn.Sender != nil {
		panic("Can not set the receiver when doing a bank cheque!")
	}
	if TheBank.CashReservesInSat <= 0 {
		panic("The bank has gone bankrupt! Game Is over")
	}
	TheBank.TransactionLedger = append(TheBank.TransactionLedger, *txn)
}

func CashQuestions(txn Transaction) bool {
	if txn.Sender.CashAvailable >= txn.Amount {
		fmt.Println("Debt can be paid off by answering a few quick questions. You need", txn.Amount, "but you only have", txn.Sender.CashAvailable)
		return true
	}
	return false
}

func (txn *Transaction) mortgage() bool {
	_, props := ShowPropertiesOfPlayer(txn.Sender.PlayerNumber, BankGameState)
	for _, prop := range props {
		amount := int(half * float64(prop.PurchaseCost))
		t := Transaction{
			Sender:   nil,
			Receiver: txn.Sender,
			Amount:   amount,
		}
		t.BankCheque()
		prop.Mortgaged = true
		fmt.Println("Mortgaged", GetTheCurrentCardName(prop.PositionOnBoard, BankGameState), "for", amount)

		if txn.Sender.CashAvailable >= txn.Amount {
			fmt.Println("Debt can be paid off after mortgaging. Needed", txn.Amount, "have", txn.Sender.CashAvailable)
			txn.Amount = txn.Amount
			t.TransactWithPlayer('x')
			return true
		}
	}
	return false
}
