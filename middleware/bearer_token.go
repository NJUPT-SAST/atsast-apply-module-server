package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenParser[T any] func(*string) (T, error)

type DataInjector[T any] func(*gin.Context, T) error

func BearerTokenAuth[T any](
	parser TokenParser[T],
	injector DataInjector[T],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenSting, err := getBearerTokenString(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		data, err := parser(tokenSting)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err = injector(c, data); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func getBearerTokenString(c *gin.Context) (*string, error) {
	header := c.GetHeader("Authorization")
	if len(header) < 7 || header[0:7] != "Bearer " {
		return nil, errors.New("not a bearer token")
	}

	tokenString := header[7:]
	return &tokenString, nil
}
