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
}

func (c *EpisodeController) New(db *gorm.DB) EpisodeController {
	var service ps.EpisodeService
	service = service.New(db)
	return EpisodeController{&service}
}

func (c *EpisodeController) MarkAsPlayed(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	if items, err := c.service.MarkAsPlayed(id); err != nil {
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

	if podcast, err := c.service.SetCurrentTime(id, &request); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, podcast)
	}
}
