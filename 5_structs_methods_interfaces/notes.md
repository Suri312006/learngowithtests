# Structs Methods Interfaces


## Structs
- there are anonymous structs 
- you can use structs to create new types on them
- you can use methods to give those structs some functionality

## Methods
Method &rarr; Function with a receiver.
 - declaration binds an identifier, the method naem, to a method and associates the message with the revievers base type


 - called by invocation on an instance of a particular type

In go, it is convention for the receiver variable be the first letter of the type

``` go
// this is the reviever  w the receiver type
func (r Rectangle) Area() float64 {
    return 0
}

```
## Interfaces
 - the interface is applied to anything with the Area method that returns a float64
``` go
type Shape interface {
    Area() float64
}
```

- go's interface resolution is implicit. If the type you pass in matches what
the interface is asking for, it will compile



# General Wrapup
- use structs to have custom type
- use methods to give structs functionality
- use interfaces to create a "family of methods"? 


# Table driven tests

- more organized way of doing testing if you have a bunch of structs with 
a similar interface. look at shapes_testing.go
