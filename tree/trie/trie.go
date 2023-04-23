package trie

type Node struct {
	Children map[rune]*Node // Tree Node Child
	Value    string         // Node가 갖는 값
}

func NewNode(value string) *Node {
	return &Node{
		Children: make(map[rune]*Node),
		Value:    value,
	}
}

func Insert(root *Node, key string) bool {
	runes := []rune(key)
	cur := root
	for i, c := range runes {
		node := cur.Children[c]
		if node == nil {
			node = NewNode(string(runes[:i+1]))
			cur.Children[c] = node
		}
		cur = node
	}
	return true
}

func AutoComplete(root *Node, key string) string {
	runes := []rune(key)
	cur := root

	for _, c := range runes {
		if n, ok := cur.Children[c]; ok {
			cur = n
		} else {
			return ""
		}
	}
	for {
		if len(cur.Children) == 0 {
			break
		}
		for _, n := range cur.Children {
			cur = n
			break
		}
	}

	return cur.Value
}
