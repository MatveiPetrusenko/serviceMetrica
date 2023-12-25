package common

import (
	"fmt"
	"serviceMetrica/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// lifeTime minimum lifetime
const lifeTime = time.Minute * 10

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// CreateJwtToken create token
func CreateJwtToken(login string, role string) (string, error) {
	duration := config.New().Token.ExpiresAt
	if duration <= 0 {
		duration = lifeTime
	}

	// Payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{ // MethodHS512
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.New().Token.Issuer,
			Subject:   login,
			Audience:  jwt.ClaimStrings{config.New().Token.Issuer},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	})

	// SIGNATURE = header + payload
	tokenAsString, err := token.SignedString([]byte(config.New().Token.Secret))

	return tokenAsString, err
}

// ParseJwtToken parses jwt-token and checks it`s validity
func ParseJwtToken(tokenAsString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenAsString, new(Claims), func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v\n", token.Header["alg"])
		}

		return []byte(config.New().Token.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		for _, aud := range claims.Audience {
			if aud == config.New().Token.Issuer { // Сравниваем с издателем
				return claims, nil
			}
		}
		return nil, fmt.Errorf("this key is not intended for this service")
	}
	return nil, fmt.Errorf("invalid JWT-token")
}
