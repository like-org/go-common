package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/like-org/common/app/service/account/client"
)

type Server struct {
	router *gin.Engine
	accSrv *client.AccountClient
}

func NewServer() *Server {
	server := &Server{}
	server.setupRouter()
	return server
}

func (srv *Server) setupRouter() {
	router := gin.Default()

	router.GET("/ping", srv.Ping)

	router.POST("/members", srv.CreateMember)
	router.GET("/verify_code", srv.GetVerfiyCode)

	srv.router = router
}
func (srv *Server) Start(address string) error {
	return srv.router.Run(address)
}
func (src *Server) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
