package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HandlerMakeFriends struct {
	log Logger
}

func NewHandlerMakeFriends(log Logger) *HandlerMakeFriends {
	return &HandlerMakeFriends{
		log: log,
	}
}

func (handler *HandlerMakeFriends) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		makeNewFriend *MakeNewFriendInfo
		allUsersInDB  Users
	)

	dataReadAllUsersInDB, err := os.ReadFile("./database.json")
	if err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.Unmarshal(dataReadAllUsersInDB, &allUsersInDB); err != nil {
		handler.log.Printf("err = %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := json.Unmarshal(body, &makeNewFriend); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	sourceUser, err := FindUserByID(allUsersInDB, makeNewFriend.SourceID)
	if err != nil {
		handler.log.Printf("cannot find user %d", makeNewFriend.SourceID)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	targetUser, err := FindUserByID(allUsersInDB, makeNewFriend.TargetID)
	if err != nil {
		handler.log.Printf("cannot find user %d", makeNewFriend.TargetID)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.ID)
	targetUser.Friends = append(targetUser.Friends, sourceUser.ID)

	dataWriteAllUsersInDB, err := json.Marshal(allUsersInDB)
	if err != nil {
		handler.log.Printf("cannot marshal allUsersInDB")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := os.WriteFile("./database.json", dataWriteAllUsersInDB, os.ModePerm); err != nil {
		handler.log.Printf("cannot write allUsersInDB in file")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := fmt.Sprintf("%s and %s are friends now", sourceUser.Name, targetUser.Name)

	if _, err := writer.Write([]byte(response)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
