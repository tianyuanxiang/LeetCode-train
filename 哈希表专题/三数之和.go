力扣链接：https://leetcode.cn/problems/3sum/
给你一个整数数组 `nums` ，判断是否存在三元组 `[nums[i], nums[j], nums[k]]` 满足 `i != j`、`i != k` 且 `j != k` ，同时还满足 `nums[i] + nums[j] + nums[k] == 0` 。请

你返回所有和为 `0` 且不重复的三元组。

**注意：**答案中不可以包含重复的三元组。

---------------------------------- C++ --------------------------------------------------------
// 三数之和
// 三个数的位置下标不能重复，三个数可以重复
// 绕人: 去重不好理解，

vector<vector<int>> threeSum(vector<int>& nums) 
{
  vector<vector<int>> result;
  sort(nums.begin(), nums.end());
  // 循环遍历数组，然后设置双指针，left指针指向k之后的第一个数，right指向最后一个数，
  // 如果nums[k] + nums[left] + nums[right] > 0,right左移，<0的话，left右移
  for(int k = 0;k < nums.size();k ++){
​    // 排序完事后第一个数还大于0，直接完事，这个数组不符合。
​    if(nums[k] > 0) break;
​    // i>0是防止越界; nums[i]==nums[i-1]是看看前面是不是已经出现过i了，如果前面出现过i，直接找下一个；如果没出现，那以i为第一个数，收集left和right了
​    if(i > 0 && nums[i] == nums[i - 1]){
​      continue;	
​    }
​    // 从这才开始处理b和c，前面都是准备工作
​    int left = k + 1, right = nums.size() - 1;
​    while(left < right)
​    {
​      if(nums[k] + nums[left] + nums[right] > 0){
​        right --;
​      }
​      else if(nums[k] + nums[left] + nums[right] < 0){
​        left ++;
​      }
​      else{
​        result.push_back(vector<int>{nums[k], nums[left], nums[right]});
​        // 找到了一个结果之后，怎么办？对left 和right 去重，防止加上了重复的b或c
​        while(right > left && nums[left] == nums[left + 1]){
​          left ++;
​        }
​        while(right > left && nums[right] == nums[right - 1]){
​          right --;
​        }
​        // 找到答案后
​        left ++;
​        right --
​      }
​    }
  }
}

------------------------------------ GO ------------------------------------------------
// 三数之和
// a,b,c三数求和要牵扯到三次去重：
// 1.a去重(只是a，b和c有自己的去重策略)
// arr[i]!=arr[i-1],为啥不是arr[i]==arr[i+1]?
// 对于结果集来说，arr[i]==arr[i+1]是判断结果集里是否有重复的元素。这样会漏掉答案,
// arr[i]!=arr[i-1]是从未来推现在,先往结果集里放元素
// 2.b,c 去重
// 由于已经排好序了，在找到合适的b,c之后,只需判断前n个有没有相同元素即可
// 先对元素排序，

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// 嵌套数组
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}
