package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
)

type TokenParser func(tokenString *string) (interface{}, error)

type DataInjector func(c *gin.Context, data interface{}) error

func BearerAuth(parse TokenParser, inject DataInjector) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := error(nil)
		defer func() {
			if err != nil {
				stdrsp.ResponseResult(c, stdrsp.Response(stdrsp.AuthFailed).Msg(stdrsp.AuthErrMsg))
				c.Abort()
			}
		}()
		tokenString, err := getBearerTokenString(c)
		if err != nil {
			return
		}
		data, err := parse(tokenString)
		if err != nil {
			return
		}
		err = inject(c, data)
		if err != nil {
			return
		}
	}
}

func getBearerTokenString(c *gin.Context) (*string, error) {
	header := c.GetHeader("Authorization")
	if len(header) < 7 || header[0:7] != "Bearer " {
		return nil, errors.New("a valid bearer token is required")
	}
	tokenString := header[7:]
	return &tokenString, nil
}
