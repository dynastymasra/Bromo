package handlers

import (
	"bromo/settings"
	"bromo/utils"
	"net/http"

	"bromo/models"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// CreateStickerHandler used
func CreateStickerHandler(c *gin.Context) {

	var pack models.Pack

	c.Header("Content-Type", "application/json")
	if c.BindJSON(&pack) == nil {

		db, err := settings.Connector()
		if err != nil {
			log.Errorf("create - Connector : %v", err)
			c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error()))
		}
		defer db.Close()

		packPassed := models.CheckDataPack(pack)

		stickerPack := models.StickerPack{}
		stickerPack.MerchantID = pack.CreatedBy
		stickerPack.Status = true
		stickerPack.PackID = packPassed.ID

		db.Save(packPassed)
		db.Save(stickerPack)

		c.JSON(http.StatusCreated, utils.ObjectResponse(packPassed))
	} else {
		log.Errorf("create - BindJSON : %v", c.BindJSON(&pack))
		c.JSON(http.StatusBadRequest, utils.FailResponse("Data request not valid"))
	}
}
