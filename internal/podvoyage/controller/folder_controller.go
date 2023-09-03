package podvoyagecontroller

import (
	"net/http"
	"strconv"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	"github.com/airelcamilo/podvoyage-backend/internal/podvoyage/model"
	ps "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type FolderController struct {
	service *ps.FolderService
	auth    *utils.AuthUtils
}

func (c *FolderController) New(db *gorm.DB) FolderController {
	var service ps.FolderService
	service = service.New(db)
	var auth utils.AuthUtils
	auth = auth.New(db)
	return FolderController{&service, &auth}
}

func (c *FolderController) GetAllFolder(w http.ResponseWriter, r *http.Request) {
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if items, err := c.service.GetAllFolder(user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, items)
	}
}

func (c *FolderController) GetFolder(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if podcast, err := c.service.GetFolder(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, podcast)
	}
}

func (c *FolderController) SaveFolder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.Folder
	utils.ParseBody(w, r, &request)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if podcast, err := c.service.SaveFolder(&request, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusCreated, podcast)
	}
}

func (c *FolderController) CheckInFolder(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if podcast, err := c.service.CheckInFolder(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, podcast)
	}
}

func (c *FolderController) ChangeFolder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folderId, err := strconv.Atoi(vars["folderId"])
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, nil)
		return
	}
	podId, err := strconv.Atoi(vars["podId"])
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, nil)
		return
	}
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if podcast, err := c.service.ChangeFolder(folderId, podId, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusCreated, podcast)
	}
}

func (c *FolderController) RemoveFolder(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if err := c.service.RemoveFolder(id, user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "Bad Request")
	} else {
		utils.FormatResponse(w, http.StatusOK, id)
	}
}
