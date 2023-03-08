package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"health-service/app/domain/usercases/user/repo"
)

func GetCurrentUser(userName string) (*repo.UserRepo, error) {
	return repo.User.GetUserByUserName(userName)
}

func IsAdmin(ctx *fiber.Ctx) bool {
	role := ctx.Get("role")
	roleId, _ := strconv.Atoi(role)
	return roleId == 1 || roleId == 2
}
