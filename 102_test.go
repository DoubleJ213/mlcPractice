package mleetcode

import (
	"fmt"
	"testing"
)

// 二叉树，BFS

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	// 递归 实现
	res := make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

func TestAl102(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(levelOrder(root))
	fmt.Println("done")

}
