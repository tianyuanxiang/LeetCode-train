力扣链接：https://leetcode.cn/problems/reverse-string-ii/

给定一个字符串 `s` 和一个整数 `k`，从字符串开头算起，每计数至 `2k` 个字符，就反转这 `2k` 字符中的前 `k` 个字符。

- 如果剩余字符少于 `k` 个，则将剩余字符全部反转。
- 如果剩余字符小于 `2k` 但大于或等于 `k` 个，则反转前 `k` 个字符，其余字符保持原样。

------------------------------- C++ ------------------------------
  
// 遍历时，以2k为间隔往后跑
void reverseString_flower(vector<char> &s, int k)
{
  for(int i = 0;i < s.size();i += (2*k)){
​    // 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
​    if(i + k <= s.size()){
​      // 难点
​      // 一次翻转k个字符，而不是2k个字符
​      reverse(s.begin() + i, s.begin() + i + k);
​    }
​    // 剩余字符少于 k 个
​    else{
​      reverse(s.begin() + i, s.end());
​    }
  }
}
int main()
{
  char s;
  int k;
  vector<char> str;
  while((s = getchar()) != '\n')
​    str.push_back(s);
  cin >> k;
  reverseString_flower(str, k);
  for(int i = 0;i < str.size();i ++){
​    cout << str[i] << " ";
  }
  return 0;
}
------------------------------- GoLang ------------------------------
  func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
func reverseStr(s string, k int) string {
	// 把字母组成的字符串搞成byte数字
	str := []byte(s)
	length := len(s)
	for i := 0; i < length; i += (2 * k) {
		// 1. 每隔2k个字符的前k个字符进行反转
		// 2. 剩余字符小于2k但大于或等于k个，则反转前k个字符
		// 上面这俩都是反转k个字符
		if i+k < length {
			reverse(str[i : i+k])
		} else {
			reverse(str[i:length])
		}
	}
	return string(str)
}
