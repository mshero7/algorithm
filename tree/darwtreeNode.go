package tree

type drawTreeNode[T any] struct {
	Value T

	// 좌표
	X int
	Y int

	Child []*drawTreeNode[T]
}

func makeDrawTree[T any](node *TreeNode[T], level int, order *int) *drawTreeNode[T] {
	if node == nil {
		return nil
	}

	drawNode := &drawTreeNode[T]{
		Value: node.Value,
		Y:     level,
		X:     *order,
	}

	(*order)++

	// in-order
	half := len(node.Childs) / 2
}
