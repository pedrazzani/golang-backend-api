package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pedrazzani/golang/backend-api/controller"
)

var RoutersHandler http.Handler

func init() {
	routers := chi.NewRouter()

	routers.Get("/v1/orders", controller.ListOrders)
	routers.Get("/v1/orders/{id:[0-9]+}", controller.GetOrder)
	routers.Post("/v1/orders", controller.PostOrder)

	RoutersHandler = routers
}
