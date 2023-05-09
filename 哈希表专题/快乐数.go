力扣链接：https://leetcode.cn/problems/happy-number/
编写一个算法来判断一个数 `n` 是不是快乐数。

**「快乐数」** 定义为：

- 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
- 然后重复这个过程直到这个数变为 1，也可能是 **无限循环** 但始终变不到 1。
- 如果这个过程 **结果为** 1，那么这个数就是快乐数。

如果 `n` 是 *快乐数* 就返回 `true` ；不是，则返回 `false` 。
// 快乐数

// set：又排序又去重
// unorder_set: 不排序，去重

// 人是咋想的：给我一个数，我把各个位数拆开相加，取平方得到一个数，看看这个数是1不，不是的话，再拆。
// 关于拆开相加：循环求和，以前会
// 关于再拆：在死循环里，把结果数重新赋给n
// 关于是不是死循环的数(没想到，反而是本题的难点)：定义一个unorder_set，用来存储已经求和过的数

// 对于反复求和问题，用一个函数计算当前元素各位数的平方和
// 对于死循环问题，用unorder_set，存储每次运算的结果，如果是永远到不了1的，比如73，(注意个位数不停的平方，也有可能到1)，
// 在set里面找找这个数，它肯定会循环的
----------------------------- C++ --------------------------------------------
int sum(int m)
{
  int ret = 0;
  while(m){
​    ret += (m % 10) * (m % 10);
​    m /= 10;
  }
  return ret;
}
bool isHappy(int n) 
{
  unordered_set<int> set;
  int result;
  while(1)
  {
​    result = sum(n);
​    if(result == 1){
​      return true;
​    }
     // 在set中找到result了，即result是循环无果的数
​    if(set.find(result) != set.end()){
​      return false;
​    }
​    else{
​      set.insert(result);
​    }
​    // 就循环处理呗，之前没想到
​    n = result;
  }
}

----------------------------- Go --------------------------------------------
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
	}
	n /= 10
	return sum
}
func isHappy(n int) bool {
	arr := make(map[int]bool)
	// 刚开始的时候让arr数组里都为false
	// !arr[n]
	for n != 1 && !arr[n] {
		// 各位数之和如果是死循环，那么在将来的某个点，还会出现
		n := getSum(n)
		arr[n] = true
	}
	return n == 1
}
