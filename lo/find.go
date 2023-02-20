package lo

import (
	"fmt"

	"github.com/mudssky/leetcode-go/constraints"
)

// IndexOf returns the index at which the first occurrence of a value is found in an array or return -1
// if the value cannot be found.
func IndexOf[T comparable](collection []T, element T) int {
	for i, item := range collection {
		if item == element {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index at which the last occurrence of a value is found in an array or return -1
// if the value cannot be found.
func LastIndexOf[T comparable](collection []T, element T) int {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if collection[i] == element {
			return i
		}
	}

	return -1
}

// Find search an element in a slice based on a predicate. It returns element and true if element was found.
// 因为元素为空的情况,此时predicate判断为空,此时返回值都是空,没法判断有没有找到,所以加了一个bool的返回值区分
func Find[T any](collection []T, predicate func(item T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var result T
	return result, false
}

// FindIndexOf searches an element in a slice based on a predicate and returns the index and true.
// It returns -1 and false if the element is not found.
func FindIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	for i, item := range collection {
		if predicate(item) {
			return item, i, true
		}
	}

	var result T
	return result, -1, false
}

// FindLastIndexOf searches last element in a slice based on a predicate and returns the index and true.
// It returns -1 and false if the element is not found.
func FindLastIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], i, true
		}
	}

	var result T
	return result, -1, false
}

// Nth returns the element at index `nth` of collection. If `nth` is negative, the nth element
// from the end is returned. An error is returned when nth is out of slice bounds.
func Nth[T any, N constraints.Integer](collection []T, nth N) (T, error) {
	n := int(nth)
	l := len(collection)
	if n >= l || -n > l {
		var t T
		return t, fmt.Errorf("nth: %d out of slice bounds", n)
	}

	if n >= 0 {
		return collection[n], nil
	}
	return collection[l+n], nil
}
