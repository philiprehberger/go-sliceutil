// Package sliceutil provides generic slice utilities for Go.
//
// It includes common functional operations like Map, Filter, Reduce,
// as well as utilities for deduplication, chunking, partitioning, and more.
// All functions are generic and work with any compatible types.
package sliceutil

import "math/rand/v2"

// Pair holds two values of potentially different types.
type Pair[T any, U any] struct {
	First  T
	Second U
}

// Map transforms each element of the slice using the given function.
func Map[T any, R any](s []T, fn func(T) R) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate.
func Filter[T any](s []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce folds a slice into a single value by applying the function to an accumulator
// and each element, starting with the given initial value.
func Reduce[T any, R any](s []T, fn func(R, T) R, initial R) R {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

// Unique returns a new slice with duplicate elements removed, preserving order.
func Unique[T comparable](s []T) []T {
	seen := make(map[T]struct{}, len(s))
	result := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// UniqueBy returns a new slice with elements deduplicated by a key function, preserving order.
// The first element for each key is kept.
func UniqueBy[T any, K comparable](s []T, fn func(T) K) []T {
	seen := make(map[K]struct{}, len(s))
	result := make([]T, 0, len(s))
	for _, v := range s {
		key := fn(v)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Flatten flattens one level of nesting, converting a slice of slices into a single slice.
func Flatten[T any](s [][]T) []T {
	total := 0
	for _, inner := range s {
		total += len(inner)
	}
	result := make([]T, 0, total)
	for _, inner := range s {
		result = append(result, inner...)
	}
	return result
}

// FlatMap applies a function that returns a slice to each element, then flattens the result.
func FlatMap[T any, R any](s []T, fn func(T) []R) []R {
	result := make([]R, 0)
	for _, v := range s {
		result = append(result, fn(v)...)
	}
	return result
}

// Zip pairs elements from two slices. The result has the length of the shorter slice.
func Zip[T any, U any](a []T, b []U) []Pair[T, U] {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	result := make([]Pair[T, U], n)
	for i := 0; i < n; i++ {
		result[i] = Pair[T, U]{First: a[i], Second: b[i]}
	}
	return result
}

// Partition splits a slice into two groups based on a predicate.
// The first group contains elements that satisfy the predicate,
// the second contains those that do not.
func Partition[T any](s []T, predicate func(T) bool) (matched []T, unmatched []T) {
	matched = make([]T, 0)
	unmatched = make([]T, 0)
	for _, v := range s {
		if predicate(v) {
			matched = append(matched, v)
		} else {
			unmatched = append(unmatched, v)
		}
	}
	return matched, unmatched
}

// Chunk splits a slice into chunks of the given size.
// The last chunk may have fewer elements than size.
// If size is less than 1, it panics.
func Chunk[T any](s []T, size int) [][]T {
	if size < 1 {
		panic("sliceutil: chunk size must be at least 1")
	}
	chunks := make([][]T, 0, (len(s)+size-1)/size)
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

// Reverse returns a new slice with elements in reverse order.
func Reverse[T any](s []T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

// Shuffle returns a new slice with elements in random order.
func Shuffle[T any](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

// Contains reports whether the slice contains the given element.
func Contains[T comparable](s []T, elem T) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of elem in the slice,
// or -1 if the element is not found.
func IndexOf[T comparable](s []T, elem T) int {
	for i, v := range s {
		if v == elem {
			return i
		}
	}
	return -1
}

// Last returns the last element of the slice and true,
// or the zero value and false if the slice is empty.
func Last[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[len(s)-1], true
}

// First returns the first element of the slice and true,
// or the zero value and false if the slice is empty.
func First[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}
