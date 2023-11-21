package entities

const OperationBuy string = "BUY"
const OperationSell string = "SELL"

type Operation struct {
	Price  int64
	Symbol string
	Type   string
}
