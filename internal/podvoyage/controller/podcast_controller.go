package podvoyagecontroller

import (
	"net/http"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	ps "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/service"
	"gorm.io/gorm"
)

type PodcastController struct {
	service *ps.PodcastService
	auth *utils.AuthUtils
}

func (c *PodcastController) New(db *gorm.DB) PodcastController {
	var service ps.PodcastService
	service = service.New(db)
	var auth utils.AuthUtils
	auth = auth.New(db)
	return PodcastController{&service, &auth}
}

func (c *PodcastController) SearchPodcasts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.SearchRequest
	utils.ParseBody(w, r, &request)
	_, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if searchResult, err := c.service.SearchPodcasts(&request); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, searchResult)
	}
}

func (c *PodcastController) SearchPodcast(w http.ResponseWriter, r *http.Request) {
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	id := utils.GetId(w, r)
	if podcast, err := c.service.SearchPodcast(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, podcast)
	}
}

func (c *PodcastController) GetAllPodcast(w http.ResponseWriter, r *http.Request) {
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if items, err := c.service.GetAllPodcast(user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, items)
	}
}

func (c *PodcastController) GetPodcast(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if podcast, err := c.service.GetPodcast(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, podcast)
	}
}

func (c *PodcastController) SavePodcast(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.Podcast
	utils.ParseBody(w, r, &request)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if podcast, err := c.service.SavePodcast(&request, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusCreated, podcast)
	}
}

func (c *PodcastController) RemovePodcast(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if err := c.service.RemovePodcast(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "Bad Request")
	} else {
		utils.FormatResponse(w, http.StatusOK, id)
	}
}