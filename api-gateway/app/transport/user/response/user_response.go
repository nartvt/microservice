package response

import "api-gateway/app/domain/usercases/user/repo"

func NewUserResponse(userRepo repo.UserRepo) *UserResponse {
	return &UserResponse{
		Id:       userRepo.Id,
		FullName: userRepo.FullName,
	}
}
