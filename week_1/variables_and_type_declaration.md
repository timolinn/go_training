# Variables And Basic Data Types

A variable is a piece of storage containing a value.

Naming in Go follows the following simple rules:

- Must start with a letter or underscore
- can have any number of additional letters or numbers eg. _heap, heapSort, en212
- Go has 25 keywords. These keywords cannot be used as variable names

## Variable Declaration

The four major types of declaration in GO are made with: `const`, `var`, `type`, `func`.

- `var` : A var declaration create a variable of a particular type, attaches a name to it, and give it an initial value.
    basic data types in Go has a default initial value called `Zero Values`:
  - `string`: ""
  - `int`: 0
  - `float64`: 0
  - `bool`: false
  - `interfaces, slices, channels, maps, pointers and functions`: nil
  - `arrays, structs`: fields are set to their respective zero values

    ```go
        var name type = expression
        var i, j, k int = 1, 2, 3
        var l, m, n = 1.5, "one", false
    ```

    Either the type or the = expression part may be omitted, but not both. You cannot assign a diffrent value type to another variable type

    ```go
        var h int
        h = "hello" // won't compile
    ```

### Short Variable Declaration

Short varaible declaration are used to declare and assign initial value to a variable. They work only within the function scope.

```go
    name := expression
    f, err := os.Open(filepath)
```

One subtle but important point: a short variable declaration does not necessarily declare all the variables on its left-hand side. If some of them were already declared in the same lexical block (§2.7), then the short variable declaration acts like an assignment to those variables.

In the code below, the first statement declares both in and err. The second declares out but only assigns a value to the existing err variable.

```go
     in, err := os.Open(infile)
     // ...
     out, err := os.Create(outfile)
```

### Lifetime variables

Package-level variables lasts the entire lifetime of the program execution. local variables live until they become unreachable. New instances are created everytime the declarations are executed

### Tuple Assignments

Tuple assignments allows multiple assigments

```go
    x, y := 2, 4
    a[i], a[j] = a[j], a[i]
```

Some function calls return tuples, so you must enough variables on the left side to hold the returned values.

```go
    f, err := os.Open("path.txt")
    // to discard result use underscores
    _, err := os.Open("foo.txt")

    // you can also use in map lookups
    _, ok := m[key]
```

## Type Declaration
