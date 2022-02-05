package controller

import (
	"layoutrender-boilerplate/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Home_Page(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockHomeApp(ctrl)

	lr := app.NewLayoutRender("../public/templates", "default")
	lr.Page = "index"

	m.EXPECT().HomeRender().Return(lr).Times(1)

	homeController := NewHomeController(m)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	homeController.HomePage(ctx)

	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
}
