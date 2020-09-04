package slice

// UpdateStr 插入新元素，若切片中已存在，则忽略.
func UpdateStr(arr []string, s string) []string {
	if arr != nil {
		for i, v := range arr {
			if i != len(arr)-1 {
				if v != s {
					continue
				}
				break
			}
			if v != s {
				return append(arr, s)
			}
		}
	}
	return append(arr, s)
}
