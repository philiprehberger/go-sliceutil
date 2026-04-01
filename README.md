# go-sliceutil

[![CI](https://github.com/philiprehberger/go-sliceutil/actions/workflows/ci.yml/badge.svg)](https://github.com/philiprehberger/go-sliceutil/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/philiprehberger/go-sliceutil.svg)](https://pkg.go.dev/github.com/philiprehberger/go-sliceutil)
[![Last updated](https://img.shields.io/github/last-commit/philiprehberger/go-sliceutil)](https://github.com/philiprehberger/go-sliceutil/commits/main)

Generic slice utilities for Go. Map, filter, reduce, and more with type safety

## Installation

```bash
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

### Find & Search

```go
val, ok := sliceutil.Find([]int{1, 2, 3, 4}, func(n int) bool {
    return n > 2
})
// val: 3, ok: true

idx := sliceutil.FindIndex([]string{"a", "b", "c"}, func(s string) bool {
    return s == "b"
})
// 1

hasEven := sliceutil.Any([]int{1, 3, 4}, func(n int) bool {
    return n%2 == 0
})
// true

allPositive := sliceutil.All([]int{1, 2, 3}, func(n int) bool {
    return n > 0
})
// true
```

### Sort, Take & Drop

```go
type Item struct {
    Name  string
    Price int
}

sorted := sliceutil.SortBy(items, func(i Item) int {
    return i.Price
})
// sorted by price ascending

first3 := sliceutil.Take([]int{1, 2, 3, 4, 5}, 3)
// [1, 2, 3]

rest := sliceutil.Drop([]int{1, 2, 3, 4, 5}, 2)
// [3, 4, 5]
```

### Compact

```go
cleaned := sliceutil.Compact([]string{"", "hello", "", "world"})
// ["hello", "world"]

nonZero := sliceutil.Compact([]int{0, 1, 0, 2, 3})
// [1, 2, 3]
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
| `Find[T]` | First element matching predicate |
| `FindIndex[T]` | Index of first match, -1 if none |
| `Any[T]` | True if any element matches |
| `All[T]` | True if all elements match |
| `SortBy[T, K]` | Sort by extracted key |
| `Take[T]` | Return first n elements |
| `Drop[T]` | Skip first n elements |
| `Compact[T]` | Remove zero-value elements |
| `Intersect[T]` | Elements in both slices |
| `Difference[T]` | Elements in first but not second |
| `Union[T]` | All unique elements from both |
| `SymmetricDifference[T]` | Elements in either but not both |

## Development

```bash
go test ./...
go vet ./...
```

## Support

If you find this project useful:

⭐ [Star the repo](https://github.com/philiprehberger/go-sliceutil)

🐛 [Report issues](https://github.com/philiprehberger/go-sliceutil/issues?q=is%3Aissue+is%3Aopen+label%3Abug)

💡 [Suggest features](https://github.com/philiprehberger/go-sliceutil/issues?q=is%3Aissue+is%3Aopen+label%3Aenhancement)

❤️ [Sponsor development](https://github.com/sponsors/philiprehberger)

🌐 [All Open Source Projects](https://philiprehberger.com/open-source-packages)

💻 [GitHub Profile](https://github.com/philiprehberger)

🔗 [LinkedIn Profile](https://www.linkedin.com/in/philiprehberger)

## License

[MIT](LICENSE)
