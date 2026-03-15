package sliceutil

// Intersect returns a new slice containing elements that appear in both a and b.
// The order follows the order of elements in a.
func Intersect[T comparable](a, b []T) []T {
	set := make(map[T]struct{}, len(b))
	for _, v := range b {
		set[v] = struct{}{}
	}
	result := make([]T, 0)
	for _, v := range a {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

// Difference returns a new slice containing elements that are in a but not in b.
// The order follows the order of elements in a.
func Difference[T comparable](a, b []T) []T {
	set := make(map[T]struct{}, len(b))
	for _, v := range b {
		set[v] = struct{}{}
	}
	result := make([]T, 0)
	for _, v := range a {
		if _, ok := set[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

// Union returns a new slice containing all unique elements from both a and b.
// Elements from a appear first (in order), followed by unique elements from b.
func Union[T comparable](a, b []T) []T {
	seen := make(map[T]struct{}, len(a)+len(b))
	result := make([]T, 0, len(a)+len(b))
	for _, v := range a {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	for _, v := range b {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// SymmetricDifference returns a new slice containing elements that are in either
// a or b, but not in both. Elements from a appear first, followed by elements from b.
func SymmetricDifference[T comparable](a, b []T) []T {
	setA := make(map[T]struct{}, len(a))
	for _, v := range a {
		setA[v] = struct{}{}
	}
	setB := make(map[T]struct{}, len(b))
	for _, v := range b {
		setB[v] = struct{}{}
	}
	result := make([]T, 0)
	for _, v := range a {
		if _, ok := setB[v]; !ok {
			result = append(result, v)
		}
	}
	for _, v := range b {
		if _, ok := setA[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}
