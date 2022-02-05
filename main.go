package main

import (
	"layoutrender-boilerplate/app"
	"layoutrender-boilerplate/config"
	"layoutrender-boilerplate/controller"
	"layoutrender-boilerplate/repository"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENV")

	if env == "" {
		panic("ENV cannot be empty")
	}

	conf := config.NewConfig(".conf", env)

	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	homeRepository := repository.Home{}
	homeApp := app.NewHomeApp(homeRepository)
	homeController := controller.NewHomeController(homeApp)

	router := gin.Default()

	router.LoadHTMLGlob("public/templates/**/*")

	router.GET("/", homeController.HomePage)

	router.GET("/static/*path", func(c *gin.Context) {
		p := c.Param("path")
		file, err := Asset("public/assets" + p)

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		if strings.Contains(p, ".css") {
			c.Writer.Header().Add("Content-Type", "text/css")
		} else if strings.Contains(p, ".js") {
			c.Writer.Header().Add("Content-Type", "text/javascript")
		} else if strings.Contains(p, ".png") {
			c.Writer.Header().Add("Content-Type", "image/png")
		} else if strings.Contains(p, ".png") {
			c.Writer.Header().Add("Content-Type", "image/png")
		} else if strings.Contains(p, ".png") {
			c.Writer.Header().Add("Content-Type", "image/jpeg")
		}

		c.Writer.Write(file)
	})

	router.Run()
}
