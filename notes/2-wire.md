# Tools for DI

## In Go

Dependency injection is so important, that there are quite a few solutions for this in the Golang community already, such as dig from Uber and inject from Facebook. Both of them implement runtime dependency injection through Reflection Mechanisms.

## Wire

Wire is a code generation tool that automates connecting components using dependency injection. Dependencies between components are represented in Wire as function parameters, encouraging explicit initialization instead of global variables.

### Why use Wire?

1. As said, `Wire` is a **code generation** tool, means that the generated code is **obvious** and **readable**.
2. `Wire` runs in **compile-time**, this provide-us an easy debug. If any dependency is missing or being unused, an error will be reported during compiling.
