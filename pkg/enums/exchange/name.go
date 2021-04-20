package exchange

type Name string

const (
	NewAnalysis Name = "new-analysis"
)

func (n Name) ToString() string {
	return string(n)
}
