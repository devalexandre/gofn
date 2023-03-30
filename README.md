# gofn
![Coverage](https://img.shields.io/badge/Coverage-84.4%25-brightgreen)
[![Go](https://github.com/devalexandre/gofn/actions/workflows/go.yml/badge.svg)](https://github.com/devalexandre/gofn/actions/workflows/go.yml)

# library to use golang functional

# Summary

- [Array](#array)
    - [Usage](#usage)
        - [array.Filter](#arrayfilter)
        - [array.Find](#arrayfind)
        - [array.Map](#arraymap)
        - [array.FlatMap](#arrayflatmap)
        - [array.Reduce](#arrayreduce)
        - [array.Any](#arrayany)
        - [array.Some](#arraysome)
        - [array.Every](#arrayevery)
        - [array.Sum](#arraysum)
        - [array.Max](#arraymax)
        - [array.Min](#arraymin)
        - [array.Product](#arrayproduct)
        - [array.Reverse](#arrayreverse)
        - [array.Shuffle](#arrayshuffle)
        - [array.Unique](#arrayunique)
        - [array.Union](#arrayunion)
        - [array.Fill](#arrayfill)
        - [array.Join](#arrayjoin)
        - [array.Pop](#arraypop)
        - [array.Push](#arraypush)
        - [array.Shift](#arrayshift)
        - [array.Sort](#arraysort)
    - [chaining functions](#chaining-functions)


## Usage

### array.Filter

Filter a array , with a function that return a boolean.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Filter(data, func(i int) bool {
        return i%2 == 0
    })

    fmt.Println(result) // [2 4 6 8 10]
	
```

### array.Find

Find a element in the array , with a function that return a boolean.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Find(data, func(i int) bool {
        return i%2 == 0
    })

    fmt.Println(result) // 2

```


### array.Map

Map a array , with a function that return a value.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Map(data, func(i int) int {
        return i * 2
    })

    fmt.Println(result) // [2 4 6 8 10 12 14 16 18 20]

```

### array.FlatMap

FlatMap a array , with a function that return a slice.
```go

a := []int{1, 2, 3, 4, 5}
b := array.FlatMap(a, func(x int) []int { return []int{x, x * 2} })

fmt.Println(b) // [1 2 2 4 3 6 4 8 5 10]

```

### array.Reduce

Reduce a array , with a function that return a value.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Reduce(data, func(i, j int) int {
        return i + j
    })

    fmt.Println(result) // 55

```

### array.Any

Check if any element in the array match the condition.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Any(data, func(i int) bool {
        return i%2 == 0
    })

    fmt.Println(result) // true

```

### array.Some

Check if some element in the array match the condition.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Some(data, func(i int) bool {
        return i%2 == 0
    })

    fmt.Println(result) // true

```

### array.Every

Check if every element in the array match the condition.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Every(data, func(i int) bool {
        return i%2 == 0
    })

    fmt.Println(result) // false

```

### array.Sum

Sum all elements in the array.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Sum(data)

    fmt.Println(result) // 55

```

### array.Max

Get the max element in the array.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Max(data)

    fmt.Println(result) // 10

```

### array.Min

Get the min element in the array.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Min(data)

    fmt.Println(result) // 1

```

### array.Product

Get the product of all elements in the array.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Product(data)

    fmt.Println(result) // 3628800

```

### array.Reverse

Reverse te order of the array.

```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Reverse(data)

    fmt.Println(result) // [10 9 8 7 6 5 4 3 2 1]

```

### array.Shuffle

Shuffle the array.

```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Shuffle(data)

    fmt.Println(result) // [10 9 8 7 6 5 4 3 2 1]

```

### array.Unique

Get unique elements in the array.

```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Unique(data)

    fmt.Println(result) // [1 2 3 4 5 6 7 8 9 10]

```
### array.Union

Get union of two arrays.

```go

a := []int{1, 2, 3, 4, 5}
b := []int{6, 7, 8, 9, 10}

    // filter
    result := array.Union(a, b)

    fmt.Println(result) // [1 2 3 4 5 6 7 8 9 10]

```

### array.Fill

Fill the array with a value.

```go

a := []int{1, 2, 3, 4, 5}
b := array.Fill(a, 0)

fmt.Println(b) // [0 0 0 0 0]

```


### array.Join

Join the array with a separator.

```go

a := []int{1, 2, 3, 4, 5}
b := array.Join(a, "-")

fmt.Println(b) // 1-2-3-4-5

```

### array.Pop

Pop the last element of the array, and return the removed element and the new array.

```go

a := []int{1, 2, 3, 4, 5}
b, c := array.Pop(a)

fmt.Println(b) // 5
fmt.Println(c) // [1 2 3 4]

```

### array.Push

Push an element to the array, and return the new array.

```go

a := []int{1, 2, 3, 4, 5}
b := array.Push(a, 6)

fmt.Println(b) // [1 2 3 4 5 6]

```

### array.Shift

Shift the first element of the array, and return the removed element and the new array.

```go

a := []int{1, 2, 3, 4, 5}
b, c := array.Shift(a)

fmt.Println(b) // 1
fmt.Println(c) // [2 3 4 5]

```
### array.Sort

Sort the array using the function.

```go
a := []int{3, 2, 1, 5, 4}
b := array.Sort(a, func(i, j int) bool { return a[i] < a[j] })

fmt.Println(b) // [1 2 3 4 5]

```

## chaining functions

You can chain the functions together.

```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  a := array.NewArray(data).
    Filter(func(i int) bool {
      return i%2 == 0
    }).
    Map(func(i int) int {
      return i * 2
    }).
    Value()
  
    fmt.Println(a) // [4 8 12 16 20]

```