package mapper

import (
	"github.com/Kei-K23/mock-commerce-api/dto"
	"github.com/Kei-K23/mock-commerce-api/models"
)

func MatchUserRequestToUser(req dto.UserRequest) *models.User {
	return &models.User{
		Username:      req.Username,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Password:      req.Password,
		Phone:         req.Phone,
		City:          req.City,
		Street:        req.Street,
		AddressNumber: req.AddressNumber,
		ZipCode:       req.ZipCode,
		Lat:           req.Lat,
		Long:          req.Long,
	}
}

func MatchUserToUserRequest(user models.User) *dto.UserRequest {
	return &dto.UserRequest{
		ID:            user.ID,
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Password:      user.Password,
		Phone:         user.Phone,
		City:          user.City,
		Street:        user.Street,
		AddressNumber: user.AddressNumber,
		ZipCode:       user.ZipCode,
		Lat:           user.Lat,
		Long:          user.Long,
	}
}
