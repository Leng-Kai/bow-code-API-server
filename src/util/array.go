package util

func Contain(array []interface{}, item interface{}) int {
	for i, ele := range array {
		if ele == item {
			return i
		}
	}
	return -1
}