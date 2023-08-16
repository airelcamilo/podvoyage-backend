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
}

func (c *FolderController) New(db *gorm.DB) FolderController {
	var service ps.FolderService
	service = service.New(db)
	return FolderController{&service}
}

func (c *FolderController) GetAllFolder(w http.ResponseWriter, r *http.Request) {
	if items, err := c.service.GetAllFolder(); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, items)
	}
}

func (c *FolderController) GetFolder(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	if podcast, err := c.service.GetFolder(id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, podcast)
	}
}

func (c *FolderController) SaveFolder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request model.Folder
	utils.ParseBody(w, r, &request)

	if podcast, err := c.service.SaveFolder(&request); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusCreated, podcast)
	}
}

func (c *FolderController) CheckInFolder(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	if podcast, err := c.service.CheckInFolder(id); err != nil {
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

	if podcast, err := c.service.ChangeFolder(folderId, podId); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusCreated, podcast)
	}
}

func (c *FolderController) RemoveFolder(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(w, r)
	if err := c.service.RemoveFolder(id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "Bad Request")
	} else {
		utils.FormatResponse(w, http.StatusOK, id)
	}
}
