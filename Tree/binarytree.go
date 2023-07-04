package tree

// define a node which has left child and right child
type Node struct {
	Val int
	Left *Node
	Right *Node
}

func NewFromArray(arr []int) *Node {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	root := &Node{Val: arr[0]}
	queue := []*Node{root}

	for i := 1; i < len(arr); i++ {
		queue = append(queue, &Node{Val: arr[i]})
	}

	for i := 0; i < len(arr); i++ {
		if 2 * i + 1 < len(arr) {
			queue[i].Left = queue[2 * i + 1]
		}

		if 2 * i + 2 < len(arr) {
			queue[i].Right = queue[2 * i + 2]
		}
	}

	return root
}

func (n *Node) PreOrder() []int {
	if n == nil {
		return nil
	}

	res := []int{n.Val}
	res = append(res, n.Left.PreOrder()...) // ... means unpacking
	res = append(res, n.Right.PreOrder()...)

	return res
}

func (n *Node) Inorder() []int {
	if n == nil {
		return nil
	}

	res := n.Left.Inorder()
	res = append(res, n.Val)
	res = append(res, n.Right.Inorder()...)

	return res
}

func (n *Node) Postorder() []int {
	if n == nil {
		return nil
	}

	res := n.Left.Postorder()
	res = append(res, n.Right.Postorder()...)
	res = append(res, n.Val)

	return res
}

func (n *Node) Levelorder() []int {
	if n == nil {
		return nil
	}

	res := []int{n.Val}

	queue := []*Node{n.Left, n.Right}

	for len(queue) > 0 {

		// get the first element of the queue and pop it
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}

		// add the poped element to the result
		// and add its left and right child to the queue
		res = append(res, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	return res
}

func (n *Node) CheckSameTree(other *Node) bool {
	if n == nil && other == nil {
		return true
	}

	if n == nil || other == nil {
		return false
	}

	if n.Val != other.Val {
		return false
	}

	return n.Left.CheckSameTree(other.Left) && n.Right.CheckSameTree(other.Right)
}

func (n *Node) CheckFather(suspectson *Node) bool {
	if n == nil {
		return false
	}

	// tree conditions, it will be true when anyone is true
	// 1. n is the father of suspected son
	// 2. n's left subtree is the father of suspected son
	// 3. n's right subtree is the father of suspected son
	return checksame(n, suspectson) || n.Left.CheckFather(suspectson) || n.Right.CheckFather(suspectson)
}

// tool function: check if two trees' structure and values are the same
func checksame(treeA, treeB *Node) bool {
	if treeA == nil && treeB == nil {
		return true
	}

	if treeA == nil || treeB == nil {
		return false
	}

	if treeA.Val != treeB.Val {
		return false	
	}
	return checksame(treeA.Left, treeB.Left) && checksame(treeA.Right, treeB.Right)
}