package pkg

import (
	"errors"
	"time"

	"github.com/abdallahelassal/Store/config"
	"github.com/abdallahelassal/Store/internal/modules/user/domain"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string	    `json:"user_id"`
	Email  string      	`json:"email"`
	Role   domain.Role 	`json:"role"`
	jwt.RegisteredClaims
}

type JWTService struct{
	secretKey		string
	exprictionTime 	time.Duration
	refreshExpiry 	time.Duration
}

func NewJWTService(sec config.JWTConfig, expMinutes int, refreshDay int) *JWTService{
	return &JWTService{
		secretKey: sec.Secret,
		exprictionTime: time.Duration(expMinutes) * time.Minute,
		refreshExpiry: time.Duration(refreshDay)  * 24 * time.Hour,
	}
}

func (j *JWTService) GenerateToken(userID uuid.UUID, email, role string)(string, error){
	claims := Claims{
		UserID: userID.String(),
		Email: email,
		Role: domain.Role(role),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "store-api",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.exprictionTime)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) GenerateRefreshToken(userID uuid.UUID)(string, error){
	claims := &jwt.RegisteredClaims{
		Issuer: "store-api",
		Subject: userID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshExpiry)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(tokenString string)(*Claims, error){
	token , err := jwt.ParseWithClaims(tokenString,&Claims{},func(t *jwt.Token) (interface{}, error) {
		if _, ok :=t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil , errors.New("unexpected signing method")
		}
		return []byte(j.secretKey),nil 
	})
	if err != nil {
		return nil , err
	}
	if claims , ok := token.Claims.(*Claims); ok && token.Valid{
		return claims, nil 
	}
	return nil ,errors.New("invalid token")
}