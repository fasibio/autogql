package helper

import (
	"fmt"
	"strings"
)

func GetArrayOfInterface[K comparable](v interface{}) []K {
	aInterface := v.([]interface{})
	aGen := make([]K, len(aInterface))
	for i, v := range aInterface {
		aGen[i] = v.(K)
	}
	return aGen
}

func GetGormValue(gormDirectiveValue, key string) (string, error) {
	if strings.Contains(gormDirectiveValue, key+"=") {
		values := strings.Split(gormDirectiveValue, ";")
		for _, v := range values {
			if strings.Contains(v, key) {
				return strings.SplitN(v, "=", 1)[1], nil
			}
		}
	}
	return "", fmt.Errorf("not found %s", key)
}
