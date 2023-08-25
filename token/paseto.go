package token

import (
	"time"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	secretKey []byte
}

func NewPasetoMaker(secretKey string) (Maker, error) {
	return &PasetoMaker{secretKey: []byte(secretKey)}, nil
}

// VerifyToken(token string) (*Payload, error)

func (p *PasetoMaker) GenerateToken(account_id uint, username string, duration time.Duration) (string, error) {
	payload := NewPayload(account_id, username, duration)
	v2 := paseto.NewV2()

	token, err := v2.Encrypt(p.secretKey, payload, nil)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (p *PasetoMaker) VerifyToken(tokenString string) (*Payload, error) {
	var payload Payload
	v2 := paseto.NewV2()

	err := v2.Decrypt(tokenString, p.secretKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
