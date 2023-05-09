力扣链接 ： https://leetcode.cn/problems/top-k-frequent-elements/
//	前 K 个高频元素
//
// 3步：1.统计元素出现的频率 2.对频率数进行排序 3.找出前k个高频元素
func topKFrequent(nums []int, k int) []int {
	res := []int{}
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}

	for re, _ := range map_num {
		res = append(res, re)
	}

	// 你光对map里的value排序，你怎么知道key是啥的？
	// 奥奥，map_num[res[出现的次数]]
	sort.Slice(res, func(a, b int) bool {
		return map_num[res[a]] > map_num[res[b]]
	})
	return res[:k]

}
