// 力扣题目链接：https://leetcode.cn/problems/evaluate-reverse-polish-notation/
// 逆波兰表达式
// 逆波兰表达式的形成其实就是二叉树后序遍历的方式
// 思路：
// 1. 创建一个栈，然后遍历字符串，如果字符等于符号，就弹出两个元素做相应运算，然后根据具体的符号类型，将得出的结果再存入栈中
// 2. 如果非符号，直接入栈

------------------------  C++ --------------------------
int evalRPN(vector<string>& tokens) 
{
  stack<long long> st;
  for(int i = 0;i < tokens.size();i ++){
​    if(tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/"){
​      long long num1 = st.top();
​      st.pop();
​      long long num2 = st.top();
​      st.pop();
    
​      // 我草，这里还有个坑
​      // 因为是栈，所以顺序不能错
​      if(tokens[i] == "+") st.push(num2 + num1);
​      if(tokens[i] == "-") st.push(num2 - num1);
​      if(tokens[i] == "*") st.push(num2 * num1);
​      if(tokens[i] == "/") st.push(num2 / num1);
​    }
​    else{
​      // 类型转换在这呢
​      // stoi():将字符串转换为int型
​      // stoll():将字符串转换为long long
​      // stof():将字符串转换为float型
​      // stod():将字符串转换为double型
​      st.push(stoll(tokens[i]));
​    }
  }
  // 说实话还是没理解透：最后的结果其实就是处理完后栈中的唯一元素
  // 2023.2.6：对啊，每两个数都做运算，最后可不就剩一个数了吗
  int result = st.top();
  st.pop();
  return result;
}
int main()
{
  vector<string> str{"10","6","9","3","+","-11","*","/","*","17","+","5","+"};
  int k = evalRPN(str);
  return 0;
}
-------------------------- go ----------------------------------------
//	逆波兰表达式求值（二叉树的后序遍历）
// 计算机将其转换为逆波兰表达式进行运算，怎么转换的？对用户输入数据进行二叉树的后序遍历，二叉树怎么来的？不知道
// 思路：
// 1.设置栈，遍历字符串，非运算符就入栈
// 2.遇到运算符，判断是什么运算符，从栈里弹出两个元素进行操作，将得出来的数再次入栈（重要！）
// 3.运算符如何判断？（两层判断：1.总体 2.具体）字符和数字如何转换(重新存储数字)？
func evalRPN(tokens []string) int {
	st := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			var num2 int = st[len(st)-1]
			var num1 int = st[len(st)-2]
			st = st[:len(st)-2]
			if tokens[i] == "+" {
				var sum int = num1 + num2
				st = append(st, sum)
			} else if tokens[i] == "-" {
				var sum int = num1 - num2
				st = append(st, sum)
			} else if tokens[i] == "*" {
				var sum int = num1 * num2
				st = append(st, sum)
			} else if tokens[i] == "/" {
				var sum int = num1 / num2
				st = append(st, sum)
			}
		} else {
			val, _ := strconv.Atoi(tokens[i])
			st = append(st, val)
		}
	}
	return st[0]
}
