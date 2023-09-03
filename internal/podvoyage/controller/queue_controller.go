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
	auth *utils.AuthUtils
}

func (c *QueueController) New(db *gorm.DB) QueueController {
	var service ps.QueueService
	service = service.New(db)
	var auth utils.AuthUtils
	auth = auth.New(db)
	return QueueController{&service, &auth}
}

func (c *QueueController) GetAllQueue(w http.ResponseWriter, r *http.Request) {
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}
	if queue, err := c.service.GetAllQueue(user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, queue)
	}
}

func (c *QueueController) AddToQueue(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.Episode
	utils.ParseBody(w, r, &request)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if episode, err := c.service.AddToQueue(&request, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, episode)
	}
}

func (c *QueueController) RemoveInQueue(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if err := c.service.RemoveInQueue(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "Bad Request")
	} else {
		utils.FormatResponse(w, http.StatusOK, id)
	}
}
