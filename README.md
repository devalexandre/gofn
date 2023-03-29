# gofn
library to use golang functional


## Usage

### array.Filter
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Filter(data, func(i int) bool {
        return i%2 == 0
    })

    fmt.Println(result)
	
```

### array.Map
```go

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // filter
    result := array.Map(data, func(i int) int {
        return i * 2
    })

    fmt.Println(result)

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

    fmt.Println(result)

```