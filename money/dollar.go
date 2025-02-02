package money

type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{
		Money{
			amount:   amount,
			currency: "USD",
		},
	}
}

// Times returns a new Dollar instance multiplied by the specified rate.
func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}
