# Injector in Wire

## Concept

A function that use providers to inject the dependencies. With wire you can create an injector with a simple function that call `wire.Build`.

```go
package example

func InitializeExample() (Example, error) {
    wire.Build(Provide1, Provide2, Provide3)

    return Example, nil
}
```

After that, wire can produce a implementation of the injection running, simply running `wire` in console. A file with this implementation will be created.
