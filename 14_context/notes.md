# Context


- Context package provides functions to derive new Context values from existing ones
. these values form a tree: when a Context is canceled, all Contexts
derived from it are also canceled




- incoming requests to server should create a Context and outgoing calls should accept a Context
    - when a Context is cancelled, all contexts derived from it are also cancelled



- at google its required to pass a Context parameter as the first argument to every function on the 
call path between incoming and outgoing requests. This allows good interoperation
