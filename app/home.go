package app

import "github.com/gin-gonic/gin"

//go:generate mockgen -source home.go -destination mock_home.go -package=app
type HomeRepository interface {
	GetH1() string
}

type Home struct {
	Repository HomeRepository
}

func NewHomeApp(repository HomeRepository) *Home {
	return &Home{
		Repository: repository,
	}
}

func (ha *Home) HomeRender() LayoutRender {
	lr := NewLayoutRender("public/templates", "default")
	lr.Page = "index"

	lr.Data = gin.H{
		"title": "Home Page Title",
		"h1":    ha.Repository.GetH1(),
	}

	return lr
}
