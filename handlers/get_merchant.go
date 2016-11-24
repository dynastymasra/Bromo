package handlers

import (
	"bromo/models"
	"bromo/settings"
	"bromo/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// GetMerchantStickerHandler used
func GetMerchantStickerHandler(c *gin.Context) {

	id := c.Param("id")

	db, err := settings.Connector()
	if err != nil {
		log.Errorf("get_merchant - Connector : %v", err)
		c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error()))
	}
	defer db.Close()

	packs := []models.Pack{}

	db.Where("created_by = ?", id).Find(&packs)
	for i := range packs {
		db.Model(packs[i]).Related(&packs[i].Stickers)
	}

	c.JSON(http.StatusOK, utils.ObjectResponse(packs))
}
