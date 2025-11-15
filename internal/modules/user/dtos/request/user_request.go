package request


type RegisterRequest struct{
	Name 		string		`json:"name" binding:"required,min=3,max=30"`
	Email		string		`json:"email" binding:"required,min=3,email"`
	Password	string		`json:"password" binding:"required,min=6,max=100"`	
}

type LoginRequest struct{
	Email		string		`json:"email" binding:"required,email,min=3"`
	Password	string		`json:"password" binding:"required"`
}

type UpdateProfileRequest struct{
	Name		string		`json:"name" binding:"required,min=3,max=30"`
	Email		string		`json:"email" binding:"required,email"`
	Password	string		`json:"password" binding:"required"`
}

type RefreshTokenRequest struct{
	RefreshToken string 	`json:"refresh_token" binding:"required"`
}