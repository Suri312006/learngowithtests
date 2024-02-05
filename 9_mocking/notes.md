# Mocking

### Its an important skill to be able to slice up requirements as small as you can so you can have working software


we are trying to spy on sleep, so we can test that sleeping works, without the test taking
the same time as sleep


### if mocking seems evil, youre likely doing something wrong

- if it seems complicated, then it probably is too complicated, break things down

- The thing you are testing is having to do too many things (because it has too many dependencies to mock)
    - break the module apart so it does less
- if the dependencies are too fine grained (whatever that means)
    - think about how you can consolidate some of these dependencies into one meaningful module
- your test is too concerened with the implementation details
    - favor testing expected behavior rather than the implementation
