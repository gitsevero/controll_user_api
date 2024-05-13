package UserModel

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20,containsany=!@#$%&*"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int    `json:"age" binding:"required,min=6,max=120"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age" `
}
