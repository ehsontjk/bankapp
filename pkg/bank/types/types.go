package types
type Money int64
type Balance int64
type Currency string
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Card struct {
	ID int
	PAN string
	Currency Currency
	Balance Money
	Color string
	Name string
	Active bool
	MinBalance Money
	

    
}
type PaymentSource struct {
	Type string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}