# Packages

Go comes with over 100 standard packages, there are several thousands written by a thriving Go community.

Packages group related features together, by doing this managing large programs become practical. Packages provide Go with an OOP feature known as `encapsulation` by controlling the visibility of package members, only exposing exported members as the package API.

```go
    package main

    import (
        "fmt"
        "crypto/rand"
        math "math/rand"

        _ "github.com/lib/pq" // blank import
        _ "github.com/go-sql-driver/mysql"
    )
```

## Packages and Naming

Use short names but not so short to be cryptic. Also be descriptive and unambigous eg. Don't name your package `util`, instead use specific names like `imageutil` or `ioutil`.

> Avoid package level state - This reduces the possibility of introducing breaks when a change is made.

## The Go tool
