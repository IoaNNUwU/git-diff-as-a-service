package files_transport_http

import (
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_http_request "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/request"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

type CreateFileRequest struct {
	FileName string  `json:"file_name" validate:"required,min=3,max=100"`
	OwnerID  int     `json:"owner_id" validate:"required,min=0"`
	Content  *string `json:"content" validate:"omitempty,min=10,max=20"`
}

type CreateFileResponse struct {
	ID       int    `json:"id"`
	Version  int    `json:"version"`
	FileName string `json:"file_name"`
}

func (h *FilesHTTPHandler) CreateFile(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := FilesHTTPTransportLogger(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	var request CreateFileRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(request)

	user, err := h.usersService.CreateFile(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to create file")
		return
	}

	response := dtoFromDomain(user)

	responseHandler.JSONResponse(&response, http.StatusCreated)
}

func domainFromDTO(dto CreateFileRequest) domain.File {

	var fileContent string
	if dto.Content != nil {
		fileContent = *dto.Content
	} else {
		fileContent = ""
	}

	return domain.NewFileUninitialized(dto.FileName, dto.OwnerID, fileContent)
}

func dtoFromDomain(file domain.File) CreateFileResponse {
	return CreateFileResponse{
		ID:       file.ID,
		Version:  file.Version,
		FileName: file.FileName,
	}
}
