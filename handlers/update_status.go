package handlers

import (
	"bromo/settings"
	"bromo/utils"
	"net/http"

	"bromo/models"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// UpdateStatusHandler used
func UpdateStatusHandler(c *gin.Context) {

	var pack models.StickerPack

	id := c.Param("id")

	c.Header("Content-Type", "application/json")
	if c.BindJSON(&pack) == nil {

		db, err := settings.Connector()
		if err != nil {
			log.Errorf("create - Connector : %v", err)
			c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error()))
		}
		defer db.Close()
		db.LogMode(true)

		pack.MerchantID = id

		if db.Where("merchant_id = ? AND pack_id = ?", id, pack.PackID).First(&models.StickerPack{}).RecordNotFound() {
			db.Save(pack)
		} else {
			db.Model(&pack).Where("merchant_id = ? AND pack_id = ?", id, pack.PackID).Update("status", pack.Status)
		}

		c.JSON(http.StatusOK, utils.SuccessResponse())
	} else {
		log.Errorf("create - BindJSON : %v", c.BindJSON(&pack))
		c.JSON(http.StatusBadRequest, utils.FailResponse("Data request not valid"))
	}
}
