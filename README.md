# gofn
![Coverage](https://img.shields.io/badge/Coverage-80.7%25-brightgreen)
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
    - [Use two functions](#use-two-functions)

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

Reverse the array.
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Reverse(data)

    fmt.Println(result) // [10 9 8 7 6 5 4 3 2 1]

```

## Use two functions
You can use two types of function, like Filter and Map.
```go

    // create slice
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Map(array.Filter(data, func(i int) bool {
        return i%2 == 0
    }), func(i int) int {
        return i * 2
    })

    fmt.Println(result) // [4 8 12 16 20]

```