This is an example of using Github actions to perform integration tests between golang and postgreSQL.

This example package contains an interface `models.UserModel`.

```go
type UserModel interface {
	Get(id int) (*User, error)
}
```

In a production application you'll want to use the `postgres.UserModel` that implements this interface and connects to postgreSQL.

In your unit tests you'll want to mock this model with `mocks.UserModel`. And in your integration tests you'll want to connect to a database instance to test you postgres models.

This is an example of how to achieve all this with Github actions.
