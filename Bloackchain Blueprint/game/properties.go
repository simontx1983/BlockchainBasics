package main

import (
	"fmt"
)

type Property struct {
	Name         string // Name of Property
	Position     int    // Position on Board
	PurchaseCost int    // cost to buy property
	NodeCost     int    // cost to buy a node
	Rent         []int  // Rent cost
	Owner        *Player
	NodeCount    int // number of nodes owned
}

type Player struct {
	Name   string // Name of Player
	Balance int   // Balance of player
}

func main() {
	bitcoinProperty := Property{
		Name:         "Bitcoin",
		Position:     0,
		PurchaseCost: 500,
		NodeCost:     250,
		Rent:         []int{},
		Owner:        nil,
		NodeCount:    0,
	}

	ethereum := Property{
		Name:         "Ethereum",
		Position:     1,
		PurchaseCost: 400,
		NodeCost:     250,
		Rent:         []int{75, 225, 300, 500, 800},
		Owner:        nil,
		NodeCount:    0,
	}

	solana := Property{
		Name:         "Solana",
		Position:     2,		
		PurchaseCost: 300,
		NodeCost:     200,
		Rent:         []int{50, 150, 200, 350, 600},
		Owner:        nil,
		NodeCount:    0,
	}
	Binace := Property{
		Name:         "Binance",
		Position:     3,
		PurchaseCost: 200,
		NodeCost:     150,
		Rent:         []int{25, 75, 100, 150, 300},
		Owner:        nil,
		NodeCount:    0,
	}
	Ton := Property{	
		Name:         "Ton",
		Position:     4,
		PurchaseCost: 200,
		NodeCost:     150,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	XRP := Property{
		Name:         "XRP",
		Position:     5,
		PurchaseCost: 200,
		NodeCost:     150,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Avalanche := Property{
		Name:         "Avalanche",
		Position:     6,
		PurchaseCost: 200,
		NodeCost:     150,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Lido DAO := Property{
		Name:         "Lido DAO",
		Position:     7,
		Owner:        nil,
		NodeCount:    0,
	}
	Cardano := Property{
		Name:         "Cardano",
		Position:     8,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Bitcoin Cash := Property{
		Name:         "Bitcoin Cash",
		Position:     9,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Tezos := Property{
		Name:         "Tezos",
		Position:     10,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Hedera Hashgraph := Property{
		Name:         "Hedera Hashgraph",
		Position:     11,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Cosmos Hub := Property{
		Name:         "Cosmos",
		Position:     12,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Chainlink := Property{
		Name:         "Chainlink",
		Position:     13,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}

	Injective := Property{
		Name:         "Injective",
		Position:     15,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Marker Dao := Property{
		Name:         "Marker Dao",
		Position:     14,
		Owner:        nil,
		NodeCount:    0,
	}

	Uniswap := Property{
		Name:         "Uniswap",
		Position:     16,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Arbitrum One := Property{
		Name:         "Arbitrum One",
		Position:     17,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Algorand := Property{
		Name:         "Algorand",
		Position:     18,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Polygon POS := Property{
		Name:         "Polygon POS",
		Position:     19,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Opitmism := Property{
		Name:         "Opitmism",
		Position:     20,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Riot Blockchain := Property{
		Name:         "Riot Blockchain",
		Position:     21,
		Owner:        nil,
		NodeCount:    0,
	}

	Near Protocol := Property{
		Name:         "Near Protocol",
		Position:     22,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Kava := Property{
		Name:         "Kava",
		Position:     23,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Kujira := Property{
		Name:         "Kujira",
		Position:     24,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Archway := Property{
		Name:         "Archway",
		Position:     25,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Filecoin := Property{
		Name:         "Filecoin",
		Position:     26,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Dogecoin := Property{
		Name:         "Dogecoin",
		Position:     27,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Ethereum Classic := Property{
		Name:         "Ethereum Classic",
		Position:     28,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Polkadot := Property{
		Name:         "Polkadot",
		Position:     29,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Uniswap := Property{
		Name:         "Uniswap",
		Position:     30,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Fetch.ai := Property{
		Name:         "Fetch.ai",
		Position:     31,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Internet Computer := Property{
		Name:         "Internet Computer",
		Position:     32,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Cronos := Property{
		Name:         "Cronos",
		Position:     33,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Stella := Property{
		Name:         "Stella",
		Position:     34,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	MetaMask := Property{
		Name:         "MetaMask",
		Position:     35,
		Owner:        nil,
		NodeCount:    0,
	}
Fantom := Property{
		Name:         "Fantom",
		Position:     36,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Theta := Property{
		Name:         "Theta",
		Position:     37,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Thorchain := Property{
		Name:         "Thorchain",
		Position:     38,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Celestia := Property{
		Name:         "Celestia",
		Position:     39,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Aave := Property{
		Name:         "Aave",
		Position:     40,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Starknet := Property{
		Name:         "Starknet",
		Position:     41,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Jupiter := Property{
		Name:         "Jupiter",
		Position:     42,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	The Graph := Property{
		Name:         "The Graph",
		Position:     43,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount:    0,
	}
	Litecoin := Property{
		Name:         "Litecoin",
		Position:     44,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount: 0,
	}
	Shiba Inu := Property{
		Name:         "Shiba Inu",
		Position:     45,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount: 0,
	}

	}
	VeChain := Property{
		Name:         "VeChain",
		Position:     46,
		PurchaseCost: 150,
		NodeCost:     100,
		Rent:         []int{10, 30, 50, 100, 200},
		Owner:        nil,
		NodeCount: 0,
	}
	Tether := Property{
		Name:         "Tether",
		Position:     47,
		Owner:        nil,
		NodeCount: 0,
	// Print the properties






	fmt.Println(bitcoinProperty)
	fmt.Println(ethereum)
}
