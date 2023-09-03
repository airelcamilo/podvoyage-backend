package podvoyagecontroller

import (
	"net/http"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	ps "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/service"
	"gorm.io/gorm"
)

type ItemController struct {
	service *ps.ItemService
	auth *utils.AuthUtils
}

func (c *ItemController) New(db *gorm.DB) ItemController {
	var service ps.ItemService
	service = service.New(db)
	var auth utils.AuthUtils
	auth = auth.New(db)
	return ItemController{&service, &auth}
}

func (c *ItemController) GetAllItem(w http.ResponseWriter, r *http.Request) {
	user, err := c.auth.GetUser(w, r)
	if err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, "error while reading cookie")
		return
	}

	if items, err := c.service.GetAllItem(user.Id); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, items)
	}
}

