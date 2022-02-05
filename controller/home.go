package controller

import (
	"layoutrender-boilerplate/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source home.go -destination mock_home.go -package=controller
type HomeApp interface {
	HomeRender() app.LayoutRender
}

type HomeController struct {
	Home HomeApp
}

func NewHomeController(homeApp HomeApp) *HomeController {
	return &HomeController{
		Home: homeApp,
	}
}

func (hc *HomeController) HomePage(c *gin.Context) {
	c.Render(http.StatusOK, hc.Home.HomeRender())
}
