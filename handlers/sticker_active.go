package handlers

import (
	"bromo/models"
	"bromo/settings"
	"bromo/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// GetAllStickerActiveMerchant used
func GetAllStickerActiveMerchant(c *gin.Context) {

	id := c.Param("id")

	db, err := settings.Connector()
	if err != nil {
		log.Errorf("sticker_active - Connector : %v", err)
		c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error()))
	}
	defer db.Close()

	stickerPack := []models.StickerPack{}

	db.Where("status = ? AND merchant_id = ?", true, id).Find(&stickerPack)

	packs := make([]models.Pack, 0, len(stickerPack))
	for i := range stickerPack {
		res := models.Pack{}
		db.Where("id = ?", stickerPack[i].PackID).First(&res)
		packs = append(packs, res)
	}

	for i := range packs {
		db.Model(packs[i]).Related(&packs[i].Stickers)
	}

	c.JSON(http.StatusOK, utils.ObjectResponse(packs))
}
