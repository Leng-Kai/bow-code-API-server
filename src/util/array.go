package util

import (
	"github.com/Leng-Kai/bow-code-API-server/schemas"
)

func Contain_str(array []string, item string) bool {
	for _, ele := range array {
		if ele == item {
			return true
		}
	}
	return false
}

func Contain_ID(array []schemas.ID, item schemas.ID) bool {
	for _, ele := range array {
		if ele == item {
			return true
		}
	}
	return false
}