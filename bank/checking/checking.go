package checking

type Checking struct {
	money       int64
	timesUnder0 int
}

func NewChecking() *Checking {
	return &Checking{}
}
