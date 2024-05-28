package game

type Property struct{
Name string // Name of Property
Position int // Postion on Board
PurchaseCost int // cost to buy property
NodeCost int // cost to buy a node
Rent [] int // Rent cost
Owner *Player
NodeCount int // number of nodes owned
}

type Player struct{
	Name string // Name of Player
	Blance int // Blance of player
}
func main(){
bitcoinProperty := Property{
	Name:         "Bitcoin",
	Position:     0,
	PurchaseCost: 500,
	NodeCost:     250,
	Rent:         []int{},
	Owner:        nil,
	NodeCount:    0,
}
ethereum := Property
	Name: "Etherum",
	Postion: 1,
	PurchaseCost: 400,
	NodeCost:250,
	Rent[] int {75,225,300,500,800},
	Owner: nil,
	NodeCount: 0,

}

