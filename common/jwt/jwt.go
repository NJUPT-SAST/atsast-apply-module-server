package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/njupt-sast/atsast-apply-module-server/config"
)

var (
	parserOption	 = jwt.WithValidMethods([]string{"HS512"})
	parseKeySelector = func(_ *jwt.Token) (any, error) {
		return config.JwtSecret, nil
	}
)

func NewString(claims map[string]any) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(claims))
	identityJwtString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return nil, err
	}

	return &identityJwtString, nil
}

func Parse(jwtString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, parseKeySelector, parserOption)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return &claims, nil
}

func extractUUid(claims *jwt.MapClaims, field string) (*uuid.UUID, error) {
	fieldStringValue, ok := (*claims)[field].(string)
	if !ok || fieldStringValue == "" {
		return nil, errors.New("extracting field error")
	}

	fieldUuidValue, err := uuid.Parse(fieldStringValue)
	if err != nil {
		return nil, err
	}

	return &fieldUuidValue, nil
}
