力扣题目链接：https://leetcode.cn/problems/implement-queue-using-stacks/
---------------------------- CPP -------------------------------------
stack<int> stIn;
stack<int> stOut;
MyQueue() {
}
// 入队
void push(int x) {
  stIn.push(x);
}

// 出队
int pop() {
  if(stOut.empty()){
​    while(!stIn.empty()){
​      stOut.push(stIn.top());
​      stIn.pop();
​    }
  }
  // 直接把队首元素存储
  int result = stOut.top();
  stOut.pop();
  return result;
}

int peek() {
  int res = pop();
  // 你把人家pop出来了，不得给人家放回去
  // 没想到
  stOut.push(res);
  return res;
}
bool empty() {
  return stIn.empty() && stOut.empty();
}

----------------------------GO-------------------------------------
  type MyQueue struct {
	stackIn  []int // 输入栈
	stackOut []int // 输出栈
}

func Constructor() MyQueue {
	return MyQueue{
		// 创建切片
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0), // 输出栈
	}
}

// 入队列
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

// 出队列
func (this *MyQueue) Pop() int {
	inLen := len(this.stackIn)
	outLen := len(this.stackOut)
	// 出不去出去
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		for i := inLen - 1; i >= 0; i-- {
			// 腾挪
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		// 腾挪后清空
		this.stackIn = []int{}
		outLen = len(this.stackOut)
	}
	val := this.stackOut[outLen-1]
	this.stackOut = this.stackOut[0 : outLen-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
		return -1
	}
	// 为啥直接append就可以了？
	// peek在stackOut的最后面，不是0位置，append就回到原来的位置上了
	this.stackOut = append(this.stackOut, val)
	return val

}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
