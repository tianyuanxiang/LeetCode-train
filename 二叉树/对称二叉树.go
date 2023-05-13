力扣链接：https://leetcode.cn/problems/symmetric-tree/

给你一个二叉树的根节点 `root` ， 检查它是否轴对称。
// 递归法
// 对称二叉树
func compare(left *TreeNode, right *TreeNode) bool {
	// 左右子树都为空
	if left == nil && right == nil {
		return true
	}
	// 左子树为空或右子树为空
	if left == nil || right == nil {
		return false
	}
	// 左右子树都不为空了，那就开始比较子树上的数值了
	if left.Val != right.Val {
		return false
	}
	// 如果左右子树都不为空，且数值相同的情况。
	// 比较二叉树外侧是否对称：传入的是左节点的左孩子，右节点的右孩子。
	// 比较内测是否对称，传入左节点的右孩子，右节点的左孩子。
	// 如果左右都对称就返回true ，有一侧不对称就返回false。
	// 处理外侧
	var outSide bool = compare(left.Left, right.Right)
	var inSide bool = compare(left.Right, right.Left)
	var issame bool = outSide && inSide
	return issame
}
func isSymmetric(root *TreeNode) bool {
    if root == nil {
		return false
	}
	return compare(root.Left, root.Right)
}

// 非递归法
func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	queue := list.New()
	queue.PushBack(root.Lchild) // 加入左子树节点
	queue.PushBack(root.Rchild) // 加入右子树节点
	for queue.Len() > 0 {
		leftNode := queue.Remove(queue.Front()).(*Node)
		rightNode := queue.Remove(queue.Front()).(*Node)
		// 如果左右节点都为空，说明到底了，直接再次循环
		if leftNode == nil && rightNode == nil {
			continue
		}
		// 如果两个节点有一个不为空，或者都不为空但数值不同，说明不对称
		if leftNode == nil || rightNode == nil || leftNode.val != rightNode.val {
			return false
		}
		// 以上情况都不存在，说明左右子树都在 && 值相等，接下来处理接下来的节点
		queue.PushBack(leftNode.Lchild)
		queue.PushBack(rightNode.Rchild)
		queue.PushBack(leftNode.Rchild)
		queue.PushBack(rightNode.Lchild)
	}
	return true
}
