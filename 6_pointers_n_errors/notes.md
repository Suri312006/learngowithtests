# Pointers & Errors

## Pointers
- Pointers let us point to some values and then let us change them

```go
// this wouldnt actually work because the w in `(w Wallet)` is a copy
// of whatever we called the method from
func (w Wallet) Deposit(amt int ) {
    // so when we try to change its value, we arent actually changing its value 
    w.balance += amt
    
    fmt.Printf("address of balance in test is %p \n", &w.balance)
}
```

- Pointers can be nil
- when a function returns a pointer to something
make sure you check if its nil, otherwise a runtime exception could be raised

## Errors
- way to signify failure when calling function / method
- use errors in tests! 


