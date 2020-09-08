package slice

// UpdateStr 插入新元素，若切片中已存在，则忽略.
func UpdateStr(arr []string, s string) []string {
	if arr != nil {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] == s {
				return arr
			}
		}
	}
	return append(arr, s)
}
