package token

import "time"

// this is an functional intercafe
type Maker interface {
	// Token for specifik username
	GenerateToken(account_id uint, username string, duration time.Duration) (string, *Payload, error)
	// if success return the payload
	VerifyToken(token string) (*Payload, error)
}
