// 有效的括号
// 思路：

// 说实话还是有点迷糊，比如字符串是“({[]})”、“({[])})、({{]]})”那用这个方法肯定没问题
// 如果出现“([{)]}”或“)[]({}”这样就没法搞啊。除非你有约束条件：所有的左右括号都是根据类型以同样顺序排列，那就假设这样吧
// 1.当出现左号时，把右号放进去
// 2.遇到右号时，判断栈顶元素是不是相等
// 3.三种意外情况：（1）字符串全部遍历完了，但栈不空 （2）右边没有与左匹配的号 （3）遍着遍着栈空了，字符串还有
// 总结：两个要点，三个细节:括号反着放，别墅靠大海；三种问题有，你就跟我走

bool isValid(string s) 
{
  if(s.size() % 2 != 0)
​    return false;
  stack<char> st;
  for(int i = 0;i < s.size();i ++){
​    if(s[i] == '(') st.push(')');
​    else if(s[i] == '[') st.push(']');
​    else if(s[i] == '{') st.push('}');

​    // 第三种情况：遍历字符串匹配的过程中，栈已经为空了，没有匹配的字符了，说明右括号没有找到对应的左括号 return false
​    // 第二种情况：遍历字符串匹配的过程中，发现栈里没有我们要匹配的字符。所以return false 
​    // 一行搞定
​    else if(st.empty() || st.top() != s[i]){
​      return false;
​    }
​    // 满足
​    else{
​      st.pop();
​    }
  }
  // 第一种情况，直接看空否
  return st.empty();
}

--------------------------GO-----------------------------------
  // 有效地括号
// 三种情况：
// 3.三种意外情况：（1）字符串全部遍历完了，但栈不空(左边多了) （2）右边没有与左匹配的号(不匹配) （3）遍着遍着栈空了，字符串还有(右边多了)
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0)
	// 右方向是key,左方向是value
	// 由于go语言字符串和int转换费劲，借助map进行巧妙转换
	val := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			// 只要进入这个if，说明指定遍历着右括号了，开始判断
			// len(stack) > 0:第二、三种情况
		} else if len(stack) > 0 && val[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
