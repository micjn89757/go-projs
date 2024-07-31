package common

// ChechInNumberSlice 检查slice中是否含有元素target
func CheckInNumberSlice[T comparable](target T, arr []T) bool {
	for _, val := range arr {
		if target == val {
			return true
		}
	}

	return false
}

// DelEleInSlice 删除arr中的target，只支持元素不重复，删除失败返回原arr
func DelEleInSlice[T comparable](target T, arr []T) (newArr []T) {
	for i, val := range arr {
		if target == val {
			newArr = append(newArr[:i], newArr[i+1:]...)
			return
		}
	}

	return arr
}