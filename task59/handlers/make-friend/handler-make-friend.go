package makefriend

import (
	"curse/task59/logger"
	"encoding/json"
	"io"
	"net/http"
)

type MakeFriend struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type UserActionsForHandlerMakeFriendForUsers interface {
	MakeFriend(sourceID, targetID int) (string, error)
}

type HandlerForMakeFriendForUsers struct {
	log         logger.Logger
	userActions UserActionsForHandlerMakeFriendForUsers
}

func NewHandlerForMakeFriend(log logger.Logger, userActions UserActionsForHandlerMakeFriendForUsers) *HandlerForMakeFriendForUsers {
	result := &HandlerForMakeFriendForUsers{
		log:         log,
		userActions: userActions,
	}

	return result
}

func (handler *HandlerForMakeFriendForUsers) prepareRequest(request *http.Request) (*MakeFriend, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var makeFriendFromClient *MakeFriend

	if err := json.Unmarshal(body, &makeFriendFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	return makeFriendFromClient, nil
}

func (handler *HandlerForMakeFriendForUsers) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	shouldMakeFriend, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		write.WriteHeader(http.StatusBadRequest)

		return
	}

	status, err := handler.userActions.MakeFriend(shouldMakeFriend.SourceID, shouldMakeFriend.TargetID)
	if err != nil {
		handler.log.Printf("can not make friends: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := write.Write([]byte(status)); err != nil {
		handler.log.Printf("can not write: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}
}
