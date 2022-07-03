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
