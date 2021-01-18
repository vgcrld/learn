package pwtools

import (
	b64 "encoding/base64"
	"errors"
)

type Credentials struct {
	User, plain, Encrypt string
}

func NewCredentials(user, pw string) Credentials {
	c := Credentials{}
	c.Encrypt = b64.StdEncoding.EncodeToString([]byte(pw))
	c.plain = pw
	c.User = user
	return c
}

func (c Credentials) GetPlain(u string) (string, error) {
	if u == c.User {
		return c.plain, nil
	} else {
		return "", errors.New("that user is not the correct user")
	}
}

func (c Credentials) IsValid() bool {
	val, _ := b64.StdEncoding.DecodeString(c.Encrypt)
	if c.plain == string(val) {
		return true
	}
	return false
}
