package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	NewJWT(id int, ttl time.Duration) (string, error)
	Parse(accessToken string) (int, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signinKey string
}

func NewManager(signinKey string) (TokenManager, error) {
	if signinKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{
		signinKey: signinKey,
	}, nil
}

func (m *Manager) NewJWT(id int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   strconv.Itoa(id),
	})

	return token.SignedString([]byte(m.signinKey))
}

func (m *Manager) Parse(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(m.signinKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims := token.Claims.(jwt.MapClaims)
	adminIdToFloat64, ok := claims["admin_id"].(float64)
	if !ok {
		return 0, errors.New("admin_id is not a number")
	}
	adminID := int(adminIdToFloat64)
	return adminID, nil

}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil

}
