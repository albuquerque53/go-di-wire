# Provider in Wire

## Concept

Provider is a function that can produce some value. An ordinary function that instance something.

```go
package example

type Example struct {
    exVal int
}

func NewExample() Example {
    return Example{exVal: 0}
}
```

The provider functions must be exported to instance values in other packages.

## With providers you can

1. Specify the dependencies with params.
2. Return errors.
3. Create group of provider sets
4. Add others providers into provider sets
