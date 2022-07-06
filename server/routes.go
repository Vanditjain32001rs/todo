package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"todo/handler"
	"todo/helper"
)

type Server struct {
	chi.Router
}

func SetUpRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/user", func(api chi.Router) {
		api.Post("/add-user", handler.CreateUser)
		api.Post("/sign-in", handler.SignIn)
		api.Group(func(r chi.Router) {
			r.Use(helper.Middleware)
			r.Use(helper.RecoveryMiddleware)
			r.Get("/fetch-task", handler.FetchTask)
			r.Post("/add-task", handler.AddTask)
			r.Put("/update-task", handler.UpdateTask)
			r.Delete("/delete-task", handler.DeleteTask)
			r.Put("/completed-task", handler.MarkTaskComplete)
			r.Put("/log-out", handler.LogOut)
		})
	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
