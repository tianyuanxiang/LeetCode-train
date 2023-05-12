力扣：https://leetcode.cn/problems/binary-tree-level-order-traversal/

----------------------- C++ ----------------------------------------------
// 二叉树的层序遍历
// 思路：
// 1. 建队列，从根节点开始遍历，循环队列不空，弹出队列的元素，找弹出的元素的左右孩子，分别入队
// 2. 用一个size记录当前层的节点数，每个节点的孩子们入队时，size++
// 3. 怎么遍历这一层呢？还是用size控制，size存储的是当前层的节点数，所以while(size--)，队列弹出元素

vector<vector<int>> levelOrder(Bitree root) 
{
  queue <Bitree> que;
  vector<vector<int>> result;
  if(root != NULL) que.push(root);
  while(!que.empty()){
​    int size = que.size();
​    vector<int> num;
​    while(size--){
​      Bitree cnt = que.front();
​      que.pop();
​      num.push_back(cnt->data);
​      if(cnt->lchild){
​        que.push(cnt->lchild);
​      }
​      if(cnt->rchild){
​        que.push(cnt->rchild);
​      }
​    }
​    result.push_back(num);
  }
  return result;
}

-------------------------------------- go ----------------------------------------------
func levelOrder(root *TreeNode) [][]int {
    // 一级结果数组
	res1 := [][]int{}
	if root == nil {
		return res1
	}
	queue := list.New() // 创建列表模拟队列
	queue.PushBack(root)
	// 二级结果数组
	res2 := []int{}
	// 呼应下方node.Lchild/Rchild,随遍随加
	for queue.Len() > 0 {
		size := queue.Len() // 控制每一层的遍历数量
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			res2 = append(res2, node.Val)
		}
		res1 = append(res1, res2)
		res2 = []int{} // 清空级结果数组的数据
	}
	return res1
}
