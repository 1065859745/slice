package slice

// IncludeStr 判断切片中是否存在某个元素，有一个bool返回值，存在返回true，不存在返回false。
func IncludeStr(arr []string, s string) bool {
	b := false
	for _, v := range arr {
		if v == s {
			b = true
			break
		}
	}
	return b
}

// IncludeSameStr 判断切片中是否存在相同的元素，有一个bool返回值，存在返回true，不存在返回false
func IncludeSameStr(arr []string) bool {
	for i := 0; i < len(arr)-1; i++ {
		for _, m := range arr[i+1:] {
			if arr[i] == m {
				return true
			}
		}
	}
	return false
}
