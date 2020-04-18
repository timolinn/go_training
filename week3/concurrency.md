# Concurrency

> Concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the final outcome. - Wikipedia

Concurrency in a simpler terms means executing multiple tasks in a disorganised order, thereby making progress on more than one task simultaneously. Go has a rich support for concurrency using goroutines, channels and the sync package.

## Goroutines

Goroutines are lightweight threads that can be referred to as `green threads`. Threads are a known component of an operating system's process, they contain tasks or procedures to be executed in a process. A process is basically a program in execution. Your OS creates new processes for each program you launch. Processes are managed directly by the Operating system, memory allocations, CPU resources etc. To create a goroutine use the `go` keyword.


```go
    import (
        "fmt"
        "time"
    )

    func say(s string) {
        for i := 0; i < 5; i++ {
            time.Sleep(100 * time.Millisecond)
            fmt.Println(s)
        }
    }

    func main() {
        go say("world") // this spawns a new goroutine
        say("hello")
    }
```

We should note that the `main` function itself is a goroutine. Try running the code snippet above and see what happens.

Notice that in our `main` function the first call to the `say` function does not block the code execution? This is because the `go` keyword created a new goroutine that.

There are multiple concurrency models, they include:

+ Events
+ Process Threads
+ Callbacks and Promises
+ Communication Sequential Process (CSP)

The Go concurrency model implements the communicating sequential process or CSP. This models does not follow the usual way of passing information between threads by sharing memory (common in multithreaded proramming), instead goroutines share memory by communicating through `channels`.

## Channels

A channel is a communication mechanism that lets goroutines send values to another goroutine. To create a channel that accepts only `int`s we use the `make` function.

```go
    ch := make(chan int)
```

A channel can either `send` or `receive` data.

```go
    ch <- x // send statement
    x = <- ch // receive statement, received data is assigned to x
    <- ch // received data is discarded
```

You can close a channel with the `close` function.

```go
    close(ch)
```

A channel can also be buffered, buffered channels cannot receive or send more than the specified capacity.

```go
    ch := make(chan int, 2) // capacity of 2
```

A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel, at which point the value is transmitted and both goroutines may continue.
