package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/config"
)

func NewString(claims map[string]interface{}) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(claims))
	identityJwtString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return nil, err
	}
	return &identityJwtString, nil
}

func parseKeySelector(_ *jwt.Token) (interface{}, error) {
	return config.JwtSecret, nil
}

var parserOption = jwt.WithValidMethods([]string{"HS512"})

func Parse(jwtString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, parseKeySelector, parserOption)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, err
	} else {
		return &claims, nil
	}
}

func ExtractUUid(claims *jwt.MapClaims, field string) (*uuid.UUID, error) {
	fieldStringValue, ok := (*claims)[field].(string)
	if !ok {
		return nil, errors.New("extracting field error")
	}

	fieldUuidValue, err := uuid.Parse(fieldStringValue)
	if err != nil {
		return nil, err
	}

	return &fieldUuidValue, nil
}
