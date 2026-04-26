package users_transport_http

type UsersHTTPHandler struct {
	usersService UsersService
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

type UsersService interface {

}