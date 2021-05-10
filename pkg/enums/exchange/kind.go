package exchange

type Kind string

const (
	Topic  Kind = "topic"
	Fanout Kind = "fanout"
)

func (k Kind) ToString() string {
	return string(k)
}
