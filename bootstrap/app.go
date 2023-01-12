package bootstrap

import (
	"database/sql"
	"log"
	"server/database/store"
)

type Application struct {
	Env              Env
	Store            store.Store
	PostgresqlClient *sql.DB
}

func App() Application {
	app := &Application{}
	env, err := NewEnv(".", "app")
	if err != nil {
		log.Fatal(err)
	}

	app.Env = env
	app.PostgresqlClient = NewPostgresqlClient(&env)
	app.Store = store.NewStore(app.PostgresqlClient)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresqlConnection(app.PostgresqlClient)
}
