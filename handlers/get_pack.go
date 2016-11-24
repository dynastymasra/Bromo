package handlers

import (
	"bromo/models"
	"bromo/settings"
	"bromo/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// GetPackByIDHandler used
func GetPackByIDHandler(c *gin.Context) {

	id := c.Param("id")

	db, err := settings.Connector()
	if err != nil {
		log.Errorf("get_pack - Connector : %v", err)
		c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error()))
	}
	defer db.Close()

	pack := models.Pack{}

	db.Where("id = ?", id).First(&pack)
	db.Model(&pack).Related(&pack.Stickers)

	c.JSON(http.StatusOK, utils.ObjectResponse(pack))
}
