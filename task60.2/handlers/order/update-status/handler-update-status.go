package updatestatus

import (
	"curse/task60.2/logger"
	"curse/task60.2/models"
	"encoding/json"
	"io"
	"net/http"
)

type UpdateStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type OrderActionsForHandlerUpdateStatus interface {
	UpdateStatus(orderID int, status string) (*models.Order, error)
}

type HandlerForUpdateStatus struct {
	log          logger.Logger
	orderActions OrderActionsForHandlerUpdateStatus
}

func NewHandlerForUpdateStatus(log logger.Logger, orderActions OrderActionsForHandlerUpdateStatus) *HandlerForUpdateStatus {
	result := &HandlerForUpdateStatus{
		log:          log,
		orderActions: orderActions,
	}

	return result
}

func (handler *HandlerForUpdateStatus) prepareRequest(request *http.Request) (*models.Order, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var newStatusFromClient *UpdateStatus

	if err := json.Unmarshal(body, &newStatusFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newStatus := &models.Order{
		ID:     newStatusFromClient.ID,
		Status: newStatusFromClient.Status,
	}

	return newStatus, nil
}

func (handler *HandlerForUpdateStatus) sendResponse(writer http.ResponseWriter, updatedStatus *models.Order) {
	updatedStatusMarshaled, err := json.Marshal(updatedStatus)
	if err != nil {
		handler.log.Printf("cannot marshal updated status: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(updatedStatusMarshaled); err != nil {
		handler.log.Printf("cannot send to client updated status: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerForUpdateStatus) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldUpdateStatus, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	updatedStatus, err := handler.orderActions.UpdateStatus(shouldUpdateStatus.ID, shouldUpdateStatus.Status)
	if err != nil {
		handler.log.Printf("cannot update status: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, updatedStatus)
}
