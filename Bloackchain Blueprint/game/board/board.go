package game

type Board struct {
	BoardSpace [48]Square
}

type Square struct {
}

const (
	BuildableProperty int = iota
	DAO
	MiningPool
	CommunityNode
	SmartContract
	TransactionFee
	HODL
	SlashedCards
)

func GetPropertyType(number int) string {
	var propType string
	switch number {
	case BuildableProperty:
		propType = "Property"
	case DAO:
		propType = "DAO"
	case MiningPool:
		propType = "Mining Pool"
	case CommunityNode:
		propType = "Community Node"
	case SmartContract:
		propType = "Smart Contract"
	case TransactionFee:
		propType = "Transaction Fees"
	case HODL:
		propType = "HODL"
	case SlashedCards:
		propType = "Missed Block"

	}
	return propType
}
