# Type Declarations

Type declaration is where we declare a new __named type__, like `int`, `float64` etc, Go gives us the ability to create our own named types that shares the same behaviour with an existing type and at the same time flexible enough that we cannot confuse two separate _named types_ `type` as being the same. For eg. `int value` == `int value` is true, but `named type` == `named type` may not be true. The underlying type of a _named type_ determines its str ucture and represent ation, and also the set of int rinsic operations it supports, which are the same as if the underlying type had been used directly. To declare a _named type_:

```go
    type NewType struct {}
    type AliasFloat float64
```

## Pointers

A pointer value is the address of a variable. A pointer is thus the location at which a value is stored. Not every value has an address, but every variable does. If a variable is declared var x int, the expression &x ("address of x") yields a pointer to an integer variable, that is, a value of type *int, which is pronounced "pointer to int."

```go
    var n *int
    n := new(int) // returns *int
```

Each call to new returns a distinct variable with a unique address:

```go
     p := new(int)
     q := new(int)
     fmt.Println(p == q) // "false"
```

There is one exception to this rule: two variables whose type carries no information and is therefore of size zero, such as `struct{}` or `[0]int`, may, depending on the implementation, have the same address.

Pointers makes it possible to pass a parameter by reference, as opposed to Go's default _pass by value_. Consider the code below, what do you think it prints?

```go
    import "fmt"

    type Elves struct {
        King   string
        Number int
    }

    func main() {
        e := Elves{}
        Number(&e)
        fmt.Println(e.King)
    }

    func SetKing(e *Elves) {
        e.King = "Fingon"
    }
```

### The `reference` (&) and the `dereference` (*) operators

The reference operator can also be referred to as `address of` operator, it is used to fetch the memory address of a value or a variable. As mentioned above, you can fetch the memory address of variable `x` by prefixing it with the reference operator ie. `&` (ampersand) the returned value is otherwise known as a pointer. In other words, a pointer value of variable `x` is `&x`.

The deference operator does the opposite of what the reference operator does but can only work or can only be used on pointer expressions eg. `&x`. Assuming we get the pointer to a variable `y` by doing `var y = new(int)` or `var y = &x`, you'll notice we cannot assign a new value with to `y`  without using the `*` operator. Trying to print `y` using the `fmt.Println` method would return a memory address instead of the expected value. Consider the code below:

```go
    func main() {
        var x = 5
        var y = &x // assigns variable x's memory address to "y"
        fmt.Println(y)
    }
```

Running the code above would print a memory address not the actual value. Now, prefix `y` with the `*` operator ie. `fmt.Println(*y)` and try again. The `*` operator deferences the value from the location where it's stored in your computer memory. This simply means going to the location and retriving the actual value. Consider thecode below:

```go
    func main() {
        var y = new(int) // the new built-in function returns a pointer to int ie. "*int"
        y = 5 // compile error
    }
```

The above code won't compile because `y` is of type `*int` (pointer to an `int`), while the value `5` is of type `int`, classic "type mismatch error". To fix this, we can go ahead and prefix `y` with the deference operator `*` ie `*y = 5`, then try compiling again.

## Functions and Methods

Functions are first class types in Go. You'll see them "everywhere you Go", they provide us with the means to wrap a sequence of statements and expressions into a single unit.

```go
    type Add func(a, b int) int

    func Name(parameters type) (return types) {
        return type, type
    }
```

Functions can return multiple values in Go, once a return type has been defined the function must contain a return statement, unless the function clearly going to panic your code won't compile.

### Methods

Methods are functions attached to a type. Like in Java or PHP, the functions of an existing class is called a method, they are responsible for adding arbitrary behaviour to Objects. The same principle applies to methods in Go. However, Go is pro `composition over inheritance`, hence methods are `attached` to structs in a completely different way as in Java et al.

```go

    func(e Elves) GetKing() string {
        return e.king
    }

    func main() {
        // ...
        fmt.Println(elf.GetKing())
    }
```
