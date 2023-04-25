package main

import (
	"curse/task60.3/datasource/dispatcher"
	"curse/task60.3/datasource/mongo"
	handler_create "curse/task60.3/handlers/order/create"
	handler_read_all_orders "curse/task60.3/handlers/order/read-all"
	handler_update_address "curse/task60.3/handlers/order/update-address"
	handler_update_status "curse/task60.3/handlers/order/update-status"
	handler_update_telephone "curse/task60.3/handlers/order/update-telephone"
	"curse/task60.3/logger"
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
					"id": 5,
					"address": "new address"
				}

				response: 200

		--------------

				POST /api/v1/orders/update_status
				body:
				{
					"id": 5,
					"status": "in-route" | "delivered" | "canceled"
				}

				response: 200

		--------------

				POST /api/v1/orders/update_telephone
				body:
				{
					"id": 5,
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

		curl -i -X POST --data-binary '{"address": "second order","telephone": "+375 whatever"}' \
			http://localhost:8080/api/v1/orders/create

		curl -i -X POST --data-binary '{"id":2,"telephone":"my-telephone"}' \
			http://localhost:8080/api/v1/orders/update_telephone

		curl -i -X POST --data-binary '{"id":2,"status":"in-route"}' \
			http://localhost:8080/api/v1/orders/update_status

		curl -i -X POST --data-binary '{"id":2,"address":"KRY tam zhiviet"}' \
			http://localhost:8080/api/v1/orders/update_address

		curl -i -X GET http://localhost:8080/api/v1/orders
	*/

	router := chi.NewRouter()
	log := logger.New()

	moooong := mongo.New(log, "", "", []string{"localhost:27017"}, "my-database")
	dispatcher := dispatcher.NewDispatcher(log, moooong)

	handlerCreateOrder := handler_create.NewHandlerForCreateOrder(log, dispatcher)
	router.Method(http.MethodPost, "/api/v1/orders/create", handlerCreateOrder)

	handlerReadOrders := handler_read_all_orders.NewHandlerForReadOrder(log, dispatcher)
	router.Method(http.MethodGet, "/api/v1/orders", handlerReadOrders)

	handlerUpdateAddress := handler_update_address.NewHandlerForUpdateAddress(log, dispatcher)
	router.Method(http.MethodPost, "/api/v1/orders/update_address", handlerUpdateAddress)

	handlerUpdateStatus := handler_update_status.NewHandlerForUpdateStatus(log, dispatcher)
	router.Method(http.MethodPost, "/api/v1/orders/update_status", handlerUpdateStatus)

	handlerUpdateTelephone := handler_update_telephone.NewHandlerForUpdateTelephone(log, dispatcher)
	router.Method(http.MethodPost, "/api/v1/orders/update_telephone", handlerUpdateTelephone)

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
