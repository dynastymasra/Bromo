package main

import (
	"flag"
	"time"

	"bromo/models"
	"bromo/settings"
	"bromo/utils"
	"net/http"

	"bromo/routes"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

var (
	address = flag.String("addr", ":8080", "Address to run web")
	mode    = flag.String("mode", "release", "Mode to run http server api")
	logPath = flag.String("log-path", "access.log", "Log file to redirect all logrus log messages")
	env     = flag.String("env-file", "resources/environment", "Dotenv filepath")
)

func init() {
	flag.Parse()
	gin.SetMode(*mode)

	err := godotenv.Load(*env)
	if err != nil {
		log.Fatalf("main - Load : %v", err)
	}

	db, err := settings.Connector()
	if err != nil {
		log.Fatalf("main - Connector : %v", err)
	}
	defer db.Close()
	db.AutoMigrate(models.Pack{}, models.Sticker{}, models.StickerPack{})

	log.Info("Initialize engine...")
}

func main() {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	router.NoRoute(func(c *gin.Context) {
		log.Warnf("main - NoRoute : endpoint not found %v", c.Request.URL.Path)
		c.JSON(http.StatusNotFound, utils.FailResponse("URL not found check again"))
	})

	v1 := router.Group("/v1")
	{
		routes.EndpointHandler(v1)
	}

	router.Run(*address)
}
