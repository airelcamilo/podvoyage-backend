package router

import (
	"github.com/airelcamilo/podvoyage-backend/internal/pkg/db"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	DB := db.Connect()
	r := mux.NewRouter()

	PodvoyageRouter(r, DB)
	UserRouter(r, DB)

	return r
}
