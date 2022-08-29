package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/isaquesr/users-test-golang/auth"
	"github.com/isaquesr/users-test-golang/domain"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *domain.User `json:"name"`
}

type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTL time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, u *domain.User) error {
	pwd := sha1.New()
	pwd.Write([]byte(u.Password))
	pwd.Write([]byte(a.hashSalt))

	user := &domain.User{
		Name:     u.Name,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
		Address:  u.Address,
		Age:      u.Age,
		Email:    u.Email,
	}

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, name, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, name, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*domain.User, error) {
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
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
