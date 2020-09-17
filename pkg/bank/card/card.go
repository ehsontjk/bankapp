package card
import ("bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	cards = []types.Card {
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
	var paymentSources []types.PaymentSource // == nil
    for _, v := range paymentSources {
         
            paymentSources = append(paymentSources,v)
        }
    
    return paymentSources
}