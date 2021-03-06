package slice

// DelStr 删除切片中某一个元素.
func DelStr(arr []string, s ...string) []string {
	var ar []string
	for i, n := range arr {
		for _, m := range s {
			if n != m {
				ar = append(ar, arr[i])
			}
		}
	}
	return ar
}

// DelSameStr 删除string切片中相同的元素.
func DelSameStr(arr []string) []string {
	m := 1

a:
	for i := 0; i < len(arr)-1; i++ {
		for j := m; j < len(arr); j++ {
			if arr[i] == arr[j] {
				arr = append(arr[:j], arr[j+1:]...)
				i--
				continue a
			}
			m++
		}
		m = i + 2
	}
	return arr
}

// DelNearbyStr 删除切片中相邻相同的元素
func DelNearbyStr(arr []string) []string {
a:
	for {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] == arr[i+1] {
				arr = append(arr[:i], arr[i+1:]...)
				i--
				continue a
			}
		}
		break
	}
	return arr
}
