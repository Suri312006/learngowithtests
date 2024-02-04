# Arrays and Slices

- Arrays have a fixed capacity, which you define and declare the variable
    - initialized in 2 ways
        - [N] type {val 1, val 2, ..., val N}
            - number:= [5]int{1, 2, 3, 4, 5}
        - [...]type{val 1, val 2, ..., val N}
            - numbers := [...]int{1, 2, 3, 4, 5}

- Arrays size is encoded in its type

- slices do not encode the size of the collection, and can have any size
    - like vector in rust 
    - hello 
