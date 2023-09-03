package utils

import (
	"net/http"
	"strings"

	"github.com/airelcamilo/podvoyage-backend/internal/user/model"
	us "github.com/airelcamilo/podvoyage-backend/internal/user/service"
	"gorm.io/gorm"
)

type AuthUtils struct {
	userService *us.UserService
}

func (a *AuthUtils) New(db *gorm.DB) AuthUtils {
	var userService us.UserService
	userService = userService.New(db)
	return AuthUtils{&userService}
}

func (a *AuthUtils) GetUser(w http.ResponseWriter, r *http.Request) (model.User, error) {
	var user model.User
	reqToken := r.Header.Get("Authorization")
	reqToken = strings.Split(reqToken, "Bearer ")[1]
	user, err := a.userService.Validate(reqToken)
	if err != nil {
		return user, err
	}
	return user, nil
}
