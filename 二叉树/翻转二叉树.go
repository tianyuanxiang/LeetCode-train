力扣链接：https://leetcode.cn/problems/invert-binary-tree/

给你一棵二叉树的根节点 `root` ，翻转这棵二叉树，并返回其根节点


// 翻转二叉树(递归写法)
func invertTree(root *Node) *Node {
	if root == nil {
		return root
	}
	// 注意Go语言的交换(和python一样)
	root.Lchild, root.Rchild = root.Rchild, root.Lchild
	invertTree(root.Lchild)
	invertTree(root.Rchild)
	return root
}


