package fxxk

import (
	"fmt"
	"testing"
)

/*
235. 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：
“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

说明:
所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

/*
所有节点的值都是唯一的。所以这题不要存每个节点对象了，直接值判断
p、q 为不同节点且均存在于给定的二叉搜索树中。不要判断 p q是否存在root 的树中了
那么怎么找两个节点的深度最深的公共祖先。后序遍历然后判断值是否等于，或者子树是否存在这个数？
*/

// 二叉搜索树
var node235 *TreeNode

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	node235 = new(TreeNode)
	if root != nil {
		dfs235(root, p, q)
	}

	return node235
}

//   4
//  8 5
// 0 1  6
/*
pin 表示p是否存在子树中
qin 表示q是否存在子树中*/
func dfs235(root *TreeNode, p *TreeNode, q *TreeNode) {
	cur := root.Val
	if cur > q.Val && cur > p.Val {
		dfs235(root.Left, p, q)
	}

	if cur < q.Val && cur < p.Val {
		dfs235(root.Right, p, q)
	}

	if (cur >= q.Val && cur <= p.Val) || (cur >= p.Val && cur <= q.Val) {
		node235 = root
	}

	return

}

func TestAl235(t *testing.T) {
	//二叉搜索树,测试样例别给错了
	root7 := &TreeNode{9, nil, nil}
	//root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{4, nil, nil}
	root4 := &TreeNode{0, nil, nil}
	root3 := &TreeNode{8, nil, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{6, root2, root3}
	fmt.Println(lowestCommonAncestor235(root, root4, root5))
}
