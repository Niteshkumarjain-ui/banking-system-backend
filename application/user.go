package application

import (
	"banking-system-backend/domain"
	"banking-system-backend/outbound"
	"banking-system-backend/util"
	"errors"
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

func Login(request domain.UserLoginRequest) (response domain.UserLoginResponse, err error) {
	var user domain.Users
	var hashPassword domain.CheckHashPassword
	var Jwt domain.JwtGenerate

	if request.Email != "" {
		err = outbound.DatabaseDriver.Where("email = ?", request.Email).First(&user).Error
	} else {
		err = outbound.DatabaseDriver.Where("username = ?", request.Username).First(&user).Error
	}

	if err != nil {
		return
	}

	hashPassword = util.CheckPasswordHash(request.Password, user.Password)

	if !hashPassword.ValidPassword {
		err = errors.New("Invalid Password")
		return
	}

	Jwt, err = util.GenerateJWT(user)

	if err != nil {
		err = errors.New("Failed to Generate jwt Token")
		return
	}

	response.Message = "User Successfully Logged in."
	response.Token = Jwt.Token

	return
}
