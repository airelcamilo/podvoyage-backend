package podvoyagecontroller

import (
	"net/http"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/utils"
	ps "github.com/airelcamilo/podvoyage-backend/internal/podvoyage/service"
	"gorm.io/gorm"
)

type ItemController struct {
	service *ps.ItemService
}

func (c *ItemController) New(db *gorm.DB) ItemController {
	var service ps.ItemService
	service = service.New(db)
	return ItemController{&service}
}

func (c *ItemController) GetAllItem(w http.ResponseWriter, r *http.Request) {
	if items, err := c.service.GetAllItem(); err != nil {
		utils.FormatResponse(w, http.StatusBadRequest, err)
	} else {
		utils.FormatResponse(w, http.StatusOK, items)
	}
}

