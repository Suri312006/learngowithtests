# Reflection
golang challenge: write a function walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside. difficulty level: recursively.


- Reflection in computing is the ability of a program to example its own structure, particularly through types; its a form of metaprogramming. It's also a great source of confusion.




### What is interface{}

`interface{}` is kind of like `any` in typescript
- lets you write a function where you dont know the type at compile time

walk(x interface{}, fn func(string)) will accept any value for x.

- as the writer of such a function, you have to be able to inspect _anything_ that is passed into it, and try to figure out what you can do with it.
- this is done using ___reflection___, but it can be quite clumsy and less performant

#### Only use reflection if you really need to 

- if you want polymorphic functions, maybe you could design it around an interfaceo
## Which is different that `interface{}`

 
