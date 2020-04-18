# Interfaces

Interfaces are named collections of method signatures OR An `interface` is a set of method signatures. Interfaces can also be referred to as "contracts". Like structs they are also data types, but they are "abstract" data types, in that they do not hold any concrete data and implementation details. This means that they can take the form anything you want them to. Interfaces improves loose coupling by making it possible to write more flexible functions. Unlike many OOP languages, Go's interfaces are satisfied implicitly, there's no `implements` keyword.

```go
    type SocialMedia interface {
        Feed() []string
        Posts() []string
    }
```

The code snippet above is an interface type `SocialMedia`.

Like I said earlier, interfaces can be seen as contracts to be fulfilled by other types. Any type that fulfills an interface (a contract) can represent that interface.

## The empty interface

An empty interface is usually unnamed and contains zero methods. They can can hold values of any type. They are often used to handle values of unknown types.

```go
    var i interface{}
    describe(i)

    i = 42
    describe(i)

    i = "hello"
    describe(i)
    func describe(i interface{}) {
        fmt.Printf("(%v, %T)\n", i, i)
    }

    func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
    func Printf(format string, args ...interface{}) (int, error) {
        return Fprintf(os.Stdout, format, args...)
    }
    func Sprintf(format string, args ...interface{}) string {
        var buf bytes.Buffer
        Fprintf(&buf, format, args...)
        return buf.String()
    }
```
