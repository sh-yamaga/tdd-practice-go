package money

type Money struct {
	amount   int
	currency string
}

// Equals compares whether two money instances have the same amount.
func (m Money) Equals(in Money) bool {
	return m == in
}

// Times returns a new Money instance multiplied by the specified rate.
func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// NewMoney returns a new Money instance with the specified amount and currency.
func NewMoney(amount int, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

// NewDollar returns a new Dollar instance have the specified amount.
func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

// NewFranc returns a new Franc instance have the specified amount.
func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}
