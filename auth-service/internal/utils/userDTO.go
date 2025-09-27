package utils

import m "auth-service/internal/models"

func ToUserResponse(u m.User) *m.UserResponse {
	return &m.UserResponse{
		ID:          u.ID.Hex(),
		FullName:    u.Name + " " + u.Surname,
		Gender:      u.Gender,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}

}
