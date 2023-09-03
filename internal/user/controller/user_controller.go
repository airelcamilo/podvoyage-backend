package usercontroller

import (
	"net/http"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	"github.com/airelcamilo/podvoyage-backend/internal/user/model"
	us "github.com/airelcamilo/podvoyage-backend/internal/user/service"
	"gorm.io/gorm"
)

type UserController struct {
	service *us.UserService
}

func (c *UserController) New(db *gorm.DB) UserController {
	var service us.UserService
	service = service.New(db)
	return UserController{&service}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.RegisterRequest
	utils.ParseBody(w, r, &request)
	if response, err := c.service.Register(&request); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, response)
	}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.LoginRequest
	utils.ParseBody(w, r, &request)
	if response, err := c.service.Login(&request); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, response)
	}
}

func (c *UserController) Validate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.ValidateRequest
	utils.ParseBody(w, r, &request)
	if user, err := c.service.Validate(request.Token); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, user)
	}
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.ValidateRequest
	utils.ParseBody(w, r, &request)
	if user, err := c.service.Logout(request.Token); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, user)
	}
}
