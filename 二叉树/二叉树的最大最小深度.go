最大深度力扣链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/

给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
**说明:** 叶子节点是指没有子节点的节点。
**示例：**
给定二叉树 [3,9,20,null,null,15,7]，

// 二叉树的最大深度(递归)
// 1.确定递归函数返回值和参数
// 2.递归出口
// 3.单层递归逻辑
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	var depth int = max(left, right) + 1
	return depth
}

最小力扣链接：https://leetcode.cn/problems/minimum-depth-of-binary-tree/

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

**说明：**叶子节点是指没有子节点的节点。
// 二叉树的最小深度(递归)
// 1.确定递归函数返回值和参数
// 2.递归出口
// 3.单层递归逻辑
func minDepth(root *Node) int {
	if root == nil {
		return 0
	}
	// 不能这样写，这样会把没有左孩子的分支算为最短深度，而不考虑右子树真正的深度
	Left := minDepth(root.Lchild)
	Right := minDepth(root.Rchild)
	// var depth int = min(Left, Right) + 1

	// 应该这么来
	if root.Lchild == nil && root.Rchild != nil {
		return Right + 1
	}
	if root.Lchild != nil && root.Rchild == nil {
		return Left + 1
	}
	ret := min(Left, Right) + 1
	return ret
}
