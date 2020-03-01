package savings

type Savings struct {
	money          int64
	withdrawlCount int //number of times someone has withdrawn

}

func NewSavings() *Savings {
	return &Savings{}
}
