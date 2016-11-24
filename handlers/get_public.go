package handlers

import (
	"bromo/models"
	"bromo/settings"
	"bromo/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// GetAllStickerPublic used
func GetAllStickerPublic(c *gin.Context) {

	db, err := settings.Connector()
	if err != nil {
		log.Errorf("get_public - Connector : %v", err)
		c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error()))
	}
	defer db.Close()

	packs := []models.Pack{}

	db.Where("is_public = ?", true).Find(&packs)
	for i := range packs {
		db.Model(packs[i]).Related(&packs[i].Stickers)
	}

	c.JSON(http.StatusOK, utils.ObjectResponse(packs))
}
