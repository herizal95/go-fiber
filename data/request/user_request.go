package request

type CreateUserRequest struct {
	Name     string `validate:"required,min=2,max=100" json:"name"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type UpdateUserRequest struct {
	Uid      string `validate:"required"`
	Name     string `validate:"required,min=2,max=100" json:"name"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}
