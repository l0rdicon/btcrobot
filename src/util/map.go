package util

//TODO, proper comments
func MapKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for key, _ := range data {
		keys = append(keys, key)
	}
	return keys
}

func MapIntKeys(data map[int]int) []int {
	keys := make([]int, 0, len(data))
	for key, _ := range data {
		keys = append(keys, key)
	}
	return keys
}
