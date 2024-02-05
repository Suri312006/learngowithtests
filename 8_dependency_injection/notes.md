# Dependency Injection

### What is it?

```go
func Greet(name string){
    fmt.Printf("Hello, %s", name)
}
```

how would we test this code?
- cannot really test the printing to std out

- **need to be able to inject** (fancy for pass in) the dependency of printing


**Our function doesnt need to care _where_ or _how_ the printing happens, so we should accept an _interface_ than a concrete type**




## Summary
- Motivated by our tests, we refactored the code so we could control where the data was written by __injecting a dependecy__ which allowed us to 
test our code, seperate our concerns, and allow our code to be used in different contexts
