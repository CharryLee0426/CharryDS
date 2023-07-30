package binaryindexsearch

type Template1 struct {
	Nums []int
	Target int
}

func NewTemplate1(nums []int, target int) *Template1 {
	return &Template1{
		Nums: nums,
		Target: target,
	}
}

func (t1 *Template1) Search() int {
	left, right := 0, len(t1.Nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if t1.Nums[mid] == t1.Target {
			return mid
		} else if t1.Nums[mid] < t1.Target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}