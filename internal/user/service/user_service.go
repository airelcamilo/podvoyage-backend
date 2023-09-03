package userService

import (
	"errors"
	"strconv"
	"time"

	"github.com/airelcamilo/podvoyage-backend/internal/user/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SecretKey = "secret"

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) New(db *gorm.DB) UserService {
	return UserService{db}
}

func (s *UserService) Register(request *model.RegisterRequest) (model.UserResponse, error) {
	var user, temp model.User
	var userResponse model.UserResponse
	if result := s.DB.Where("username = ?", request.Username).First(&temp); result.Error == nil {
		return userResponse, errors.New("username already taken")
	}
	
	if result := s.DB.Where("email = ?", request.Email).First(&temp); result.Error == nil {
		return userResponse, errors.New("email already taken")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	user = model.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: password,
	}
	if result := s.DB.Create(&user); result.Error != nil {
		return userResponse, result.Error
	}

	userResponse, err := s.Login(&model.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	return userResponse, err
}

func (s *UserService) Login(request *model.LoginRequest) (model.UserResponse, error) {
	var user model.User
	var userResponse model.UserResponse
	if result := s.DB.Where("email = ?", request.Email).First(&user); result.Error != nil {
		return userResponse, errors.New("incorrect email")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
		return userResponse, errors.New("incorrect password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return userResponse, errors.New("could not login")
	}

	if result := s.DB.Create(&model.Session{
		Token: token,
		User:  user,
	}); result.Error != nil {
		return userResponse, result.Error
	}
	userResponse = model.UserResponse{
		Token: token,
		User:  user,
	}
	return userResponse, nil
}

func (s *UserService) Validate(token string) (model.User, error) {
	var user model.User
	var session model.Session
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return user, errors.New("unauthorized")
	}
	if result := s.DB.Preload("User").Where("token = ?", token).First(&session); result.Error != nil {
		return user, errors.New("user not found")
	}

	claims := jwtToken.Claims.(*jwt.RegisteredClaims)
	if strconv.Itoa(session.User.Id) != claims.Issuer {
		return user, errors.New("user not found")
	}
	return session.User, nil
}

func (s *UserService) Logout(token string) (string, error) {
	var session model.Session
	if result := s.DB.Where("token = ?", token).First(&session).Delete(&session); result.Error != nil {
		return "", errors.New("user not found")
	}
	return "logout", nil
}
