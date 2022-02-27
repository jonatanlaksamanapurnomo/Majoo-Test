package util

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"majoo-be-test/enums"
	"majoo-be-test/models"
	"net/http"
	"os"
	"time"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Id uint
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(id uint) (signedToken string, err error) {
	claims := &JwtClaim{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

//ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return

}

func GetValueEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

//generate jwt function
func GenerateJWT(w http.ResponseWriter, user models.User) string {
	jwtWrapper := JwtWrapper{
		SecretKey:       GetValueEnv("JWT_SECRET"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	token, tokenError := jwtWrapper.GenerateToken(user.ID)
	if tokenError != nil {
		http.Error(w, enums.ERROR, enums.BAD_REQUEST)
	}
	return token
}

func GetJWTValue(signedToken string) (claims *JwtClaim, err error) {
	jwtWrapper := JwtWrapper{
		SecretKey:       GetValueEnv("JWT_SECRET"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	data, tokenError := jwtWrapper.ValidateToken(signedToken)
	return data, tokenError

}
