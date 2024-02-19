# Select

## deferring
- runs statement at the end of the _containing_ function
- helps with making sure you dont forget to clean up resources




## select

- Helps you wait on mulitple channels.
- Sometimes you';; want to include time.After in one of your cases to prevent your system from
blocking forever (since select blocks until a channel responds)


## httptest

- convinient way to create test servers so you can have reliable and controllable tests
- uses the same interface was the "real" net/http servers, so its very consistent
