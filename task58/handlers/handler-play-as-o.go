package handlers

import "net/http"

type HandlerForPlayAsO struct {
	log Logger
}

func NewHandlerForPlayingO(log Logger) *HandlerForPlayAsO {
	return &HandlerForPlayAsO{
		log: log,
	}
}

func (handler *HandlerForPlayAsO) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}
