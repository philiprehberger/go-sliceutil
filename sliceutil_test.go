package sliceutil

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	result := Map(input, func(n int) string {
		return fmt.Sprintf("%d", n)
	})
	expected := []string{"1", "2", "3"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map(%v) = %v, want %v", input, result, expected)
	}
}

func TestMap_Empty(t *testing.T) {
	result := Map([]int{}, func(n int) int { return n * 2 })
	if len(result) != 0 {
		t.Errorf("Map on empty slice should return empty slice, got %v", result)
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	result := Filter(input, func(n int) bool { return n%2 == 0 })
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter(%v) = %v, want %v", input, result, expected)
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Reduce(input, func(acc int, n int) int { return acc + n }, 0)
	if result != 15 {
		t.Errorf("Reduce sum(%v) = %d, want 15", input, result)
	}
}

func TestUnique(t *testing.T) {
	input := []int{1, 2, 2, 3, 3}
	result := Unique(input)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unique(%v) = %v, want %v", input, result, expected)
	}
}

func TestUniqueBy(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}
	input := []person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 25},
	}
	result := UniqueBy(input, func(p person) int { return p.Age })
	if len(result) != 2 {
		t.Errorf("UniqueBy should return 2 elements, got %d", len(result))
	}
	if result[0].Name != "Alice" || result[1].Name != "Charlie" {
		t.Errorf("UniqueBy kept wrong elements: %v", result)
	}
}

func TestFlatten(t *testing.T) {
	input := [][]int{{1, 2}, {3, 4}}
	result := Flatten(input)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Flatten(%v) = %v, want %v", input, result, expected)
	}
}

func TestFlatMap(t *testing.T) {
	input := []int{1, 2, 3}
	result := FlatMap(input, func(n int) []int { return []int{n, n * 10} })
	expected := []int{1, 10, 2, 20, 3, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FlatMap(%v) = %v, want %v", input, result, expected)
	}
}

func TestZip(t *testing.T) {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}
	result := Zip(a, b)
	if len(result) != 3 {
		t.Fatalf("Zip length = %d, want 3", len(result))
	}
	if result[0].First != 1 || result[0].Second != "a" {
		t.Errorf("Zip[0] = %v, want {1, a}", result[0])
	}
	if result[2].First != 3 || result[2].Second != "c" {
		t.Errorf("Zip[2] = %v, want {3, c}", result[2])
	}
}

func TestZip_UnequalLength(t *testing.T) {
	a := []int{1, 2, 3}
	b := []string{"a", "b"}
	result := Zip(a, b)
	if len(result) != 2 {
		t.Errorf("Zip with unequal length = %d, want 2", len(result))
	}
}

func TestPartition(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	matched, unmatched := Partition(input, func(n int) bool { return n%2 == 0 })
	expectedMatched := []int{2, 4}
	expectedUnmatched := []int{1, 3, 5}
	if !reflect.DeepEqual(matched, expectedMatched) {
		t.Errorf("Partition matched = %v, want %v", matched, expectedMatched)
	}
	if !reflect.DeepEqual(unmatched, expectedUnmatched) {
		t.Errorf("Partition unmatched = %v, want %v", unmatched, expectedUnmatched)
	}
}

func TestChunk(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Chunk(input, 2)
	expected := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Chunk(%v, 2) = %v, want %v", input, result, expected)
	}
}

func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := Reverse(input)
	expected := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse(%v) = %v, want %v", input, result, expected)
	}
	// Ensure original is not modified.
	if !reflect.DeepEqual(input, []int{1, 2, 3, 4}) {
		t.Error("Reverse modified the original slice")
	}
}

func TestShuffle(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := Shuffle(input)
	if len(result) != len(input) {
		t.Errorf("Shuffle length = %d, want %d", len(result), len(input))
	}
	// Check same elements exist.
	sortedInput := Unique(input)
	sortedResult := Unique(result)
	if len(sortedInput) != len(sortedResult) {
		t.Error("Shuffle changed the elements")
	}
	// Ensure original is not modified.
	if !reflect.DeepEqual(input, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Error("Shuffle modified the original slice")
	}
}

func TestContains(t *testing.T) {
	input := []int{1, 2, 3}
	if !Contains(input, 2) {
		t.Error("Contains should find 2")
	}
	if Contains(input, 5) {
		t.Error("Contains should not find 5")
	}
}

func TestIndexOf(t *testing.T) {
	input := []string{"a", "b", "c"}
	if idx := IndexOf(input, "b"); idx != 1 {
		t.Errorf("IndexOf(b) = %d, want 1", idx)
	}
	if idx := IndexOf(input, "z"); idx != -1 {
		t.Errorf("IndexOf(z) = %d, want -1", idx)
	}
}

func TestFirst(t *testing.T) {
	v, ok := First([]int{10, 20, 30})
	if !ok || v != 10 {
		t.Errorf("First([10,20,30]) = (%d, %v), want (10, true)", v, ok)
	}
	_, ok = First([]int{})
	if ok {
		t.Error("First on empty slice should return false")
	}
}

func TestLast(t *testing.T) {
	v, ok := Last([]int{10, 20, 30})
	if !ok || v != 30 {
		t.Errorf("Last([10,20,30]) = (%d, %v), want (30, true)", v, ok)
	}
	_, ok = Last([]int{})
	if ok {
		t.Error("Last on empty slice should return false")
	}
}
