package router

import (
	uc "github.com/airelcamilo/podvoyage-backend/internal/user/controller"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserRouter(r *mux.Router, DB *gorm.DB) {
	var userController uc.UserController
	userController = userController.New(DB)

	r.HandleFunc("/api/register", userController.Register).Methods("POST")
	r.HandleFunc("/api/login", userController.Login).Methods("POST")
	r.HandleFunc("/api/validate", userController.Validate).Methods("POST")
	r.HandleFunc("/api/logout", userController.Logout).Methods("POST")
}