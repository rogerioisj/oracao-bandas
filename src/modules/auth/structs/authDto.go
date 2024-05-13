package auth

type LoginDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}
