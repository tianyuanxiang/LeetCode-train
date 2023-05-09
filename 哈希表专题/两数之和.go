力扣链接：https://leetcode.cn/problems/two-sum/

**给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。**
**你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。**
**示例:**
**给定 nums = [2, 7, 11, 15], target = 9**
**因为 nums[0] + nums[1] = 2 + 7 = 9**
**所以返回 [0, 1]**

-----------------------C++---------------------------------
vector<int> twoSum(vector<int>& nums, int target) 
{
  // std:如果没有using namespace std，需要手动注明命名空间
  std:unordered_map<int, int> map;
  for(int i = 0;i < nums.size();i ++){
​    // find: 查找以 key 为键的键值对(对应本题就是元素值)，
​    // 如果找到，则返回一个指向该键值对的正向迭代器；反之，则返回一个指向容器中最后一个键值对之后位置的迭代器（如果 end() 方法返回的迭代器）。
​    auto iter = map.find(target - nums[i]);
​    // 返回一个指向容器中最后一个键值对之后位置的迭代器
​    // 意思就是找到了，即不等于end()
​    if(iter != map.end()){
​      // 返回两个下标
​      return {iter->second, i};
​    }
​    // 如果没找到,把元素值作为key,把下标作为value。为什么不把下标作为key?应该是find函数的事，查查
​    map.insert(pair<int, int>(nums[i], i));
  }
}

---------------------------- GO ------------------------------------
// 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，
// 并返回它们的数组下标。你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 思路：
// 注：Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
// 1、遍历数组，在map中查找目标值-数组元素，是否存在，如存在直接返回下标，否则加入map中()
func twoSum(nums []int, target int) []int {
	// 定义map
	arr := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 查找一下arr[*]这个键里是否有值，有的话，ok==true，value等于值
		// arr['*']括号里面是键；ok：bool;value:键所对应的值
		value, ok := arr[target-nums[i]]
		if ok {
			// value:原来里面已经有的值，所以要排在前面
			return []int{value, i}
			// 给键值对赋值
		} else {
			arr[nums[i]] = i
		}
	}
	return []int{}
}
