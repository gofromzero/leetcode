// Copyright 2021 The GoLand Authors.
// Author: Assistant
// Modified: 2025/1/21

package makefancystring

func makeFancyString(s string) string {
	res := []rune{} // 删除后的字符串
	// 遍历 s 模拟删除过程
	for _, ch := range s {
		n := len(res)
		if n >= 2 && res[n-1] == ch && res[n-2] == ch {
			// 如果 res 最后两个字符与当前字符均相等，则不添加
			continue
		}
		// 反之则添加
		res = append(res, ch)
	}
	return string(res)
}
