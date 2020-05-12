# Testing

Unlike some programming languages, Go comes with a robust test tool that makes writing automatted tests for your programs really easy.

The first things you should know about testing in Go are the conventions that are recognised by the Go compiler.

+ Files that contains test cases end with `_test`.
+ Test case functions start with a `Test` prefix.
  For example: say we have a function `Fibonacci(n int)`, that computes the sum of fibonacci numbers through `n`, if we were to write a function to test our functionality we'll go ahead and call it `TestFibonacci`
+ Go comes with a tool called `go test`. This command scans through your project for the test files, and executes all availble test cases. It will return an error for failed tests and vice versa.

Go ahead and run `go test -v` command withing the current folder. Try simulating a failing test by modifying the `Fibonacci` function.
