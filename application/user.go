package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/outbound"
	"banking-system-backend/util"
	"time"
)

func Register(request domain.UserRegisterRequest) (response domain.UserRegisterResponse, err error) {

	var hashedPassword domain.HashPassword

	hashedPassword, err = util.HashPassword(request.Password)
	if err != nil {
		return
	}

	userRegister := domain.Users{
		Username:  request.Username,
		Email:     request.Email,
		Password:  hashedPassword.EncryptPassword,
		Role:      request.Role,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err = outbound.DatabaseDriver.Create(&userRegister).Error

	if err != nil {
		return
	}

	response.Status = "User successfully Registered."

	return
}
