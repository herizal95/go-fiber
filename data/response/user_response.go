package response

import "github.com/google/uuid"

type UserResponse struct {
	Uid      uuid.UUID `json:"uid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
