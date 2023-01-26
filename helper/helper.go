package helper

func GetArrayOfInterface[K comparable](v interface{}) []K {
	aInterface := v.([]interface{})
	aGen := make([]K, len(aInterface))
	for i, v := range aInterface {
		aGen[i] = v.(K)
	}
	return aGen
}
