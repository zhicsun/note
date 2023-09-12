# [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
```go
func lengthOfLongestSubstring(s string) int {
	// 返回值
	res := 0
	// 窗口
	window := make(map[byte]int, 0)
	for fast, slow := 0, 0; fast < len(s); fast++ {
		// 放进窗口并加一
		window[s[fast]]++

		// 窗口有重复值，慢指针移动到快指针的位置，清除窗口内的所有值
		for window[s[fast]] > 1 {
			window[s[slow]]--
			slow++
		}

		// 无重复值，计算长度并和现有结果比较，取最大值
		length := fast - slow + 1
		if length > res {
			res = length
		}
	}

	// 返回结果
	return res
}
```
# [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/)
```go
func minWindow(s string, t string) string {
    // 初始化返回结果
    res := ""

    // 格式化要查找的字符串
    need := make(map[byte]int, 0)
    for i:=0; i<len(t);i++{
        need[t[i]]++
    }

    // 初始化滑动窗口和查到的目标字符串计数值
    count := 0
    window := make(map[byte]int, 0)

    // 依次遍历所有字符串
    for slow, fast := 0, 0; fast < len(s); fast++ {
        // 窗口值累加
        window[s[fast]]++

        // 窗口值小于等于目标值时，计数目标字符串长度值才累加
        if window[s[fast]] <= need[s[fast]] {
            count++
        }

        // 除去左边无用的字符串，缩小窗口大小
        for slow <= fast && window[s[slow]] > need[s[slow]] {
            window[s[slow]]--
            slow++
        }

        // 窗口字符串包含目标字符串，计算最小字符串并赋值
        if count == len(t){
            if res == "" || fast-slow+1 < len(res) {
                res = s[slow:fast+1]
            }
        }
    }

    // 返回结果
    return res
}
```
# [438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/)
```go
func findAnagrams(s string, p string) []int {
    res := make([]int, 0)
    // 格式化需要查找的字符串
    need := make(map[byte]int, 0)
    for i:=0; i<len(p); i++ {
        need[p[i]]++
    }

    // 初始化窗口和找到的字符串
    count := 0
    window := make(map[byte]int, 0)

    // 依次遍历所有字符串
    for slow,fast := 0, 0; fast < len(s); fast++ {
        // 窗口值增加
        window[s[fast]]++

        // 窗口值小于等于目标值时，计数目标字符串长度值才累加
        if window[s[fast]] <= need[s[fast]] {
            count++
        }

        // 除去左边无用的字符串，缩小窗口大小
        for slow <= fast && window[s[slow]] > need[s[slow]] {
            window[s[slow]]--
            slow++
        }

        // 窗口包含目标字符串和大小等于目标字符串, 加入到结果集中
        if count == len(p) && fast-slow + 1 == len(p) {
            res = append(res, slow)
        }
    }

    // 返回结果
    return res
}
```
# [567. 字符串的排列](https://leetcode.cn/problems/permutation-in-string/)
```go
func checkInclusion(s1 string, s2 string) bool {
    // 格式化需要查找的字符串
    need := make(map[byte]int, 0)
    for i:=0; i<len(s1); i++ {
        need[s1[i]]++
    }

    // 初始化窗口和找到的字符串
    count := 0
    window := make(map[byte]int, 0)

    // 依次遍历所有字符串
    for slow,fast := 0, 0; fast < len(s2); fast++ {
        // 窗口值增加
        window[s2[fast]]++

        // 窗口值小于等于目标值时，计数目标字符串长度值才累加
        if window[s2[fast]] <= need[s2[fast]] {
            count++
        }

        // 除去左边无用的字符串，缩小窗口大小
        for slow <= fast && window[s2[slow]] > need[s2[slow]] {
            window[s2[slow]]--
            slow++
        }

        // 窗口包含目标字符串和大小等于目标字符串，返回结果
        if count == len(s1) && fast-slow + 1 == len(s1) {
            return true
        }
    }

    // 返回结果
    return false
}
```
