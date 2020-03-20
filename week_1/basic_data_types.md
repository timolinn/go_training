# Basic Data Types in Go

Go types are divided into four categories:

+ basic types - numbers, strings, floats and booleans
+ reference types - pointers, slices, maps, functions and channels
+ aggregate types - arrays and structs
+ interface types - interfaces

They include: `int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64 bool byte rune string error`.

## Integers

Go has multiple size numerical data types as you can see above. There are many integer sizes, floats and complex numbers.

A rune can be seen as int32, runes are a Unicode code point values. Singned values are represented in 2s eg:

```go
    int8 = -2^(n-1) => 2^(n-1) - 1
    unint8 = 0 => 2^n - 1
    int8 = -127 => 128
    unint8 = 0 => 255
```

Go's arithmetic, logic and comparison operators are listed below in order off decreasing precedence:

```go
    * / % << >> & &^ + - | ^
    == != < <= > >=
    &&
    ||
```

Type conversion: This is not the same as the regular string to int conversion mechansims in many other programming languages.

```go
    var num int16 = 1
    var num int = num // compile error
    var num int = int(num) // works
```

## Floating Point Numbers

Go has two floating number types `float32` and `float64`. Float32 has 6 digit preicision and float64 has 15 digits precision point. Casting floats to int causes them to loose the precision point.

```go
    var PI float32 = 3.142
```

## Booleans

Booleans are either true or false. 1 or 0 cannot be used as true or false substitute in a comparison statement

## Strings

A string is an immutable sequence of bytes. Strings may contain arbitrary data, including bytes with value 0, but usually they contain human-readable text. Text strings are convention- ally interpreted as UTF-8-encoded sequences of Unicode code points (runes)

```go
    s := "hello, world"
     fmt.Println(len(s))     // "12"
     fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')
```

To substring a string we can do `s[j:l]`, it returns the char from j to and not including l. j and l can be omitted

```go
    fmt.Println(s[:5]) // "hello"
     fmt.Println(s[7:]) // "world"
     fmt.Println(s[:])  // "hello, world"
```

Concatenation can be done using the unary operator (+). Also, Strings a immutable. You cannot do `s[0] = 'L' // compile error: cannot assign to s[0]`
