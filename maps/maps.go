package maps

import (
	"sort"
)

// SortKeyRise 升序
func SortKeyRise(result map[string]interface{}) map[string]interface{} {
	keys := make([]string, 0)
	for key := range result {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	rsp := make(map[string]interface{}, 0)
	for _, k := range keys {
		rsp[k] = result[k]
	}

	return rsp
}

// SortKeyDrop 升序
func SortKeyDrop(result map[string]interface{}) map[string]interface{} {
	keys := make([]string, 0)
	for key := range result {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	rsp := make(map[string]interface{}, 0)
	for _, k := range keys {
		rsp[k] = result[k]
	}

	return rsp
}
