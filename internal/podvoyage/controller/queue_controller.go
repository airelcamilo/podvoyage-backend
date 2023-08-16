package podvoyagecontroller

import (
	"net/http"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	ps "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/service"
	"gorm.io/gorm"
)

type QueueController struct {
	service *ps.QueueService
}

func (c *QueueController) New(db *gorm.DB) QueueController {
	var service ps.QueueService
	service = service.New(db)
	return QueueController{&service}
}

func (c *QueueController) GetAllQueue(w http.ResponseWriter, r *http.Request) {
	if queue, err := c.service.GetAllQueue(); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, queue)
	}
}

func (c *QueueController) AddToQueue(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.Episode
	utils.ParseBody(w, r, &request)

	if episode, err := c.service.AddToQueue(&request); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, episode)
	}
}

func (c *QueueController) RemoveInQueue(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	if err := c.service.RemoveInQueue(id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "Bad Request")
	} else {
		utils.FormatResponse(w, http.StatusOK, id)
	}
}
