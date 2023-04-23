package main

import (
	"curse/task60.2/datasource/dispatcher"
	handler_create "curse/task60.2/handlers/order/create"
	handler_read_all_orders "curse/task60.2/handlers/order/read-all"
	"curse/task60.2/logger"
	"net/http"

	"github.com/go-chi/chi"
)

const addr = ":8080"

func main() {
	/*
		--------------

				POST /api/v1/orders/create
				body:
				{
					"address": "where to delivery",
					"telephone": "+375 whatever",
				}

				response: 200
				{
					"id": 444,
					"address": "where to delivery",
					"telephone": "+375 whatever",
					"status": "in-route"
				}

		--------------

				POST /api/v1/orders/update_address
				body:
				{
					"address": "new address"
				}

				response: 200

		--------------

				POST /api/v1/orders/update_status
				body:
				{
					"status": "in-route" | "delivered" | "canceled"
				}

				response: 200

		--------------

				POST /api/v1/orders/update_telephone
				body:
				{
					"telephone": "+375 new phone"
				}

				response:200

		--------------

				GET /api/v1/orders
				response: 200
				[
				{
					"id": 445,
					"address": "where to delivery",
					"telephone": "+375 whatever",
					"status": "in-route"
				},
				{
					"id": 444,
					"address": "where to delivery",
					"telephone": "+375 whatever",
					"status": "canceled"
				},
				...
				]

		--------------
	*/

	router := chi.NewRouter()
	log := logger.New()

	dispatcher := dispatcher.NewDispatcher(log)

	handlerCreateOrder := handler_create.NewHandlerForCreateOrder(log, dispatcher)
	router.Method(http.MethodPost, "/api/v1/orders/create", handlerCreateOrder)

	handlerReadOrders := handler_read_all_orders.NewHandlerForReadOrder(log, dispatcher)
	router.Method(http.MethodGet, "/api/v1/orders", handlerReadOrders)

	server := NewServer(addr, router)

	log.Printf("Serving at [%s]", addr)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func NewServer(address string, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
