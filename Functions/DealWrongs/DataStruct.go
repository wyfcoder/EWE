package DealWrongs

type DealWrong struct {
	Wrong string
	Solve string
}

func (dealWrong *DealWrong) GetWrong() string{
	return dealWrong.Wrong
}

func (dealWrong *DealWrong) GetSolution() string {
	return dealWrong.Solve
}
