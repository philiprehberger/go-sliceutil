# go-sliceutil

[![CI](https://github.com/philiprehberger/go-sliceutil/actions/workflows/ci.yml/badge.svg)](https://github.com/philiprehberger/go-sliceutil/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/philiprehberger/go-sliceutil.svg)](https://pkg.go.dev/github.com/philiprehberger/go-sliceutil)
[![License](https://img.shields.io/github/license/philiprehberger/go-sliceutil)](LICENSE)

Generic slice utilities for Go. Map, filter, reduce, and more with type safety.

## Installation

```sh
go get github.com/philiprehberger/go-sliceutil
```

## Usage

### Transform & Filter

```go
import "github.com/philiprehberger/go-sliceutil"

doubled := sliceutil.Map([]int{1, 2, 3}, func(n int) int {
    return n * 2
})
// [2, 4, 6]

evens := sliceutil.Filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
    return n%2 == 0
})
// [2, 4]
```

### Reduce

```go
sum := sliceutil.Reduce([]int{1, 2, 3, 4}, func(acc, n int) int {
    return acc + n
}, 0)
// 10
```

### Unique & Flatten

```go
unique := sliceutil.Unique([]int{1, 2, 2, 3, 3})
// [1, 2, 3]

flat := sliceutil.Flatten([][]int{{1, 2}, {3, 4}})
// [1, 2, 3, 4]
```

### Set Operations

```go
common := sliceutil.Intersect([]int{1, 2, 3}, []int{2, 3, 4})
// [2, 3]

diff := sliceutil.Difference([]int{1, 2, 3}, []int{2, 3, 4})
// [1]

all := sliceutil.Union([]int{1, 2, 3}, []int{2, 3, 4})
// [1, 2, 3, 4]
```

### Partition

```go
evens, odds := sliceutil.Partition([]int{1, 2, 3, 4, 5}, func(n int) bool {
    return n%2 == 0
})
// evens: [2, 4], odds: [1, 3, 5]
```

## API

| Function | Description |
|---|---|
| `Map[T, R]` | Transform each element |
| `Filter[T]` | Keep elements matching predicate |
| `Reduce[T, R]` | Fold into single value |
| `Unique[T]` | Remove duplicates, preserve order |
| `UniqueBy[T, K]` | Deduplicate by key function |
| `Flatten[T]` | Flatten one level of nesting |
| `FlatMap[T, R]` | Map then flatten |
| `Zip[T, U]` | Pair elements from two slices |
| `Partition[T]` | Split by predicate |
| `Chunk[T]` | Split into chunks of given size |
| `Reverse[T]` | Return reversed copy |
| `Shuffle[T]` | Return shuffled copy |
| `Contains[T]` | Check if element exists |
| `IndexOf[T]` | Return index or -1 |
| `First[T]` | Return first element |
| `Last[T]` | Return last element |
| `Intersect[T]` | Elements in both slices |
| `Difference[T]` | Elements in first but not second |
| `Union[T]` | All unique elements from both |
| `SymmetricDifference[T]` | Elements in either but not both |

## License

MIT
