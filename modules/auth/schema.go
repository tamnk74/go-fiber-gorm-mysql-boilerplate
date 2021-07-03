package auth

type LoginSchema struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterSchema struct {
	Name     string `form:"name" json:"name" binding:"required" validate:"required,min=3,max=32"`
	Email    string `form:"email" json:"email" binding:"required" validate:"required,email,min=6,max=32"`
	Password string `form:"password" json:"password" binding:"required" validate:"required,min=8,max=32"`
}
