package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func main() {}

func MapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func Filter[T any](s []T, f func(T) bool) []T {
	var ret []T
	for _, v := range s {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func ToCommaSeparatedString(intSlice []int) string {
	stringSlice := make([]string, len(intSlice))
	for index, i := range intSlice {
		stringSlice[index] = strconv.Itoa(i)
	}
	return strings.Join(stringSlice, ",")
}

func Toi[T constraints.Integer | constraints.Float](nn []T) []int {
	if nn == nil {
		return nil
	}
	is := make([]int, len(nn))
	for i, n := range nn {
		is[i] = int(n)
	}
	return is
}

func Toi8[T constraints.Integer | constraints.Float](nn []T) []int8 {
	if nn == nil {
		return nil
	}
	is := make([]int8, len(nn))
	for i, n := range nn {
		is[i] = int8(n)
	}
	return is
}

func Toi32[T constraints.Integer | constraints.Float](nn []T) []int32 {
	if nn == nil {
		return nil
	}
	is := make([]int32, len(nn))
	for i, n := range nn {
		is[i] = int32(n)
	}
	return is
}

func Toi64[T constraints.Integer | constraints.Float](nn []T) []int64 {
	if nn == nil {
		return nil
	}
	is := make([]int64, len(nn))
	for i, n := range nn {
		is[i] = int64(n)
	}
	return is
}

func I32ArrayContains(s []int32, e []int32) bool {
	var matchCount int
	for _, ss := range s {
		for _, ee := range e {
			if ee == ss {
				matchCount++
				continue
			}
		}
	}
	return len(e) == matchCount
}

func I8sToStrings(ii []int8) []string {
	if ii == nil {
		return nil
	}
	ss := make([]string, len(ii))
	for idx, i := range ii {
		ss[idx] = fmt.Sprint(i)
	}
	return ss
}

func IsToStrings(ii []int) []string {
	if ii == nil {
		return nil
	}
	ss := make([]string, len(ii))
	for idx, i := range ii {
		ss[idx] = fmt.Sprint(i)
	}
	return ss
}

func StringContainsSubStrs(substrs []string, s string) bool {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

func Int8Sort(i []int8) []int8 {
	if i == nil {
		return nil
	}
	int8s := make([]int8, 0, len(i))
	int8s = append(int8s, i...)
	sort.Sort(Int8Slice(int8s))
	return int8s
}

func StrToi(ss []string) ([]int, error) {
	result := make([]int, len(ss))
	for i, s := range ss {
		intValue, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result[i] = intValue
	}
	return result, nil
}

func StrToi32(ss []string) ([]int32, error) {
	result := make([]int32, len(ss))
	for i, s := range ss {
		intValue, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result[i] = int32(intValue)
	}
	return result, nil
}

type Int8Slice []int8

func (p Int8Slice) Len() int           { return len(p) }
func (p Int8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TrimSpaceFromStrings(s []string) []string {
	if s == nil {
		return nil
	}

	ss := make([]string, len(s))
	for index, v := range s {
		ss[index] = strings.TrimSpace(v)
	}

	return ss
}
