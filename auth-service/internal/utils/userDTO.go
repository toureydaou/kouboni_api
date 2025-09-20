package utils

import m "auth-service/internal/models"

func ToUserResponse(u m.User) m.UserResponse {
	return m.UserResponse{
		ID:          u.ID.Hex(),
		FullName:    u.FirstName + " " + u.Surname,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}

}
