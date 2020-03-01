package savings

type Savings {
	money int64
	withdrawlCount int //number of times someone has withdrawn
	
}

func NewSavings() *Savings{
    return &Savings{}
}