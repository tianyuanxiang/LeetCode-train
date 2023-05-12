
------------------------ C++ --------------------------------
// 前序非递归遍历
// 思路：
// 针对一棵二叉树，从首节点开始先入栈，然后依次入右节点和左节点（栈先弹出左节点，判断左节点是否还有孩子）
// 依次弹出每个子节点，将其转化为父节点来判断是否有孩子，直到栈为空
vector<int> preorderTraversal(Bitree root) {
  stack<Bitree> st;
  vector<int> result;
  if(root == NULL){
​    return result;
  }
  // 又晕了，root存的不是整颗树，而是头结点的地址
  st.push(root);
  while(!st.empty()){
​    Bitree bi = st.top();
​    st.pop();
​    result.push_back(bi->data);
​    if(bi->rchild){
​      st.push(bi->rchild);
​    }
​    if(bi->lchild){
​      st.push(bi->lchild);
​    }
  }
  return result; 
}

// 中序非递归遍历（左中右）
// 思路：(没理解透)
// 为什么中序这么特殊？前后序在遍历的时候，都是以根节点为中心，找左右孩子，也就是说：每个根是一层
// 你遍历到的点，就是你要访问的点，（要访问的元素和要处理的元素顺序是一致的，都是中间节点）但是中序遍历是一直向左
// 借助指针和栈进行存储
// 设置一个指针指向头结点，空栈一个
// 指针循环往左,每搂到一个元素就放栈里，一旦左为空(else)，指针指向栈头(很重要！更新数)弹栈存数，指针指向右子树，如果右子树也为空，继续弹栈存数
// 那显而易见，循环条件就是栈不空或cur(当前指向的元素)不空
vector<int> inorderTraversal(Bitree root) 
{
	stack<Bitree> st;
	vector<int> result;
	BiTNode *cur = root;
	while(cur != NULL || !st.empty()){
		if(cur){
			st.push(cur);
			cur = cur->lchild;
		}else{
			// 及时存储
			cur = st.top();
			st.pop();
			result.push_back(cur->data);
			cur = cur->rchild;
		}
	}
	return result;
}

// 后序非递归遍历
// 思路：
// 前序遍历是中左右，后序是左右中，那就把前序遍历改成中右左，然后直接翻转就行了（亲测有用）
vector<int> postorderTraversal(Bitree root) {
  stack<Bitree> st;
  vector<int> result2;
  if(root == NULL) return result2;
  st.push(root);
  while(!st.empty()){
​    Bitree BI = st.top();
​    st.pop();
​    result2.push_back(BI->data);
​    if(BI->lchild){
​      st.push(BI->lchild);
​    }
​    if(BI->rchild){
​      st.push(BI->rchild);
​    }
  }
  reverse(result2.begin(), result2.end());
  return result2;
}  

------------------------ go --------------------------------
package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	val    int
	Lchild *Node
	Rchild *Node
}

// CreatTree 构建树并且返回树的头节点、arr：树节点的扁平化数组，-1表示空节点、index：当前节点在数组中的下标、n:数组长度
func CreatTree(arr []int, index, n int) *Node {
	if n == 0 || index >= n || arr[index] == -1 {
		return nil
	}
	return &Node{
		val:    arr[index],
		Lchild: CreatTree(arr, index*2+1, n),
		Rchild: CreatTree(arr, index*2+2, n),
	}
}

// 二叉树的非递归前序遍历
// 思路：和层序遍历一样，从根节点开始，先遍历右孩子，再遍历左孩子，并分别入栈
// 然后弹出栈头元素，并将其val保存至结果数组，之后再遍历该元素的左右孩子，直到栈为空
func preorderTraversal(root *Node) []int {
	// 创建结果数组
	ans := []int{}
	if root == nil {
		return ans
	}
	// 创建列表模拟栈
	st := list.New()
	// 把头结点搞里头
	st.PushBack(root)
	for st.Len() > 0 {
		// 把st.Back()先保存在node里，然后从st里移除
		node := st.Remove(st.Back()).(*Node)
		ans = append(ans, node.val)
		if node.Rchild != nil {
			st.PushBack(node.Rchild)
		}
		if node.Lchild != nil {
			st.PushBack(node.Lchild)
		}
	}
	return ans
}
// 后序非递归遍历
// 思路：
// 前序遍历是中左右，后序是左右中，那就把前序遍历改成中右左，然后直接翻转就行了

func main() {
	arr := []int{0, 1, 2, 3, 4, 1, 6}
	root := CreatTree(arr, 0, 7)
	res := preorderTraversal(root)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])

	}
}

// 二叉树的非递归中序遍历（左中右）
// 其实还好，主要问题在于“访问的节点并不是要处理的节点”，要处理的节点是啥？（从根节点开始访问，但最左边的节点才是处理的）在哪？（一路向左）
// 思路：
// 1.定义一个栈；一个装结果的数组；一个放节点的变量。设置for循环，条件为：当前节点不为空||栈不空
// 2.如果当前节点不为空，一路向左并把元素放入栈中
// 3.如果当前节点为空，说明到底了，把该节点的数存为结果，并从栈里弹出，然后<处理该节点的右节点>关键。如果右节点也为空，那么就要重新从栈里找数了
func inorderTraversal(root *Node) []int {
	// 创建结果数组
	ans := []int{}
	if root == nil {
		return ans
	}
	// 创建列表模拟栈
	st := list.New()
	// 把头结点搞里头
	// st.PushBack(root)
	cur := root
	for st.Len() > 0 || cur != nil {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Lchild
		} else {
			// 左节点为空且右节点也为空，说明是子树，接下来就要处理栈里的了
			cur = st.Remove(st.Back()).(*Node)
			ans = append(ans, cur.val)
			cur = cur.Rchild
		}
	}
	return ans
}
