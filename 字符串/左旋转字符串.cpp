力扣链接：https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
'''
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"
'''
  
#### 启示：
涉及字符串各种翻转问题的，都可从“双指针”和“分段翻转”维度进行思考。

```c++
string reverseLeftWords(string s, int n) 
{
  // 先翻转前n个字符
  reverse(s.begin(), s.begin() + n);
  // 再翻转n到最后个字符
  reverse(s.begin() + n, s.end());
  // 整体翻转
  reverse(s.begin(), s.end());
  return s;
}
```

```go
func reverseLeftWords(s string, n int) string {
    str := []byte(s)
	reverse(str[0:n])
	reverse(str[n:len(str)])
	reverse(str)
	return string(str)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
```

