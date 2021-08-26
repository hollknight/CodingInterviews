package T34_pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var nums []int
	var ans [][]int
	findSum(root, target, nums, &ans)

	return ans
}

func findSum(root *TreeNode, target int, nums []int, ans *[][]int) {
	if root == nil {
		return
	}

	nums = append(nums, root.Val)
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*ans = append(*ans, tmp)
	}
	findSum(root.Left, target, nums, ans)
	findSum(root.Right, target, nums, ans)
}
