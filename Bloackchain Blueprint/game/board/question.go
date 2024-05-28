package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Question struct {
	Text         string   // The text of the question
	Options      []string // The multiple-choice options
	CorrectIndex int      // The index of the correct answer in the Options slice
	Reward       int      // The discount or reward for answering correctly
}

func main() {
	questions := []Question{
		{
			Text:         "What is the maximum supply of Bitcoin?",
			Options:      []string{"21 million", "42 million", "18 million", "50 million"},
			CorrectIndex: 0,
			Reward:       50,
		},
		{
			Text:         "Which blockchain is known for its smart contract functionality?",
			Options:      []string{"Bitcoin", "Ethereum", "Litecoin", "Ripple"},
			CorrectIndex: 1,
			Reward:       50,
		},
		{
			Text:         "What is the smallest unit of Bitcoin called?"
			Options:      []string,{"Satoshi", "Ether", "Finney", "Ada"}
			CorrectIndex: 0,
			Reward:       50,
		},
		{
			Text:         "What is the process of validating transactions and adding them to the blockchain called?",
			Options:      []string{"Mining", "Trading", "Hodling", "Staking"},
			CorrectIndex: 0,
			Reward:       50,
		},
		{
			Text:         "What is the first cryptocurrency ever created?",
			Options:      []string{"Bitcoin", "Ethereum", "Ripple", "Litecoin"},
			CorrectIndex: 0,
			Reward:       50,
		},
		{
			Text:         "What is the name of the consensus algorithm used by Bitcoin?",
			Options:      []string{"Proof of Work", "Proof of Stake", "Delegated Proof of Stake", "Proof of Authority"},
			CorrectIndex: 0,
			Reward:       50,
		},
		{
			Text:         "What is the name of the creator of Bitcoin?",
			Options:      []string{"Satoshi Nakamoto", "Vitalik Buterin", "Charlie Lee", "Brad Garlinghouse"},
			CorrectIndex: 0,
			Reward:       50,
		},
		{
			Text:         "What is the name of the programming language used to write smart contracts on Ethereum?",
			Options:      []string{"Solidity", "Java", "Python", "C++"},
			CorrectIndex: 0,
			Reward:       50,
		},
		{
			Text: "What is the name of the first cryptocurrency exchange?",
			Options: []string{"Mt. Gox", "Binance", "Coinbase", "Kraken"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{

			Text: "What is the name of the first cryptocurrency?",
			Options: []string{"Bitcoin", "Ethereum", "Ripple", "Litecoin"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is a Decentralized application (DApp)?",
			Options: []string{"An application that runs on a decentralized network", "An application that runs on a centralized network", "An application that runs on a private network", "An application that runs on a public network"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What does the term "HODL" mean in the context of cryptocurrencies?",
			Options: []string{"Hold on for dear life", "Buy low, sell high", "Diversify your portfolio", "Invest in ICOs"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the name of the process by which new bitcoins are created and added to the circulating supply?",
			Options: []string{"Mining", "Trading", "Hodling", "Staking"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is a fork in the context of cryptocurrencies?",
			Options: []string{"A change in the protocol of a blockchain that results in two separate versions of the blockchain", "A change in the price of a cryptocurrency", "A change in the consensus algorithm of a blockchain", "A change in the supply of a cryptocurrency"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the name of the process by which new bitcoins are created and added to the circulating supply?",
			Options: []string{"Mining", "Trading", "Hodling", "Staking"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the name of the first ever blocked mined on a blockchain?",
			Options: []string{"Genesis block", "Block 0", "Block 1", "Block A"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the purpose of a whitepaper in the context of cryptocurrencies?",
			Options: []string{"To explain the technology and vision behind a cryptocurrency project", "To provide a summary of the price of a cryptocurrency", "To outline the marketing strategy of a cryptocurrency project", "To list the team members of a cryptocurrency project"},
			CorrectIndex: 0,
			Reward: 50,
		},

		{

			Text: "What is the term used to describe the reduction in the reward for mining a block on a blockchain?",
			Options: []string{"Halving", "Doubling", "Tripling", "Quadrupling"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the name of the process by which new bitcoins are created and added to the circulating supply?",
			Options: []string{"Mining", "Trading", "Hodling", "Staking"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the purpose of the Lightning Network in the context of Bitcoin?",
			Options: []string{"To enable faster and cheaper transactions on the Bitcoin network", "To increase the supply of Bitcoin", "To improve the security of the Bitcoin network", "To create a new cryptocurrency"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the main advantage of a Decentralized Autonomous Organization (DAO)?",
			Options: []string{"It is governed by smart contracts and operates without human intervention", "It is controlled by a central authority and operates like a traditional organization", "It is regulated by a government agency and operates within the law", "It is managed by a team of developers and operates like a software company"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the main purpose of a token in the context of cryptocurrencies?",
			Options: []string{"To represent an asset or utility on a blockchain", "To provide a summary of the price of a cryptocurrency", "To outline the marketing strategy of a cryptocurrency project", "To list the team members of a cryptocurrency project"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is main fuction of a blockchain Node?",
			Options: []string{"To validate transactions and add them to the blockchain", "To create new bitcoins and add them to the circulating supply", "To regulate the price of a cryptocurrency", "To manage the marketing strategy of a cryptocurrency project"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is a node in the context of cryptocurrencies?",
			Options: []string{"A computer that participates in the network and helps validate transactions", "A type of cryptocurrency", "A type of smart contract", "A type of blockchain"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "gas" used for in the context of Ethereum?",
			Options: []string{"The fee required to execute a transaction or run a smart contract on the Ethereum network", "The price of a cryptocurrency", "The reward for mining a block on the Ethereum network", "The supply of Ethereum"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "staking" used for in the context of cryptocurrencies?",
			Options: []string{"The process of participating in the proof-of-stake consensus algorithm and earning rewards", "The process of participating in the proof-of-work consensus algorithm and earning rewards", "The process of participating in the delegated proof-of-stake consensus algorithm and earning rewards", "The process of participating in the proof-of-authority consensus algorithm and earning rewards"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text:"What is the term "FOMO" used for in the context of cryptocurrencies?",
			Options: []string{"Fear of missing out", "Fear of making an investment", "Fear of losing money", "Fear of selling too soon"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "FUD" used for in the context of cryptocurrencies?",
			Options: []string{"Fear, uncertainty, and doubt", "Fear, uncertainty, and despair", "Fear, uncertainty, and danger", "Fear, uncertainty, and denial"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "DYOR" used for in the context of cryptocurrencies?",
			Options: []string{"Do your own research", "Do your own reading", "Do your own reporting", "Do your own review"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "ATH" used for in the context of cryptocurrencies?",
			Options: []string{"All-time high", "All-time low", "All-time high price", "All-time high trading volume"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "ICO" used for in the context of cryptocurrencies?",
			Options: []string{"Initial coin offering", "Initial currency offering", "Initial crypto offering", "Initial cash offering"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Fiat" used for in the context of cryptocurrencies?",
			Options: []string{"Government-issued currency", "Digital currency", "Cryptocurrency", "Commodity currency"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Altcoin" used for in the context of cryptocurrencies?",
			Options: []string:{"Any cryptocurrency other than Bitcoin", "A type of smart contract", "A type of blockchain", "A type of token"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Whale" used for in the context of cryptocurrencies?",
			Options: []string{"An individual or entity that holds a large amount of cryptocurrency", "A type of smart contract", "A type of blockchain", "A type of token"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Bear market" used for in the context of cryptocurrencies?",
			Options: []string{"A market with declining prices and pessimism", "A market with rising prices and optimism", "A market with stable prices and uncertainty", "A market with volatile prices and speculation"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Bull market" used for in the context of cryptocurrencies?",
			Options: []string{"A market with rising prices and optimism", "A market with declining prices and pessimism", "A market with stable prices and uncertainty", "A market with volatile prices and speculation"},
			CorrectIndex: 0,
			Reward: 50,
		},

		{
			Text: "What is the term "Market cap" used for in the context of cryptocurrencies?",
			Options: []string{"The total value of a cryptocurrency in circulation", "The price of a cryptocurrency", "The reward for mining a block on a blockchain", "The supply of a cryptocurrency"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Wallet" used for in the context of cryptocurrencies?",
			Options: []string{"A digital storage for cryptocurrencies", "A physical storage for cryptocurrencies", "A type of smart contract", "A type of blockchain"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Private key" used for in the context of cryptocurrencies?",
			Options: []string{"A secret code that allows the owner to access their cryptocurrency", "A public code that allows the owner to access their cryptocurrency", "A type of smart contract", "A type of blockchain"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Public key" used for in the context of cryptocurrencies?",
			Options: []string{"A public code that allows the owner to receive cryptocurrency", "A secret code that allows the owner to access their cryptocurrency", "A type of smart contract", "A type of blockchain"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Exchange" used for in the context of cryptocurrencies?",
			Options: []string:{"A platform for buying and selling cryptocurrencies", "A type of smart contract", "A type of blockchain", "A type of token"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text:"What does the acronym "KYC" stand for in the context of cryptocurrencies?",
			Options: []string{"Know Your Customer", "Know Your Cryptocurrency", "Know Your Coin", "Know Your Code"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What does the acronym "AML" stand for in the context of cryptocurrencies?",
			Options: []string{"Anti-Money Laundering", "Anti-Market Liquidity", "Anti-Market Loss", "Anti-Market Leverage"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Hash" used for in the context of cryptocurrencies?",
			Options: []string{"A unique code that represents a block of transactions", "A type of smart contract", "A type of blockchain", "A type of token"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Block" used for in the context of cryptocurrencies?",
			Options: []string:{"A group of transactions that are added to the blockchain", "A type of smart contract", "A type of blockchain", "A type of token"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Blockchain" used for in the context of cryptocurrencies?",
			Options: []string:{"A decentralized and distributed digital ledger", "A type of smart contract", "A type of token", "A type of cryptocurrency"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text:"What does the acronym "DLT" stand for in the context of cryptocurrencies?",
			Options: []string{"Distributed Ledger Technology", "Digital Ledger Token", "Decentralized Ledger Technology", "Digital Ledger Transaction"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Decentralization" used for in the context of cryptocurrencies?",
			Options: []string:{"The distribution of control and ownership of a network", "A type of smart contract", "A type of token", "A type of blockchain"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Consensus" used for in the context of cryptocurrencies?",
			Options: []string:{"The process by which decisions are made on a blockchain network", "A type of smart contract", "A type of token", "A type of blockchain"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			Text: "What is the term "Smart contract" used for in the context of cryptocurrencies?",
			Options: []string:{"Self-executing contracts with the terms of the agreement between buyer and seller being directly written into lines of code", "A type of token", "A type of blockchain", "A type of cryptocurrency"},
			CorrectIndex: 0,
			Reward: 50,
		},
		{
			


	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Select a random question
	selectedQuestion := questions[rand.Intn(len(questions))]

	// Display the question and options
	fmt.Println("Question:", selectedQuestion.Text)
	for i, option := range selectedQuestion.Options {
		fmt.Printf("%d: %s\n", i+1, option)
	}
	fmt.Printf("Reward for correct answer: %d%% discount\n", selectedQuestion.Reward)
}

