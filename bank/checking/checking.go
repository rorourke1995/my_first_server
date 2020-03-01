package checking

type Checking struct {
	money int64
	timesUnder0 int
	
}

func NewChecing() *Checking{
	return &Checking{}
}