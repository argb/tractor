package httpHandler

type Node struct {
	pattern string
	part string
	children []*Node
	isWild bool
}
