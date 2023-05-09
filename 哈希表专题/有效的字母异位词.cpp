力扣链接：https://leetcode.cn/problems/valid-anagram/

给定两个字符串 `*s*` 和 `*t*` ，编写一个函数来判断 `*t*` 是否是 `*s*` 的字母异位词。

**注意：**若 `*s*` 和 `*t*` 中每个字符出现的次数都相同，则称 `*s*` 和 `*t*` 互为字母异位词。

------------------------------ C++ -----------------------------------------------------
  // 整体思路（蛮巧妙的）
// 先遍历第一个字符串，把里面的每个元素-'a'(-'a'可得到连续的ascll码值)，然后存入hash数组中，即对应位置的元素++
// 再遍历第二个字符串，把里面每个元素-'a',在hash数组中对应位置的元素--
// 判断hash数组是否全为0即可
bool isAnagram(string s, string t) {
  // 总共就26个数
  vector<int> hash(26);
  if(s.size() != t.size()){
​    return false;
  }
  else{
​    for(int i = 0;i < s.size();i ++){
​      // hash数组中以元素值(s[i]-'a')为下标，对应的元素+1
​      hash[s[i] - 'a'] +=1;
​    }
​    for(int j = 0;j < t.size();j ++){
​      hash[t[j] - 'a'] -=1;
​    }
  }
  for(int k = 0;k < 26;k ++){
​    if(hash[k] != 0){
​      return false;
​    }
  }
  return true;
}

------------------------------ go -----------------------------------------------------
  func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
	// 定义数组
	record := [26]int{}
	// _:索引；r:值
	for _, r := range s {
		record[r-rune('a')]++
	}
	for _, r := range t {
		record[r-rune('a')]--
	}
	// 看看record是否和[26]int{}一样，都是空的
	return record == [26]int{}
}
