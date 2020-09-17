package card


import (
	"fmt"
	"bank/pkg/bank/types"

)
func ExamplePaymentSources() {
	cards := []types.Card {
		{
			
			PAN: "0001",
			Balance: 550,
			Active: true,
		},
		{
			
			PAN: "0002",
			Balance: 510,
			Active: true,
		
		},
		{
			
			PAN: "0011",
			Balance: 450,
			Active: true,
	},
}
paymentSources := PaymentSources(cards)
for _, v := range paymentSources {
			
	fmt.Println(v.Number)
}

//Output: [{0001}]
}