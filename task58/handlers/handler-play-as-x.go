package handlers

import (
	"net/http"
)

type HandlerForPlayAsX struct {
	log Logger
}

func NewHandlerForPlayingX(log Logger) *HandlerForPlayAsX {
	return &HandlerForPlayAsX{
		log: log,
	}
}

func (handler *HandlerForPlayAsX) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}
