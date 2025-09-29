package app

import (
	"beginnerGo/internal/api"
	"beginnerGo/internal/middleware"
	"beginnerGo/internal/store"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	UserHandler    *api.UserHandler
	TokenHandler   *api.TokenHandler
	Middleware     middleware.UserMiddleware
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	workoutStore := store.NewPostgresWorkoutStore(pgDB)
	userStore := store.NewPostgresUserStore(pgDB)
	tokenStore := store.NewPostgresTokenStore(pgDB)

	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)
	userHandler := api.NewUserHandler(userStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
	middlewareHandler := middleware.UserMiddleware{UserStore: userStore}

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		UserHandler:    userHandler,
		TokenHandler:   tokenHandler,
		Middleware:     middlewareHandler,
		DB:             pgDB,
	}

	return app, nil
}

// w for response writer, r for request
func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// * write response using Fprintf
	fmt.Fprintf(w, "Status is available\n")
}
