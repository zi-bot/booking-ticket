package utility

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type TokenDetail struct {
	UserID   float64 `json:"user_id"`
	UserName string  `json:"user_name"`
	Name     string  `json:"name"`
	Exp      float64 `json:"exp"`
}

func ExtractTokenMetadata(ctx context.Context) (*TokenDetail, error) {
	token, ok := ctx.Value("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("token not found")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims not found")
	}
	detail := TokenDetail{
		UserID:   claims["user_id"].(float64),
		UserName: claims["user_name"].(string),
		Name:     claims["name"].(string),
		Exp:      claims["exp"].(float64),
	}
	return &detail, nil
}
