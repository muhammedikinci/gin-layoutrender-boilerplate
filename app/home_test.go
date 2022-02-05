package app

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Home_Render(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockHomeRepository(ctrl)

	data := gin.H{
		"h1": "Test",
	}
	m.EXPECT().GetH1().Return(data["h1"]).Times(1)

	homeApp := NewHomeApp(m)

	lr := homeApp.HomeRender()
	var lrData = lr.Data.(gin.H)

	assert.Equal(t, data["h1"], lrData["h1"])
}
