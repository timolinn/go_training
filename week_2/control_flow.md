# Control Flow

In computer science, control flow (or flow of control) is the order in which individual statements, instructions or function calls of an imperative program are executed or evaluated.

## If Else

```go

    if condition {
        // do something
    } else if condition {
        // do something else
    } else {
        // do something
    }

    a := 3
    b := 4

    if a == b { // == > < >= <=
        fmt.Println("a is equal to b")
    } else {
        fmt.Println("a is not equal to b")
    }
```

## Switch

```go
    switch expression {
    case condition:
        // do something
    case condition:
        // do something
    default:
        // do something
    }
```
