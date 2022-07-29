package endpoint

import "github.com/gin-gonic/gin"

type PlatformEndPoint struct {
	router *gin.Engine
}

func NewPlatformEndPoint(router *gin.Engine) *PlatformEndPoint {
	endPoint := &PlatformEndPoint{router: router}
	return endPoint
}

func (end *PlatformEndPoint) RegisterEndPoint() {

}
