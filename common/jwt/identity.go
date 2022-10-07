package jwt

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Identity struct {
	Uid *uuid.UUID
}

var identityValidTime = 7 * 24 * time.Hour

func NewIdentityJwtString(identity *Identity) (*string, error) {
	jwtString, err := NewString(map[string]interface{}{
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(identityValidTime).Unix(),
		"uid": identity.Uid.String(),
	})
	if err != nil {
		return nil, err
	}

	return jwtString, nil
}

func ParseIdentityJwtString(identityJwtString *string) (interface{}, error) {
	claims, err := Parse(*identityJwtString)
	if err != nil {
		return nil, err
	}

	uid, err := ExtractUUid(claims, "uid")
	if err != nil {
		return nil, err
	}

	identity := Identity{
		Uid: uid,
	}
	return &identity, nil
}

func InjectIdentity(c *gin.Context, identity interface{}) error {
	c.Set("identity", identity)
	return nil
}

func ExtractIdentity(c *gin.Context) *Identity {
	return c.MustGet("identity").(*Identity)
}
