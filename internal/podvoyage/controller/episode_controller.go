package podvoyagecontroller

import (
	"net/http"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	ps "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/service"
	"gorm.io/gorm"
)

type EpisodeController struct {
	service *ps.EpisodeService
	auth *utils.AuthUtils
}

func (c *EpisodeController) New(db *gorm.DB) EpisodeController {
	var service ps.EpisodeService
	service = service.New(db)
	var auth utils.AuthUtils
	auth = auth.New(db)
	return EpisodeController{&service, &auth}
}

func (c *EpisodeController) MarkAsPlayed(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if items, err := c.service.MarkAsPlayed(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, items)
	}
}

func (c *EpisodeController) SetCurrentTime(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	defer r.Body.Close()
	var request model.CurrentTimeRequest
	utils.ParseBody(w, r, &request)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if podcast, err := c.service.SetCurrentTime(id, &request, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, podcast)
	}
}
