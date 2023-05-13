力扣链接：https://leetcode.cn/problems/count-complete-tree-nodes/description/


// 完全二叉树的节点个数
// 1. 直接层序遍历
func countNodes1(root *Node) int {
	if root == nil {
		return 0
	}
	ret := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		ret += size
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Lchild != nil {
				queue.PushBack(node.Lchild)
			}
			if node.Rchild != nil {
				queue.PushBack(node.Rchild)
			}
		}
	}
	return ret
}

// 2. 充分利用完全二叉树的性质
// 在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值(满了)，
// 并且最下面一层的节点都集中在该层最左边的若干位置(先左再右)。若最底层为第 h 层，则该层包含 1~ 2^(h-1)个节点。
// 分两种情况处理：
// 整体思路：不断细分最小完全二叉树

// 1. 全满，直接2^(h-1)，h是二叉树的深度
// 2. 最后一层不满，则分别递归左孩子和右孩子，递归到某一深度一定会有左孩子或者右孩子为满二叉树（极端情况下为最后一层），然后依然可以按照情况1来计算。
// 那么，如何递归到某一深度呢？
// 递归三部曲之
func countNodes2(root *Node) int {
	// 递归出口1:根节点为空
	if root == nil {
		return 0
	}
	// 递归出口2:左右子树深度相同
	// 根据左深度和右深度是否相同来判断该子树是不是满二叉树
	var LeftTree *Node = root.Lchild
	var RightTree *Node = root.Rchild
	LeftDepth := 0
	RightDepth := 0
	for LeftTree != nil {
		LeftTree = LeftTree.Lchild
		LeftDepth++
	}
	for RightTree != nil {
		RightTree = RightTree.Rchild
		RightDepth++
	}
	if LeftDepth == RightDepth {
		return (2 << LeftDepth) - 1
	}
	// 单层循环逻辑
	leftTreeNum := countNodes2(root.Lchild)  // 左
	rightTreeNum := countNodes2(root.Rchild) // 右
	result := leftTreeNum + rightTreeNum + 1 // 中
	return result

}

