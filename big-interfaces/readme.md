# Big Interfaces

What I like:

- I only have to pass one `DatabaseClient` to all my HTTP handlers to provide them with access to all services.
- I can use service mocks in my handler tests instead of using a real database.

What I don't like:

- Too big interfaces in package `whatever`:
  - `UserService` has basic CRUD methods plus lots getters etc.
  - `DatabaseClient` has database related methods plus getters for all services.

Setup:

~~~
$ git clone <insert URL here>
$ cd <insert dir here>
$ cd cmd/web
$ go run main.go
$ open http://localhost:8080/users/123
~~~
