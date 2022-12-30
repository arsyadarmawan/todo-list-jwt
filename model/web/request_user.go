package web

type UserCreateRequest struct {
	Name     string `json:"name"`
	Password string `validate:"required,min=8,max=20" json:"password"`
	Username string `validate:"required,min=1,max=20" json:"username"`
	Email    string `validate:"required,min=1" json:"email"`
}

type LoginRequest struct {
	Password string `validate:"required,min=8,max=20" json:"password"`
	Email    string `validate:"required,min=1" json:"email"`
}
