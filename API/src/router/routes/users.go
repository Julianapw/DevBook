package routes

import "net/http"

var controller = struct {
	CreateUser http.HandlerFunc
	GetUsers   http.HandlerFunc
	GetUser    http.HandlerFunc
	UpdateUser http.HandlerFunc
	DeleteUser http.HandlerFunc
}{
	CreateUser: func(w http.ResponseWriter, r *http.Request) {},
	GetUsers:   func(w http.ResponseWriter, r *http.Request) {},
	GetUser:    func(w http.ResponseWriter, r *http.Request) {},
	UpdateUser: func(w http.ResponseWriter, r *http.Request) {},
	DeleteUser: func(w http.ResponseWriter, r *http.Request) {},
}

var usersRoutes = []Route{
	{
		URI:          "/users",
		Method:       "POST",
		Function:     controller.CreateUser,
		AuthRequired: false,
	},

	{
		URI:          "/users",
		Method:       "GET",
		Function:     controller.GetUsers,
		AuthRequired: true,
	},

	{
		URI:          "/users/{id}",
		Method:       "GET",
		Function:     controller.GetUser,
		AuthRequired: true,
	},

	{
		URI:          "/users/{id}",
		Method:       "PUT",
		Function:     controller.UpdateUser,
		AuthRequired: true,
	},

	{
		URI:          "/users/{id}",
		Method:       "DELETE",
		Function:     controller.DeleteUser,
		AuthRequired: true,
	},
}
