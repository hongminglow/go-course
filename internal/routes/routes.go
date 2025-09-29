package routes

import (
	"beginnerGo/internal/app"

	"github.com/go-chi/chi/v5"
)

// * Using chi router to setup routes
// * chi provides a rich set of middleware and routing features
func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// * Route group for /workouts endpoints
	// * Apply authentication middleware to this group
	// * RequireUser middleware ensures that the user is authenticated
	// * If not authenticated, it will return a 401 Unauthorized response
	// * If authenticated, it will call the respective handler function
	r.Group(func(r chi.Router) {
		// * r.Use to apply middleware to all routes in the group
		// * while "app.Middleware.RequireUser" ensure that this specific middleware is applied to the route handlers
		r.Use(app.Middleware.Authenticate)
		r.Get("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleGetWorkoutByID))
		r.Post("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.HandleCreateWorkout))
		r.Put("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleUpdateWorkoutByID))
		r.Delete("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleDeleteWorkoutByID))
	})

	r.Get("/health", app.HealthCheck)
	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r
}
