package handler

import (
	"context"
	"go-4-digi-micro/auth/model"
	pb "go-4-digi-micro/auth/proto"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type authImplement struct {
	pb.UnimplementedAuthServer
	db         *gorm.DB
	signingKey []byte
}

func NewAuth(db *gorm.DB, signingKey []byte) pb.AuthServer {
	return &authImplement{
		db:         db,
		signingKey: signingKey,
	}
}

func (a *authImplement) Login(c context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	// validate
	auth := model.Auth{}
	if err := a.db.Where("username = ?", req.Username).
		First(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.Unauthenticated, "Login not valid")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	// validate
	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Login not valid")
	}

	// login valid
	token, err := a.createJWT(&auth)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AuthLoginResponse{
		Token: token,
	}, nil
}

func (a *authImplement) Validate(c context.Context, req *pb.AuthValidateRequest) (*pb.AuthValidateResponse, error) {
	// Parse the token
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(a.signingKey), nil
	})

	if err != nil || !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Set token to response
	response := &pb.AuthValidateResponse{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if authID, ok := claims["auth_id"].(float64); ok {
			response.AuthId = int64(authID)
		}
		if accountID, ok := claims["account_id"].(float64); ok {
			response.AccountId = int64(accountID)
		}
		if username, ok := claims["username"].(string); ok {
			response.Username = username
		}
	} else {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	return response, nil
}

func (a *authImplement) Upsert(c context.Context, req *pb.AuthUpsertRequest) (*emptypb.Empty, error) {
	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// prepare auth data to save
	auth := model.Auth{
		AccountID: req.AccountId,
		Username:  req.Username,
		Password:  string(hashed),
	}

	// upsert auth
	result := a.db.Clauses(
		clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"username", "password"}),
			Columns:   []clause.Column{{Name: "account_id"}},
		}).Create(&auth)
	if result.Error != nil {
		return nil, status.Error(codes.InvalidArgument, result.Error.Error())
	}

	return &emptypb.Empty{}, nil
}

func (a *authImplement) createJWT(auth *model.Auth) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["auth_id"] = auth.AuthID
	claims["account_id"] = auth.AccountID
	claims["username"] = auth.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires in 72 hours

	// encode
	tokenString, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
