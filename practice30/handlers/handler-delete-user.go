package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func DeleteUser(allUsersInDB Users, id int) Users {
	newArray := Users{}

	//for _, user := range allUsersInDB { - for dumbs
	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]

		if user.ID != id {
			newArray = append(newArray, user)
		} else if user.ID == id {
			continue
		}
	}

	return newArray
}

func DeleteUserFromFriends(allUsersInDB Users, id int) Users {
	newArrayOfFriends := Users{}

	for i := 0; i < len(allUsersInDB); i++ {
		user := allUsersInDB[i]
		arrayOfFriendsOfUser := []int{}

		for j := 0; j < len(user.Friends); j++ {
			friendID := user.Friends[j]

			if friendID != id {
				arrayOfFriendsOfUser = append(arrayOfFriendsOfUser, friendID)
			} else if friendID == id {
				continue
			}
		}

		user.Friends = arrayOfFriendsOfUser
		newArrayOfFriends = append(newArrayOfFriends, user)
	}

	return newArrayOfFriends
}

type HandlerDeleteUser struct {
	log Logger
}

func NewHandlerDeleteUser(log Logger) *HandlerDeleteUser {
	return &HandlerDeleteUser{
		log: log,
	}
}

func (handler *HandlerDeleteUser) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		deleteUser   *DeleteUserInfo
		allUsersInDB Users
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

	if err := json.Unmarshal(body, &deleteUser); err != nil {
		handler.log.Printf("cannot unmarshal body", string(body))
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	userDeleted, err := FindUserByID(allUsersInDB, deleteUser.ID)
	if err != nil {
		handler.log.Printf("cannot find user %d", deleteUser.ID)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	allUsersInDBWithoutDeletedUser := DeleteUser(allUsersInDB, userDeleted.ID)
	allFriendsWithoutDeletedUser := DeleteUserFromFriends(allUsersInDBWithoutDeletedUser, userDeleted.ID)

	dataWriteAllUsersInDB, err := json.Marshal(allFriendsWithoutDeletedUser)
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

	response, err := json.Marshal(userDeleted.ID)
	if err != nil {
		handler.log.Printf("cannot marshal userDeleted")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write([]byte(response)); err != nil {
		handler.log.Printf("err = %v", err)
	}
}
