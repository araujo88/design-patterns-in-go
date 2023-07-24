package main

import "fmt"

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

func main() {
	mysql := &MySQL{}
	postgres := &Postgres{}

	app := &Application{}
	app.SetDatabase(mysql)
	app.Start()

	app.SetDatabase(postgres)
	app.Start()
}
