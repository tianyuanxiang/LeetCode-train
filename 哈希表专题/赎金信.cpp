力扣链接：https://leetcode.cn/problems/ransom-note/

给你两个字符串：`ransomNote` 和 `magazine` ，判断 `ransomNote` 能不能由 `magazine` 里面的字符构成。

如果可以，返回 `true` ；否则返回 `false` 。

`magazine` 中的每个字符只能在 `ransomNote` 中使用一次。
-------------------------------------  C++ -------------------------------------------
// 赎金信

// 给两个字符串A、B，判断A能不能由B里的字符串组成，即B是否包含A

// 当我们遇到了要快速判断一个元素是否出现集合里的时候，就要考虑哈希法了。
bool canConstruct(string ransomNote, string magazine) 
{
  // 暴力
  // for(int i = 0;i < magazine.length();i ++){
  //   for(int j = 0;j < ransomNote.length();j ++){
  //     if(magazine[i] == ransomNote[j]){
  //       ransomNote.erase(ransomNote.begin() + j);
  //       break;
  //     }
  //   }
  // }
  // if(ransomNote.length() == 0){
  //   return true;
  // }
  // return false;
  
  // 哈希表
  // 思路：
  // 1、因为是小写字母，所以定义一个有26个元素的数组，初始值为0，将母数组中所有元素的出现次数作为数组元素值存入数组
  // 2、然后遍历子数组的每个值，在对应的record数组中--, 然后遍历record数组，如果存在小于0的元素，直接返回flase,否则就是true
  int record[26] = {0};
  if(ransomNote.length() > magazine.length()){
​    return false;
  }
  for(int i = 0;i < magazine.length();i ++){
​    record[magazine[i] - 'a'] += 1;
  }
  for(int j = 0;j < ransomNote.length();j ++){
​    record[ransomNote[j] - 'a'] -= 1;
  }
  for(int k = 0;k < 26;k ++){
​    if(record[k] < 0){
​      return false;
​      break;
​    }
  }
  return true;
}

----------------------------------------------- go----------------------------------------
  func canConstruct(ransomNote string, magazine string) bool {
	arr := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		arr[magazine[i]-'a']++
	}
	for j := 0; j < len(ransomNote); j++ {
		arr[ransomNote[j]-'a']--
		if arr[ransomNote[j]-'a'] < 0 {
			return false
		}
	}
	return true
}
