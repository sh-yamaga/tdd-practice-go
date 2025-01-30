package money

type Dollar struct {
	amount int
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.amount * multiplier}
}
