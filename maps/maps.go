package maps

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

/*
以map的key(int\float\string)排序
sortType = 0 正序
sortType = 1 倒序
eachMap      ->  待遍历的map
eachFunc     ->  map遍历接收，入参应该符合map的key和value
*/

// SortMapKey 排序
func SortMapKey(eachMap interface{}, eachFunc interface{}, sortType int) {
	eachMapValue := reflect.ValueOf(eachMap)
	eachFuncValue := reflect.ValueOf(eachFunc)
	eachMapType := eachMapValue.Type()
	eachFuncType := eachFuncValue.Type()
	if eachMapValue.Kind() != reflect.Map {
		panic(errors.New("ksort.eachMap Kind failed"))
	}
	if eachFuncValue.Kind() != reflect.Func {
		panic(errors.New("ksort.eachFunc Kind failed"))
	}
	if eachFuncType.NumIn() != 2 {
		panic(errors.New("ksort.eachFunc NumIn failed"))
	}
	if eachFuncType.In(0).Kind() != eachMapType.Key().Kind() {
		panic(errors.New("ksort.eachFunc input parameter 1 type not equal of eachMap key"))
	}
	if eachFuncType.In(1).Kind() != eachMapType.Elem().Kind() {
		panic(errors.New("ksort.eachFunc input parameter 2 type not equal of eachMap value"))
	}

	// 对key进行排序
	// 获取排序后map的key和value，作为参数调用eachFunc即可
	switch eachMapType.Key().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		keys := make([]int, 0)
		keysMap := map[int]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, int(value.Int()))
			keysMap[int(value.Int())] = value
		}
		if sortType == 1 {
			sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		} else {
			sort.Ints(keys)
		}

		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.Float64, reflect.Float32:
		keys := make([]float64, 0)
		keysMap := map[float64]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, float64(value.Float()))
			keysMap[float64(value.Float())] = value
		}
		if sortType == 1 {
			sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		} else {
			sort.Float64s(keys)
		}
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	case reflect.String:
		keys := make([]string, 0)
		keysMap := map[string]reflect.Value{}
		for _, value := range eachMapValue.MapKeys() {
			keys = append(keys, value.String())
			keysMap[value.String()] = value
		}
		if sortType == 1 {
			sort.Sort(sort.Reverse(sort.StringSlice(keys)))
		} else {
			sort.Strings(keys)
		}
		for _, key := range keys {
			eachFuncValue.Call([]reflect.Value{keysMap[key], eachMapValue.MapIndex(keysMap[key])})
		}
	default:
		panic(errors.New("eachMap key type must is int or float or string"))
	}

	return
}

// ConvertToMapStringSlice interface 转 map[string][]string
func ConvertToMapStringSlice(param interface{}) (map[string][]string, error) {
	res := make(map[string][]string)
	m, ok := param.(map[string]interface{})
	if !ok {
		return res, errors.New("参数格式错误")
	}
	for k, v := range m {
		switch val := v.(type) {
		case []string:
			res[k] = val
		case []interface{}:
			for _, s := range val {
				if str, ok := s.(string); ok {
					res[k] = append(res[k], str)
				}
			}
		case map[string]interface{}:
			subMap, err := ConvertToMapStringSlice(val)
			if err != nil {
				return res, err
			}
			if subValue, exists := subMap["value"]; exists {
				res[k] = subValue
			}
		default:
			return res, errors.New(fmt.Sprintf("参数格式错误，键：%s，值：%v", k, v))
		}
	}
	return res, nil
}

// ConvertToMapSliceString 将 interface{} 类型的模板参数转换为 []map[string]string 类型
func ConvertToMapSliceString(param interface{}) (result []map[string]string, err error) {
	if param == nil {
		err = errors.New("参数不能为空")
		return
	}

	switch p := param.(type) {
	case map[string]string:
		// 处理单个 map 对象
		strMap := make(map[string]string)
		for k, v := range p {
			strMap[k] = fmt.Sprintf("%v", v)
		}
		result = append(result, strMap)
	case map[string]interface{}:
		// 处理单个 map 对象
		strMap := make(map[string]string)
		for k, v := range p {
			strMap[k] = fmt.Sprintf("%v", v)
		}
		result = append(result, strMap)
	case []map[string]interface{}:
		// 处理 map 数组
		for _, item := range p {
			strMap := make(map[string]string)
			for k, v := range item {
				strMap[k] = fmt.Sprintf("%v", v)
			}
			result = append(result, strMap)
		}
	case []interface{}:
		// 处理 interface{} 数组
		for _, item := range p {
			if mapItem, ok := item.(map[string]interface{}); ok {
				strMap := make(map[string]string)
				for k, v := range mapItem {
					strMap[k] = fmt.Sprintf("%v", v)
				}
				result = append(result, strMap)
			}
		}
	default:
		// 不支持的类型，返回空结果
		err = errors.New(fmt.Sprintf("不支持的参数类型"))
		return
	}

	return
}
