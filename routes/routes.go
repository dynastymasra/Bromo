package routes

import (
	"bromo/handlers"

	"github.com/gin-gonic/gin"
)

// EndpointHandler used
func EndpointHandler(c *gin.RouterGroup) {
	c.POST("/stickers/pack", handlers.CreateStickerHandler)

	c.GET("/stickers", handlers.GetAllStickerPublic)
	c.GET("/stickers/pack/:id", handlers.GetPackByIDHandler)
	c.GET("/stickers/merchant/created/:id", handlers.GetMerchantStickerHandler)
	c.GET("/stickers/merchant/active/:id", handlers.GetAllStickerActiveMerchant)

	c.PUT("/stickers/merchant/active/:id", handlers.UpdateStatusHandler)

	c.DELETE("/stickers/pack/:id", handlers.DeletePackByIDHandler)
}
