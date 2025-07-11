package tree

type Dict interface {
	Search(value int) bool
	Insert(value int) bool
	Delete()
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
	path   *[]string
}

func Search(node *Node, value int) bool {
	if node.Val == value {
		return true
	}
	if value < node.Val && node.Left != nil {
		Record(node, "move left")
		node.Left.Parent = node
		return Search(node.Left, value)
	} else if value > node.Val && node.Right != nil {
		Record(node, "move right")
		node.Right.Parent = node
		return Search(node.Right, value)
	}

	return false
}

func Insert(node *Node, value int) bool {
	if node.Val == value {
		Record(node, "value exists")
		return false
	} else if value < node.Val {
		if node.Left != nil {
			// move left
			Record(node, "move left")
			node.Parent = node
			return Insert(node.Left, value)
		} else {
			// if left is nil, insert left
			Record(node, "insert left")
			node.Left = &Node{Val: value, path: node.path}
			return true
		}
	} else if value > node.Val {
		if node.Right != nil {
			// move right
			Record(node, "move right")
			node.Parent = node
			return Insert(node.Right, value)
		} else {
			// if right is nil, insert right
			Record(node, "insert right")
			node.Right = &Node{Val: value, path: node.path}
			return true
		}
	}

	Record(node, "unknown error")
	return false
}

func ASCTraversal(node *Node, path *[]int) {
	// in order traversal, left, self, right
	if node.Left != nil {
		Record(node, "move left")
		node.Left.path = node.path
		ASCTraversal(node.Left, path)
	}

	Record(node, "self")
	(*path) = append((*path), node.Val)

	if node.Right != nil {
		Record(node, "move right")
		node.Right.path = node.path
		ASCTraversal(node.Right, path)
	}
}

func DESCTraversal(node *Node, path *[]int) {
	// traversal so the return path order is DESC
	if node.Right != nil {
		Record(node, "move right")
		node.Right.path = node.path
		DESCTraversal(node.Right, path)
	}

	Record(node, "self")
	(*path) = append((*path), node.Val)

	if node.Left != nil {
		Record(node, "move left")
		node.Left.path = node.path
		DESCTraversal(node.Left, path)
	}
}

func Delete(node *Node, value int) {
	panic("unhandle error")
}

func Record(node *Node, operation string) {
	(*node.path) = append((*node.path), operation)
}
