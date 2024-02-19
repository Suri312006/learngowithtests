# Sync

### Mutex
- allows us to add locks to our data
- `Wait Group` is a means of waiting for goroutines to finish jobs

- common go newbie mistake is to use channels and goroutines just because
its possible and fun
- __Dont be afraid to use a sync / Mutex__

### use channels when passing ownership of data
### use mutexes for managing state


## `go vet`
- alerts you to subtle bugs in your code before they hit users

### __Dont use embeding because it's convenient__
- think about public facing api
- think about what is getting exposed
- remember doing this with mutexes could be disastourous
