力扣链接：https://leetcode.cn/problems/4sum-ii/
给你四个整数数组 `nums1`、`nums2`、`nums3` 和 `nums4` ，数组长度都是 `n` ，请你计算有多少个元组 `(i, j, k, l)` 能满足：

- `0 <= i, j, k, l < n`
- `nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0`
// 四数相加
// 和四数之和差不多，都是两个两个的找
// 先用两个循环，把nums1和nums2的各个元素相加，存在一个map里，并且元素之和作为下标，对应元素持续相加(不用去重)
// 然后遍历nums3和nums4,当0-（b+c）在上述map中，计数器+=1即可

int fourSumCount(vector<int>& nums1, vector<int>& nums2, vector<int>& nums3, vector<int>& nums4) 
{  
  unordered_map<int, int> umap;
  int cnt = 0;
  for(int i : nums1){
​    for(int j : nums2){
​      // 注意，是求和后+1。你总不能单个元素+1吧，那怎么求是不是满足！
​      umap[i + j] += 1;
​    }
  }
  for(int c : nums3){
​    for(int d : nums4){
​      // 找到数了
​      if(umap.find(0 - (c + d)) != umap.end();)
​        // 难点：要直接加完该元素出现的次数，反正不用去重
​        cnt += umap[0 - (c + d)];
​    }
  }
  return cnt;
}
