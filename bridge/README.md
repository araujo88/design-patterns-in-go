# Bridge

The Bridge Design Pattern is a structural design pattern that aims to decouple an abstraction from its implementation so that the two can vary independently. This is especially useful for providing platform independence and can also enhance extensibility.

In other words, Bridge Pattern is about preferring composition over inheritance. Implementation details are pushed from a hierarchy to another object with a separate hierarchy.

## Example

For example, consider different types of database drivers. We can have an interface called `DatabaseDriver` that has methods like `Connect()`, `ExecuteQuery()`, etc. Now, the actual implementations like `MySQLDriver`, `PostgresDriver`, `SQLiteDriver` will implement this `DatabaseDriver` interface and provide the actual implementation of these methods. So the abstraction layer (which could be an application using the database) will just call these methods and is not worried about what database system it is connecting to. This abstraction of implementation is the essence of the bridge pattern.

```go
type IDatabase interface {
    Connect() string
}

type MySQL struct {
}

func (m MySQL) Connect() string {
    return "Connected to MySQL Database"
}

type Postgres struct {
}

func (p Postgres) Connect() string {
    return "Connected to Postgres Database"
}

type Application struct {
    database IDatabase
}

func (a *Application) SetDatabase(database IDatabase) {
    a.database = database
}

func (a *Application) Start() {
    fmt.Println(a.database.Connect())
}

```

In this example, we define an interface `IDatabase` that each of our specific database types (`MySQL`, `Postgres`) implements. This interface includes a method `Connect()`, which would be different for each database type. The `Application` struct then has a `SetDatabase()` method that accepts anything that fulfills the `IDatabase` interface.

We then have a method `Start()` that prints the `Connect()` message. In a real-world application, it would contain the logic that is common to all databases. It could be interacting with the database, creating tables, performing CRUD operations, etc.

By using the bridge pattern in this way, we're able to decouple the abstraction (our application's interaction with a database) from the implementation (the specific calls needed to interact with a MySQL or a PostgreSQL database).
