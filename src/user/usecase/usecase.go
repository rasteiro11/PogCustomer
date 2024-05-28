package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	pbCrypto "github.com/rasteiro11/PogCustomer/gen/proto/go/crypto"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/defaultrole"
	"github.com/rasteiro11/PogCustomer/src/role"
	"github.com/rasteiro11/PogCustomer/src/user"
	"github.com/rasteiro11/PogCustomer/src/user/repository"
	"github.com/rasteiro11/PogCustomer/src/userrole"
)

type (
	UsecaseOpt func(*usecase)
	usecase    struct {
		repository      user.Repository
		roleRepo        role.Repository
		userRoleRepo    userrole.Repository
		defaultRoleRepo defaultrole.Repository
		cryptoService   pbCrypto.CryptoServiceClient
		hash            hash.Hash
	}
)

var (
	ErrEmailTaken = errors.New("email already taken")
)

var _ user.Usecase = (*usecase)(nil)

func WithCryptoClient(client pbCrypto.CryptoServiceClient) UsecaseOpt {
	return func(u *usecase) {
		u.cryptoService = client
	}
}

func WithRepository(repository user.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.repository = repository
	}
}

func WithRoleRepository(repository role.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.roleRepo = repository
	}
}

func WithUserRoleRepository(repository userrole.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.userRoleRepo = repository
	}
}

func WithDefaultRoleRepository(repository defaultrole.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.defaultRoleRepo = repository
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

func (u *usecase) Register(ctx context.Context, req *models.RegisterUserRequest) (*models.RegisterUserResponse, error) {
	newRole := &models.Role{}

	if _, err := u.repository.FindOne(ctx, &models.User{Email: req.Email}); err != nil {
		fmt.Printf("ENTERED FIND ONE CONTEXT")
		if !errors.Is(err, repository.ErrRecordNotFound) {
			fmt.Printf("ERR RECORD NOT FOUND")
			return nil, err
		}

		if err = u.repository.Tx(ctx, func(ctx context.Context) error {

			user, err := u.repository.Create(ctx, &models.User{Document: req.Document, Email: req.Email, Password: hashPassword(u.hash, req.Password)})
			if err != nil {
				fmt.Printf("ERR CREATING")
				return err
			}

			defaultRole, err := u.defaultRoleRepo.First(ctx, &models.DefaultRole{})
			if err != nil {
				fmt.Printf("ERR DEFAULT ROLE")
				return err
			}

			fmt.Printf("DEFAULT ROLE: %+v", defaultRole)

			newRole.Name = defaultRole.Name
			newRole, err = u.roleRepo.Store(ctx, newRole)
			if err != nil {
				return err
			}

			if _, err := u.userRoleRepo.Store(ctx, &models.UserRole{
				UserID: user.ID,
				RoleID: newRole.ID,
			}); err != nil {
				return err
			}

			// if _, err := u.cryptoService.RegisterUserWallet(ctx, &pbCrypto.RegisterUserWalletRequest{
			// 	UserId:  int32(user.ID),
			// 	Wallet:  req.Wallet,
			// 	Network: "Ethereum",
			// }); err != nil {
			// 	return err
			// }

			return nil
		}); err != nil {
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

	return &models.RegisterUserResponse{
		Token:     loginCreds.Token,
		ExpiresAt: loginCreds.ExpiresAt,
		Role:      newRole,
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

	userRole, err := u.userRoleRepo.Find(ctx, &models.UserRole{
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}

	role, err := u.roleRepo.Find(ctx, &models.Role{
		ID: userRole.RoleID,
	})
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
		Role:      role,
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
