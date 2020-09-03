package slice

// 删除切片中某一个元素.
<<<<<<< HEAD
func delStr(arr []string, s string) []string {
=======
func del(arr []string, s string) []string {
>>>>>>> 71dffa31e2f6ccf32da098534d56068384bf50f6
	var ar []string
	for i, v := range arr {
		if v == s {
			continue
		}
		ar = append(ar, arr[i])
	}
	return ar
}

// 删除string切片中相同的元素.
func delStrSame(arr []string) []string {
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

// 插入新元素，若切片中已存在，则忽略.
<<<<<<< HEAD
func updateStr(arr []string, s string) []string {
=======
func update(arr []string, s string) []string {
>>>>>>> 71dffa31e2f6ccf32da098534d56068384bf50f6
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

//  判断切片中是否存在某个元素，有一个bool返回值，存在返回true，否则返回false。
<<<<<<< HEAD
func includesStr(arr []string, s string) bool {
=======
func includes(arr []string, s string) bool {
>>>>>>> 71dffa31e2f6ccf32da098534d56068384bf50f6
	b := false
	for _, v := range arr {
		if v == s {
			b = true
			break
		}
	}
	return b
}

// 删除切片中相邻相同的元素
<<<<<<< HEAD
func delNearbyStr(arr []string) []string {
=======
func delNearby(arr []string) []string {
>>>>>>> 71dffa31e2f6ccf32da098534d56068384bf50f6
a:
	for {
		for i, v := range arr {
			if i != len(arr)-1 {
				if v == arr[i+1] {
					arr = append(arr[:i], arr[i+1:]...)
					continue a
				}
			}
		}
		break
	}
	return arr
}
