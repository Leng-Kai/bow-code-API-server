package util

func Contain(array []interface{}, item interface{}) bool {
	for i, ele := range array {
		if ele == item {
			return true
		}
	}
	return false
}