package money

type Money struct {
	amount   int
	currency Currency
}

type Currency string

const (
	USD Currency = "USD"
	CHF Currency = "CHF"
)

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

// Plus returns a new Sum instance with the sum of the current instance and the addend
func (m Money) Plus(addend Money) Sum {
	return Sum{
		augend: m,
		addend: addend,
	}
}

// Reduce converts the Money instance to the specified currency and returns a new Money instance
func (m Money) Reduce(to Currency) Money {
	return NewMoney(m.amount, to)
}

// NewMoney returns a new Money instance with the specified amount and currency.
func NewMoney(amount int, currency Currency) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

// NewDollar returns a new Dollar instance have the specified amount.
func NewDollar(amount int) Money {
	return NewMoney(amount, USD)
}

// NewFranc returns a new Franc instance have the specified amount.
func NewFranc(amount int) Money {
	return NewMoney(amount, CHF)
}
