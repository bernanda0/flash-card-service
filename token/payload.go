package token

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// struct of payload
type Payload struct {
	AccountID uint      `json:"account_id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// GetAudience implements jwt.Claims.
func (*Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}

// GetExpirationTime implements jwt.Claims.
func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	exp := jwt.NewNumericDate(p.ExpiredAt)
	return exp, nil
}

// GetIssuedAt implements jwt.Claims.
func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	issuedAt := jwt.NewNumericDate(p.IssuedAt)
	return issuedAt, nil
}

// GetIssuer implements jwt.Claims.
func (*Payload) GetIssuer() (string, error) {
	return "your-issuer-name", nil
}

// GetNotBefore implements jwt.Claims.
func (*Payload) GetNotBefore() (*jwt.NumericDate, error) {
	nbf := jwt.NewNumericDate(time.Now())
	return nbf, nil
}

// GetSubject implements jwt.Claims.
func (p *Payload) GetSubject() (string, error) {
	return p.Username, nil
}

func NewPayload(account_id uint, username string, duration time.Duration) *Payload {
	return &Payload{
		AccountID: account_id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) TimeValid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
