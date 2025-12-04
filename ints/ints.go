package ints

import (
	"fmt"
	"math"
	"strconv"
)

// InterfaceToInt32 converts an interface to an int32
func InterfaceToInt32(v interface{}) (int32, error) {
	switch val := v.(type) {
	case int32:
		return val, nil
	case int:
		if val > math.MaxInt32 || val < math.MinInt32 {
			return 0, fmt.Errorf("int value out of int32 range")
		}
		return int32(val), nil
	case int64:
		if val > math.MaxInt32 || val < math.MinInt32 {
			return 0, fmt.Errorf("int64 value out of int32 range")
		}
		return int32(val), nil
	case float64:
		if val > math.MaxInt32 || val < math.MinInt32 {
			return 0, fmt.Errorf("float64 value out of int32 range")
		}
		return int32(val), nil
	case string:
		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i), nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int32", v)
	}
}

// Int32Ptr returns a pointer to an int32
func Int32Ptr(v int32) *int32 {
	return &v
}
