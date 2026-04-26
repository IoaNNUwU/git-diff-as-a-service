package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	FullName string  `json:"full_name"`
	Email    *string `json:"email"`
}

type CreateUserResponce struct {
	ID       int     `json:"id"`
	Version  int     `json:"version"`
	FullName string  `json:"full_name"`
	Email    *string `json:"email"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		
	}
}
