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

func GetAllUser() (response []domain.GetUserResponse, err error) {
	var userRows []domain.Users
	var responseRow domain.GetUserResponse
	err = outbound.DatabaseDriver.Find(&userRows).Error
	if err != nil {
		return
	}

	for _, row := range userRows {
		responseRow = domain.GetUserResponse{}
		responseRow.ID = int(row.ID)
		responseRow.Role = row.Role
		responseRow.Email = row.Email
		responseRow.Name = row.Username
		response = append(response, responseRow)
	}

	return
}

func GetUser(userId int, claims domain.JwtValidate) (response domain.GetUserResponse, err error) {
	var user domain.Users

	err = outbound.DatabaseDriver.First(&user, userId).Error
	if err != nil {
		return
	}

	if claims.Claims["role"].(string) == "customer" && float64(userId) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this user")
		return
	}

	response.ID = userId
	response.Email = user.Email
	response.Name = user.Username
	response.Role = user.Role
	return
}

func UpdateUser(request domain.UpdateUserRequest, claims domain.JwtValidate) (response domain.UserResponse, err error) {
	var user domain.Users

	if claims.Claims["role"].(string) == "customer" && float64(request.ID) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this user")
		return
	}

	err = outbound.DatabaseDriver.Where("id = ?", request.ID).First(&user).Error
	if err != nil {
		return
	}

	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Name != "" {
		user.Username = request.Name
	}
	err = outbound.DatabaseDriver.Save(&user).Error
	if err != nil {
		return
	}

	response.ID = uint(request.ID)
	response.Status = "User Updated Succesfully!"
	return
}

func DeleteUser(userId int, claims domain.JwtValidate) (response domain.UserResponse, err error) {
	var user domain.Users

	if claims.Claims["role"].(string) == "customer" && float64(userId) != (claims.Claims["user_id"].(float64)) {
		err = errors.New("You are not authorized to access this user")
		return
	}

	err = outbound.DatabaseDriver.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return
	}

	err = outbound.DatabaseDriver.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return
	}

	response.ID = uint(userId)
	response.Status = "User is Successfully Deleted."
	return
}
