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
