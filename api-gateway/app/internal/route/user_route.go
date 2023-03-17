package router

import (
	userrepo "api-gateway/app/domain/usercases/user/repo"
	"api-gateway/app/middleware"
	userHandler "api-gateway/app/transport/user/handler"
)

func setupUserRoute(v1 fiber.Router) {
	groupUserHandler := userHandler.UserHandler{
		UserDomain: userrepo.User,
	}

	groupUser := v1.Group("/user/me", middleware.Auth.RequireLogin())
	{
		GET(groupUser, "", groupUserHandler.GetMe)
	}
}
