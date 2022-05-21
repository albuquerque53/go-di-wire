# Dependency Injection?

Basically, dependency injection (DI) is the pattern that consists into **inject** the dependencies of your **components** right when they are **created**.

## Why?

The DI pattern is very useful when we want separate the logic of **instance** a component from the logic of ***use** the component properly.

Supposing that we have the **ComponentA** and this component have as dependency the **ComponentB** (`ComponentB` must have others dependencies too). The DI says that **ComponentA** must not have know how to instance **ComponentB** and your dependencies.

So that, we have the concerns separated, the concern of **intance** a dependency is not attached to **use** the dependency. This separation **assert us that our code are more readable, testable and mantainable**.
