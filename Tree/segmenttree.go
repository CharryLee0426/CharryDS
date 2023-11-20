package tree

/*
SegmentTree is a data structure that can efficiently answer range queries.
It is a binary tree where each node represents an interval of the array.
The root node represents the entire array, and each child node represents
the left and right halves of the array.
*/

type SegTree struct {
	// The start and end indices of the array
	start, end int
	left, right *SegTree
	sum int
}

func GenerateSegTree(root *SegTree, nums []int) {
	if root.start == root.end {
		root.sum = nums[root.start]
		return
	}
	mid := root.start + ((root.end - root.start) >> 1)
	root.left = &SegTree{root.start, mid, nil, nil, 0}
	root.right = &SegTree{mid + 1, root.end, nil, nil, 0}
	GenerateSegTree(root.left, nums)
	GenerateSegTree(root.right, nums)
	root.sum = root.left.sum + root.right.sum
}

func UpdateSegTree(root *SegTree, index, val int) {
	if root.start == root.end {
		root.sum = val
		return
	}
	mid := root.start + ((root.end - root.start) >> 1)
	if index <= mid {
		UpdateSegTree(root.left, index, val)
	} else {
		UpdateSegTree(root.right, index, val)
	}
	root.sum = root.left.sum + root.right.sum
}

func QuerySegTree(root *SegTree, start, end int) int {
	if root.start == start && root.end == end {
		return root.sum
	}

	if root.start > end || root.end < start {
		return 0
	}

	l := QuerySegTree(root.left, start, end)
	r := QuerySegTree(root.right, start, end)

	return l + r
}