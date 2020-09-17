package card
import ("bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	payment := []types.PaymentSource{}
	for _, card:= range cards{
		if card.Balance>0 && card.Active {
			payment=append(payment, types.PaymentSource{card.Name, card.PAN, card.Balance})
			//payment.Balance=card.Balance
			//payment.Number=card.Number
		}
	}
	return payment
}