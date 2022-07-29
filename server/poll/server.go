package poll

import (
	"fmt"
	"github.com/dukryung/poll-server/server/poll/endpoint"
	"github.com/dukryung/poll-server/server/types"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	endpointManager *types.EndPointManager
}

func NewServer() *Server {
	s := &Server{
		router: gin.New(),
	}

	platformEndPoint := endpoint.NewPlatformEndPoint(s.router)
	s.endpointManager = types.NewEndPointManager(platformEndPoint)
	s.endpointManager.RegisterEndPoint()

	return s
}

func (s *Server) Run() {
	err := s.router.Run(":13578")
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) Close() {

}
