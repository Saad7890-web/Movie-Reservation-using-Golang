package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type TokenManager struct {
	secret []byte
	expiry time.Duration
}

func NewTokenManager(secret string, expHours int) *TokenManager {
	return &TokenManager{
		secret: []byte(secret),
		expiry: time.Duration(expHours) * time.Hour,
	}
}

type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

func (tm *TokenManager) Generate(c Claims) (string, error) {
	c.Exp = time.Now().Add(tm.expiry).Unix()

	header := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"alg":"HS256","typ":"JWT"}`),
	)

	payloadBytes, _ := json.Marshal(c)
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)

	message := header + "." + payload
	return message + "." + tm.sign(message), nil
}

func (tm *TokenManager) Parse(token string) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token")
	}

	message := parts[0] + "." + parts[1]
	if tm.sign(message) != parts[2] {
		return nil, errors.New("invalid signature")
	}

	data, _ := base64.RawURLEncoding.DecodeString(parts[1])
	var c Claims
	json.Unmarshal(data, &c)

	if time.Now().Unix() > c.Exp {
		return nil, errors.New("token expired")
	}

	return &c, nil
}

func (tm *TokenManager) sign(data string) string {
	h := hmac.New(sha256.New, tm.secret)
	h.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
