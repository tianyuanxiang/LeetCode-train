力扣链接：https://leetcode.cn/problems/4sum/
// 四数之和
// 思路
// 首先确定第一个和第二个元素，然后双指针算法确定第2、3、4个元素

// 确定第一个数时，不能再简单的nums[i]>target,因为targt有可能是负数，比如-4,-1,0,0;target=-5,你-4>-5直接break，把正确的过滤了

// 需要怎么办呢？当nums[i] > target && (target>=0 || nums[i]>=0)。对于target>=0,比如target=5,nums[i]=6

// 第一个和第二个元素的确定具体来说， 两层for循环，先固定i，找j和left和right里面满足的

// 第3、4个元素和<三数之和>一样

----------------------------------  C++  --------------------------------------------
vector<vector<int>> fourSum(vector<int>& nums, int target) 
{
  vector<vector<int>> result;
  sort(nums.begin(), nums.end());
  for(int i = 0;i < nums.size();i ++){
​    // 一级剪枝
​    if(nums[i] > target && nums[i] >= 0){
​        break;
​    }
​    if(nums[i] == nums[i-1]){
​      continue;
​    }

​    // 二级剪枝
​    for(int j = i + 1;j < nums.size();j ++){

​      if(nums[j] + nums[i] > target && nums[j] + nums[i] >= 0){
​        break;
​      }
​      if(j > i+1 && nums[j] == nums[j - 1]){
​        continue;
​      }
​      int left = j + 1, right = nums.size() - 1;
​      while(left < right)
​      {
​        if(nums[i] + nums[j] + nums[left] + nums[right] > target){
​          right --;
​        }
​        else if(nums[i] + nums[j] + nums[left] + nums[right] < target){
​          left ++;
​        }
​        else{
​          result.push_back(vector<int>{nums[i], nums[j], nums[left], nums[right]});
​          // 找到了一个结果之后，怎么办？对left 和right 去重，防止加上了重复的b或c
​          while(left < right && nums[left] == nums[left + 1]){
​            left ++;
​          }
​          while(left < right && nums[right] == nums[right - 1]){
​            right --;
​          }
​          // 找到答案后
​          left ++;
​          right --;
​        }
​      }
​    }
  }
  return result;
}

----------------------------------  gO  --------------------------------------------
// // 四数之和
// 整体来说和三数之和差不多，区别的地方在于，四数之和要先确定两个数，需要两层循环
// 由于给定的target不同(负数)，不能直接判断nums[0]>=0,需要nums[0]>target&&nums[0]>=0
// 还需要顾及2级剪枝处理，即确定的第二个数(nums[k]+nums[i]这俩是一个整体),if nums[k]+nums[i]>target && nums[i]+nums[k]>=0
// 把握两个点：1、数组是排序后的 2、2级剪枝的时候，把nums[i]和nums[k]看成一个整体

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 && nums[i] > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] > target && nums[i]+nums[k] >= 0 {
				break
			}
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}
			left := k + 1
			right := len(nums) - 1
			for left < right {
				if nums[i]+nums[k]+nums[left]+nums[right] < target {
					left++
				} else if nums[i]+nums[k]+nums[left]+nums[right] > target {
					right--
				} else {
					res = append(res, []int{nums[i], nums[k], nums[left], nums[right]})
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
	}
	return res
}
