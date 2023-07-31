package binaryindexsearch

type Template2 struct {
	Nums []int
	Target int
}

func NewTemplate2(nums []int, target int) *Template2 {
	return &Template2{
		Nums: nums,
		Target: target,
	}
}

func (t2 *Template2) Search() int {
	left, right := 0, len(t2.Nums)-1

	for left < right {
		mid := left + (right - left) / 2

		if t2.Nums[mid] == t2.Target {
			return mid
		} else if t2.Nums[mid] < t2.Target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left != len(t2.Nums) && t2.Nums[left] == t2.Target {
		return left
	}
	return -1
}