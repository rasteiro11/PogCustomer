package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/user"
	"github.com/rasteiro11/PogCustomer/src/user/repository"
	"hash"
	"os"
	"time"
)

type (
	UsecaseOpt func(*usecase)
	usecase    struct {
		repository user.Repository
		hash       hash.Hash
	}
)

var (
	ErrEmailTaken = errors.New("email already taken")
)

var _ user.Usecase = (*usecase)(nil)

func WithRepository(repository user.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.repository = repository
	}
}

func WithHash(hash hash.Hash) UsecaseOpt {
	return func(u *usecase) {
		u.hash = hash
	}
}

func NewUsecase(opts ...UsecaseOpt) user.Usecase {
	u := &usecase{}

	for _, opt := range opts {
		opt(u)
	}

	if u.hash == nil {
		u.hash = sha1.New()
	}

	return u
}

type claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func hashPassword(hasher hash.Hash, password string) string {
	encodedPassword := base64.RawStdEncoding.EncodeToString(hasher.Sum([]byte(password)))
	hasher.Reset()
	return encodedPassword
}

func (u *usecase) Register(ctx context.Context, req *models.User) (*models.RegisterResponse, error) {
	_, err := u.repository.FindOne(ctx, &models.User{Email: req.Email, Document: req.Document})
	if err != nil {
		if !errors.Is(err, repository.ErrRecordNotFound) {
			return nil, err
		}
		_, err := u.repository.Create(ctx, &models.User{Document: req.Document, Email: req.Email, Password: hashPassword(u.hash, req.Password)})
		if err != nil {
			return nil, err
		}
	}

	loginCreds, err := u.Login(ctx, &models.User{
		Document: req.Document,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &models.RegisterResponse{
		Token:     loginCreds.Token,
		ExpiresAt: loginCreds.ExpiresAt,
	}, nil
}

func (u *usecase) Login(ctx context.Context, req *models.User) (*models.LoginResponse, error) {
	expiresAt := time.Now().Add(time.Minute * 15)
	user, err := u.repository.FindOne(ctx, &models.User{Document: req.Document, Email: req.Email, Password: hashPassword(u.hash, req.Password)})
	if err != nil {
		return nil, err
	}

	claims := &claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
	}, nil
}

func (u *usecase) ChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (*models.ChangePasswordResponse, error) {
	user, err := u.repository.FindOne(ctx, &models.User{Email: req.Email, Password: hashPassword(u.hash, req.Password)})
	if err != nil {
		return nil, err
	}

	user.Password = hashPassword(u.hash, req.NewPassword)
	_, err = u.repository.UpdateById(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.ChangePasswordResponse{}, nil
}

func (u *usecase) extractClaims(ctx context.Context, token string) (*claims, error) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return claims, err
}

func (u *usecase) FindOne(ctx context.Context, req *models.User) (*models.User, error) {
	user, err := u.repository.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
