package bootstrap

import (
	"database/sql"
	"log"
	"server/database/store"
	"server/internal/tokenutil"

	"github.com/labstack/echo/v4"
)

type Application struct {
	Env              Env
	TokenMaker       tokenutil.Maker
	Store            store.Store
	PostgresqlClient *sql.DB
}

func App() Application {
	app := &Application{}
	env, err := NewEnv(".", "app")
	if err != nil {
		log.Fatal(err)
	}

	tokenMaker, err := tokenutil.NewPasetoMaker(env.PasetoSymmetricKey)
	if err != nil {
		log.Fatal(err)
	}

	app.Env = env
	app.TokenMaker = tokenMaker
	app.PostgresqlClient = NewPostgresqlClient(&env)
	app.Store = store.NewStore(app.PostgresqlClient)
	return *app
}

func (app *Application) Launch(e *echo.Echo) {
	log.Fatal(e.Start(app.Env.Host + ":" + app.Env.Port))
}

func (app *Application) CloseDBConnection() {
	ClosePostgresqlConnection(app.PostgresqlClient)
}
