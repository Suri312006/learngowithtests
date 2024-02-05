# Maps

- looks like map is similar to pythons dictionaries
```go

dictionary := map[string]string{"something": "is not anything"}
nums := map[int]string{2: "is not anything"}

```

- You dont have to deference maps to mutate them
```go
func (d Dictionary) Add(key string, value string) {
	d[key] = value
}
```

- this is because when you pass a map to a function / method, you are copying 
just the pointer part, not the underlying data structure

- maps can be a nil value
    - you can read an empty map
    - but writing a nil map will cause a runtime panic

### Therefore you should never instantiate an empty map variable 

##### Instead you can do the following 

```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)

```
- both approaches create an empty hashmap and point dictionary at it (ensures you will never get a runtime panic)
