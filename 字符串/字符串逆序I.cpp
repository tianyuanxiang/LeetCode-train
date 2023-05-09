力扣链接：https://leetcode.cn/problems/reverse-string/
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 `s` 的形式给出。

不要给另外的数组分配额外的空间，你必须**[原地](https://baike.baidu.com/item/原地算法)修改输入数组**、使用 O(1) 的额外空间解决这一问题。

#include <iostream>
#include <vector>
using namespace std;
void reverseString(vector<char>& s)
{
  int end = s.size() - 1;
  int start = 0;
  // 简单的双指针算法
  // 只要start和end没有错位，一直遍历即可
  while (start <= end)
  {
​    swap(s[start], s[end]);
​    start++;
​    end--;
  }
}

int main()
{
  char k;
  vector<char> str;
  while((k = getchar()) != '\n')
​    str.push_back(k);
  reverseString(str);
  for(int i = 0;i < str.size();i ++){
​    cout << str[i] << " ";
  }
  return 0;
}
