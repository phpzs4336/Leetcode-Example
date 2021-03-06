package minDepth

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 1
	} else {
		if root.Left != nil && root.Right != nil {
			return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right)))) + 1
		} else if root.Left != nil {
			return minDepth(root.Left) + 1
		} else if root.Right != nil {
			return minDepth(root.Right) + 1
		} else {
			return 1
		}
	}
}
