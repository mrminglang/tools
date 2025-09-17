package slices

import (
	"errors"
	"reflect"
	"sort"
)

// IsSlice 判断是否为slice数据
func IsSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == reflect.Slice {
		ok = true
	}

	return
}

// IsExistValue 判断切片的值是否存在 或 map类型的key是否存在
func IsExistValue(value interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)

	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == value {
				return true, nil
			}
		}
	case reflect.Map:
		// isExist map key
		if targetValue.MapIndex(reflect.ValueOf(value)).IsValid() {
			return true, nil
		}

		// todo isExist map value
	default:
		break
	}

	return false, errors.New("not in array")
}

// CreateAnyTypeSlice interface{}转为 []interface{}
func CreateAnyTypeSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := IsSlice(slice)

	if !ok {
		return nil, false
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, true
}

// RemoveSlice 删除切片的某个元素
func RemoveSlice(slice []interface{}, elem interface{}) []interface{} {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return RemoveSlice(slice, elem)
		}
	}
	return slice
}

// RemoveIntSlice 删除切片为INT类型的某个元素
func RemoveIntSlice(slice []int, elem interface{}) []int {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return RemoveIntSlice(slice, elem)
		}
	}
	return slice
}

// SliceToSqlString 字符串类型切片转SQL语句类型字符串
func SliceToSqlString(slice []string) string {
	var str string
	for _, v := range slice {
		if str == "" {
			str = "'" + v + "'"
		} else {
			str = str + "," + "'" + v + "'"
		}
	}
	return str
}

// RemoveStrSliceDuplicates 移除切片重复元素
func RemoveStrSliceDuplicates(slice []string) []string {
	encountered := map[string]struct{}{}
	var result []string

	for _, v := range slice {
		if _, has := encountered[v]; !has {
			encountered[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// StrSlice2InterfaceSlice []string 转 []interface
func StrSlice2InterfaceSlice(strArr []string) []interface{} {
	interfaceArr := make([]interface{}, len(strArr))

	for i, v := range strArr {
		interfaceArr[i] = interface{}(v)
	}

	return interfaceArr
}

// InArray 判断字符串类型数组的某个元素值是否存在(二分法)
func InArray(array []string, target string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	// index的值:[0,len(array)]
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}

// GetSliceIntersection 求两个字符串切片的交集
func GetSliceIntersection(lst1, lst2 []string) []string {
	var smaller, larger []string
	if len(lst1) < len(lst2) {
		smaller, larger = lst1, lst2
	} else {
		smaller, larger = lst2, lst1
	}

	// 较小的切片构建 map 以节省空间
	set := make(map[string]bool, len(smaller))
	for _, v := range smaller {
		set[v] = true
	}

	// 遍历较大切片查找交集
	res := make([]string, 0)
	for _, v := range larger {
		if set[v] {
			res = append(res, v)
			delete(set, v) // 避免重复添加
		}
	}
	return res
}
