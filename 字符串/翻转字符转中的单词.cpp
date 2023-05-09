力扣链接：https://leetcode.cn/problems/reverse-string/

给你一个字符串 `s` ，请你反转字符串中 **单词** 的顺序。
**单词** 是由非空格字符组成的字符串。`s` 中使用至少一个空格将字符串中的 **单词** 分隔开。
返回 **单词** 顺序颠倒且 **单词** 之间用单个空格连接的结果字符串。
**注意：**输入字符串 `s`中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

#### 移除空格：
```c++
void removeExtraSpaces(string &s)
{
  int woman = 0;
  for(int man = 0; man < s.size(); man++){
​    // 夫指针找不是空格的元素
​    if(s[man] != ' '){
​      // 夫指针刚开始打天下，妻指针肯定指向开头，因为得守老家,只有夫指针打下了一个单词，woman才能从家走
​      // 夫指针拿下一个单词之后，开始拿第二个了（下面的循环结束了），妻指针安营扎寨(放个空格)
​      if(woman != 0){
​        s[woman] = ' ';
​        woman++;
​      }
​      // 循环结束，第一个单词拿下
​      while(man < s.size() && s[man] != ' '){
​        s[woman] = s[man];
​        woman++;
​        man++;
​      }
​    }
  }
  s.resize(woman);
}
```
#### 主函数：
```c++
string reverseWords(string s) {
​    removeExtraSpaces(s);
​    reverse(s, 0, s.size() - 1);
​    int start = 0;
​    for(int i = 0;i <= s.size();i ++){
​      if(i == s.size() || s[i] == ' '){
​        reverse(s, start, i-1);
​        start = i + 1;
​      }
  }
  return s;
  }
```
#### 翻转字符串
```c++
void reverse(string& s, int start, int end)
  {
​    for(int i = start,j = end;i <j;i ++,j --){
​      swap(s[i], s[j]);
​    }
  }
```

