package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/isaquesr/users-test-golang/domain"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/isaquesr/users-test-golang/auth"
)

type AuthClaims struct {
	jwt.StandardClaims
	Login *domain.Login `json:"login"`
}

type AuthUseCase struct {
	loginRepo      auth.LoginRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	loginRepo auth.LoginRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTL time.Duration) *AuthUseCase {
	return &AuthUseCase{
		loginRepo:      loginRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &domain.Login{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.loginRepo.CreateLogin(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.loginRepo.GetLogin(ctx, username, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		Login: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*domain.Login, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.Login, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
