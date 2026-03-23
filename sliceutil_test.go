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

func TestFind(t *testing.T) {
	v, ok := Find([]int{1, 2, 3, 4, 5}, func(n int) bool { return n > 3 })
	if !ok || v != 4 {
		t.Errorf("Find(>3) = (%d, %v), want (4, true)", v, ok)
	}
}

func TestFind_NoMatch(t *testing.T) {
	_, ok := Find([]int{1, 2, 3}, func(n int) bool { return n > 10 })
	if ok {
		t.Error("Find with no match should return false")
	}
}

func TestFind_Empty(t *testing.T) {
	_, ok := Find([]int{}, func(n int) bool { return true })
	if ok {
		t.Error("Find on empty slice should return false")
	}
}

func TestFindIndex(t *testing.T) {
	idx := FindIndex([]string{"a", "b", "c"}, func(s string) bool { return s == "b" })
	if idx != 1 {
		t.Errorf("FindIndex(b) = %d, want 1", idx)
	}
}

func TestFindIndex_NoMatch(t *testing.T) {
	idx := FindIndex([]int{1, 2, 3}, func(n int) bool { return n > 10 })
	if idx != -1 {
		t.Errorf("FindIndex with no match = %d, want -1", idx)
	}
}

func TestFindIndex_Empty(t *testing.T) {
	idx := FindIndex([]int{}, func(n int) bool { return true })
	if idx != -1 {
		t.Errorf("FindIndex on empty slice = %d, want -1", idx)
	}
}

func TestAny(t *testing.T) {
	if !Any([]int{1, 2, 3}, func(n int) bool { return n == 2 }) {
		t.Error("Any should find 2")
	}
	if Any([]int{1, 2, 3}, func(n int) bool { return n == 5 }) {
		t.Error("Any should not find 5")
	}
}

func TestAny_Empty(t *testing.T) {
	if Any([]int{}, func(n int) bool { return true }) {
		t.Error("Any on empty slice should return false")
	}
}

func TestAll(t *testing.T) {
	if !All([]int{2, 4, 6}, func(n int) bool { return n%2 == 0 }) {
		t.Error("All should return true when all elements match")
	}
	if All([]int{2, 3, 6}, func(n int) bool { return n%2 == 0 }) {
		t.Error("All should return false when not all elements match")
	}
}

func TestAll_Empty(t *testing.T) {
	if !All([]int{}, func(n int) bool { return false }) {
		t.Error("All on empty slice should return true")
	}
}

func TestSortBy(t *testing.T) {
	type item struct {
		Name  string
		Value int
	}
	input := []item{
		{Name: "c", Value: 3},
		{Name: "a", Value: 1},
		{Name: "b", Value: 2},
	}
	result := SortBy(input, func(i item) int { return i.Value })
	if result[0].Name != "a" || result[1].Name != "b" || result[2].Name != "c" {
		t.Errorf("SortBy did not sort correctly: %v", result)
	}
	// Ensure original is not modified.
	if input[0].Name != "c" {
		t.Error("SortBy modified the original slice")
	}
}

func TestSortBy_Empty(t *testing.T) {
	result := SortBy([]int{}, func(n int) int { return n })
	if len(result) != 0 {
		t.Errorf("SortBy on empty slice should return empty, got %v", result)
	}
}

func TestTake(t *testing.T) {
	result := Take([]int{1, 2, 3, 4, 5}, 3)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Take(5, 3) = %v, want %v", result, expected)
	}
}

func TestTake_ExceedsLength(t *testing.T) {
	result := Take([]int{1, 2}, 5)
	expected := []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Take(2, 5) = %v, want %v", result, expected)
	}
}

func TestTake_Zero(t *testing.T) {
	result := Take([]int{1, 2, 3}, 0)
	if len(result) != 0 {
		t.Errorf("Take(3, 0) should return empty, got %v", result)
	}
}

func TestTake_Negative(t *testing.T) {
	result := Take([]int{1, 2, 3}, -1)
	if len(result) != 0 {
		t.Errorf("Take(3, -1) should return empty, got %v", result)
	}
}

func TestTake_Empty(t *testing.T) {
	result := Take([]int{}, 3)
	if len(result) != 0 {
		t.Errorf("Take on empty slice should return empty, got %v", result)
	}
}

func TestDrop(t *testing.T) {
	result := Drop([]int{1, 2, 3, 4, 5}, 2)
	expected := []int{3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Drop(5, 2) = %v, want %v", result, expected)
	}
}

func TestDrop_ExceedsLength(t *testing.T) {
	result := Drop([]int{1, 2}, 5)
	if len(result) != 0 {
		t.Errorf("Drop(2, 5) should return empty, got %v", result)
	}
}

func TestDrop_Zero(t *testing.T) {
	result := Drop([]int{1, 2, 3}, 0)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Drop(3, 0) = %v, want %v", result, expected)
	}
}

func TestDrop_Negative(t *testing.T) {
	result := Drop([]int{1, 2, 3}, -1)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Drop(3, -1) = %v, want %v", result, expected)
	}
}

func TestDrop_Empty(t *testing.T) {
	result := Drop([]int{}, 3)
	if len(result) != 0 {
		t.Errorf("Drop on empty slice should return empty, got %v", result)
	}
}

func TestCompact(t *testing.T) {
	result := Compact([]int{0, 1, 0, 2, 3, 0})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Compact = %v, want %v", result, expected)
	}
}

func TestCompact_Strings(t *testing.T) {
	result := Compact([]string{"", "a", "", "b"})
	expected := []string{"a", "b"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Compact strings = %v, want %v", result, expected)
	}
}

func TestCompact_NoZeros(t *testing.T) {
	result := Compact([]int{1, 2, 3})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Compact with no zeros = %v, want %v", result, expected)
	}
}

func TestCompact_Empty(t *testing.T) {
	result := Compact([]int{})
	if len(result) != 0 {
		t.Errorf("Compact on empty slice should return empty, got %v", result)
	}
}

func TestCompact_AllZeros(t *testing.T) {
	result := Compact([]int{0, 0, 0})
	if len(result) != 0 {
		t.Errorf("Compact on all-zero slice should return empty, got %v", result)
	}
}
