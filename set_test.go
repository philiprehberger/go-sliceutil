package sliceutil

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	result := Intersect(a, b)
	expected := []int{2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Intersect(%v, %v) = %v, want %v", a, b, result, expected)
	}
}

func TestIntersect_Empty(t *testing.T) {
	result := Intersect([]int{}, []int{1, 2, 3})
	if len(result) != 0 {
		t.Errorf("Intersect with empty slice should return empty, got %v", result)
	}
	result = Intersect([]int{1, 2, 3}, []int{})
	if len(result) != 0 {
		t.Errorf("Intersect with empty slice should return empty, got %v", result)
	}
}

func TestDifference(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	result := Difference(a, b)
	expected := []int{1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Difference(%v, %v) = %v, want %v", a, b, result, expected)
	}
}

func TestUnion(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	result := Union(a, b)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Union(%v, %v) = %v, want %v", a, b, result, expected)
	}
}

func TestSymmetricDifference(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	result := SymmetricDifference(a, b)
	expected := []int{1, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SymmetricDifference(%v, %v) = %v, want %v", a, b, result, expected)
	}
}
