力扣链接：https://leetcode.cn/problems/intersection-of-two-arrays/
给定两个数组 `nums1` 和 `nums2` ，返回 *它们的交集* 。输出结果中的每个元素一定是 **唯一** 的。我们可以 **不考虑输出结果的顺序** 

------------------------------- C++ ---------------------------------------
vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
​    unordered_set<int> result; // 用来存放结果
​    unordered_set<int> nums_set(nums1.begin(), nums1.end());// 去重后的nums1
​    for(int num : nums2){
    	// 判断条件中为什么要有!=   ?
​      if(nums_set.find(num) != nums_set.end()){
​        result.insert(num);
​      }
​    }
​    return vector<int>(result.begin(), result.end());
}

------------------------------- go ---------------------------------------
func intersection(nums1 []int, nums2 []int) []int {
	// 用map模拟set(约定俗成的)set能去重哟！
	set := make(map[int]struct{}, 0)
	// 创建一个数组
	res := make([]int, 0)
	// v:[1,2,2,1]
	for _, v := range nums1 {
		// map可以通过“comma ok”机制来获取该key是否存在,_, ok := map["key"],如果没有对应的值,ok为false
		// set问ok，我[1]里(set[1])有数嘛？
		_, ok := set[v]
		// 如果没有对应的值
		if !ok {
			// struct{}{}  构造了一个struct {}类型的值
			// 刚开始的时候，set[1]里存的是struct{}{}这么个玩意？
			set[v] = struct{}{}
		}
	}

	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		// set问ok，我[2]里(set[2])有数嘛？
		_, ok := set[v]
		// 答案是“有”
		if ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
