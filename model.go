package mongorm

type Model interface {
	Collection() string
}
