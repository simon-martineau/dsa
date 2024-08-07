package main

import (
	"github.com/simon-martineau/dsa/lc/ds"
	"math"
)

func main() {
}

func diameterOfBinaryTree(root *ds.TreeNode) int {
	_, total := diameterOfBinaryTreeRec(root)
	return total
}

func diameterOfBinaryTreeRec(root *ds.TreeNode) (longestBranch, total int) {
	if root == nil {
		return -1, -1
	}

	leftL, leftT := diameterOfBinaryTreeRec(root.Left)
	rightL, rightT := diameterOfBinaryTreeRec(root.Right)

	longestBranch = Max(leftL, rightL) + 1
	total = Max(leftT, rightT, longestBranch, leftL+rightL+1)
	return longestBranch, total
}

func Max(values ...int) int {
	m := math.MinInt
	for _, val := range values {
		if val > m {
			m = val
		}
	}
	return m
}
