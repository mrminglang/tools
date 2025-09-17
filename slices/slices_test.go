package slices_test

import (
	"fmt"
	"testing"

	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/slices"
	"github.com/stretchr/testify/assert"
)

func TestIsSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4}
	res, ok := slices.IsSlice(intSlice)
	assert.True(t, ok)
	fmt.Println(res, ok)
}

func TestIsExistValue(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	ok1, err := slices.IsExistValue("b", a)
	assert.True(t, ok1)
	assert.Nil(t, err)

	b := map[string]string{
		"test":  "test",
		"test2": "test2",
	}
	ok2, err2 := slices.IsExistValue("test1", b)
	assert.False(t, ok2)
	assert.Error(t, err2)
}

func TestRemoveSlice(t *testing.T) {
	//intSlice := []int{1,2,3,4,5,6,7,8}
	//strSlice := []string{"a","b","c","d"}
	boolSlice := []bool{true, false, false, true}

	b, ok := slices.CreateAnyTypeSlice(boolSlice)

	if ok {
		c := slices.RemoveSlice(b, false)
		fmt.Println(c)
	}
}

func TestRemoveIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := slices.RemoveIntSlice(intSlice, 3)
	dumps.Dump(b)
}

func TestRemoveStrSliceDuplicates(t *testing.T) {
	str := []string{"apple", "banana", "apple", "kiwi", "banana", "orange"}

	dumps.Dump(slices.RemoveStrSliceDuplicates(str))
}

func TestStrSlice2InterfaceSlice(t *testing.T) {
	strArr := []string{"apple", "banana", "orange"}
	dumps.Dump(slices.StrSlice2InterfaceSlice(strArr))
}

func TestInArray(t *testing.T) {
	array := []string{"2258", "2259", "2260"}
	target := "102259"
	dumps.Dump(slices.InArray(array, target))
}

func TestGetSliceIntersection(t *testing.T) {
	lst1 := []string{"aaa", "bbb", "ccc", "ddd"}
	lst2 := []string{"bbb", "ddd", "ccc", "eee"}
	dumps.Dump(slices.GetSliceIntersection(lst1, lst2))
}
