package handlers

import (
	"bromo/models"
	"bromo/settings"
	"bromo/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// DeletePackByIDHandler used
func DeletePackByIDHandler(c *gin.Context) {

	id := c.Param("id")

	db, err := settings.Connector()
	if err != nil {
		log.Errorf("delete_pack - Connector : %v", err)
		c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error()))
	}
	defer db.Close()

	db.Unscoped().Where("pack_id = ?", id).Delete(&models.Sticker{})
	db.Unscoped().Where("pack_id = ?", id).Delete(&models.StickerPack{})
	db.Unscoped().Where("id = ?", id).Delete(&models.Pack{})

	c.JSON(http.StatusOK, utils.SuccessResponse())
}
