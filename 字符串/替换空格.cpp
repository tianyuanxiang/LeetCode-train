力扣链接：https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

请实现一个函数，把字符串 `s` 中的每个空格替换成"%20"。
void replaceSpace(string s)
{
  int ret = 0;
  for(int i = 0;i < s.size();i ++){
​    if(s[i] == ' ') ret ++;
  }
  int oldSize = s.size();
  // 为啥扩充4个而不是6个？ 其实就是6个，但是空格的个数已经占去几个蹲位了。
  s.resize(s.size() + 2 * ret);
  int newSize = s.size();

  // j大，i小
  // 当i = j 的时候，说明前面已经没有空格了，后面的i--, j--表示移位完事后，指向新的待处理的数
  for(int i = oldSize, j = newSize; i < j; i--, j--)
  {
​    if(s[i] != ' '){
​      s[j] = s[i];
​    }
​    else{
​      s[j] = '0';
​      s[j - 1] = '2';
​      s[j - 2] = '%';
​      j -= 2;
​    }
  }
}

int main()
{
  char s;
  string str;
  while((s = getchar()) != '\n')
​    str += s;
  replaceSpace(str);
  for(int i = 0;i < str.size();i ++){
​    cout << str[i] << " ";
  }
  return 0;
}
