package response

type UserResponse struct{
	Name	string		`json:"name"`
	Email	string		`json:"email"`
	Role	string		`json:"role"`
}

type AuthResponse struct{
	AccessToken		string		`json:"access_token"`
	RefreshToken	string		`json:"refresh_token"`
	TokenType		string		`json:"token_type"`
	ExpiresIn		int			`json:"expires_in"`
	User 			*UserResponse `json:"user"`
}

type LoginResponse struct{
	AccessToken		string		`json:"access_token"`
	RefreshToken	string		`json:"refresh_token"`
	TokenType		string		`json:"token_type"`
	ExpiresIn		int			`json:"expires_in"`
	User 			*UserResponse `json:"user"`
}