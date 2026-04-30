package files_transport_http

import (
	"net/http"

	core_http_request "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/request"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

type DeleteFileRequest struct {
	ID int `json:"id" validate:"required,min=0"`
}

func (h *FilesHTTPHandler) DeleteFile(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := FilesHTTPTransportLogger(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	var request DeleteFileRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode and validate HTTP request")
		return
	}

	err := h.usersService.DeleteFile(ctx, request.ID)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to delete user")
		return
	}

	responseHandler.StatusCodeResponse(http.StatusCreated)
}
