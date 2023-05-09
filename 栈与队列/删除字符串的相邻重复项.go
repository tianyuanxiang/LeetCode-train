力扣链接：https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
// 删除字符串中的所有相邻重复项
// 思路：
// 1.创建一个栈，遍历原字符串依次和栈顶st.top()比较，如果元素相同，i++，st.pop()
// 2.终止条件：
func removeDuplicates(s string) string {
	st := make([]byte, 0)
	i := 0
	for i < len(s) {
		if len(st) != 0 {
			// 和栈顶元素比较
			if s[i] == st[len(st)-1] {
				// 弹出栈顶元素
				st = st[:len(st)-1]
				i++
			} else {
				st = append(st, s[i])
				i++
			}
		} else {
			st = append(st, s[i])
			i++
		}
	}
	return string(st)
}
