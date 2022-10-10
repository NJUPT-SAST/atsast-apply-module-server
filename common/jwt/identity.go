package jwt

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	identityValidTime = 7 * 24 * time.Hour
)

type Identity struct {
	Uid *uuid.UUID
}

func NewIdentityString(identity *Identity) (*string, error) {
	return NewString(map[string]any{
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(identityValidTime).Unix(),
		"uid": identity.Uid.String(),
	})
}

func ParseIdentityJwtString(identityJwtString *string) (*Identity, error) {
	claims, err := Parse(*identityJwtString)
	if err != nil {
		return nil, err
	}

	uid, err := extractUUid(claims, "uid")
	if err != nil {
		return nil, err
	}

	identity := Identity{
		Uid: uid,
	}
	return &identity, nil
}

func InjectIdentity(c *gin.Context, identity *Identity) error {
	c.Set("identity", identity)
	return nil
}

func MustExtractIdentity(c *gin.Context) *Identity {
	return c.MustGet("identity").(*Identity)
}
