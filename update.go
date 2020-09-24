package slice

// UpdateStr 插入新元素，若切片中已存在，则忽略.
func UpdateStr(arr []string, s ...string) []string {
	for i := 0; i < len(arr)-1; i++ {
		for _, v := range s {
			if arr[i] == v {
				return arr
			}
		}
	}
	return append(arr, s...)
}
