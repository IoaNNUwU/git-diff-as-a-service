package users_transport_http

import (
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	core_http_request "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/request"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName string  `json:"full_name" validate:"required,min=3,max=100"`
	Email    *string `json:"email" validate:"omitempty,min=10,max=20"`
}

type CreateUserResponse struct {
	ID       int     `json:"id"`
	Version  int     `json:"version"`
	FullName string  `json:"full_name"`
	Email    *string `json:"email"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	log.Debug("invoke CreateUser")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(request)

	user, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to create user")
	}

	response := dtoFromDomain(user)

	responseHandler.JSONResponse(&response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.Email)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:       user.ID,
		Version:  user.Version,
		FullName: user.FullName,
		Email:    user.Email,
	}
}
