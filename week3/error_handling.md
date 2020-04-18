# Error Handling

> "Don't just check errors, handle them gracefully" - Go Proverb

Error handling is as important as your application features, even better is handling them gracefully with the experience of your users in mind.

Go come's with an `error` interface that's so simple makes error handling really simple. The fact that it provides a unified interface for handling errors across different standard library packages and third party packages.

```go
    type error interface {
        Error() string
    }
```

Go also comes with an `errors` package that provides us with probably the easiest way to customize error messages. We can do this using the `New` function.

```go
    return errors.New("an error occured", err)
```

## Reference

+ <https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully>
+ <https://dave.cheney.net/2019/01/27/eliminate-error-handling-by-eliminating-errors>