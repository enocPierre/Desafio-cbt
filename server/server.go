package server

import (
	"github.com/gin-gonic/gin"
	"github.com/learming/webapi-mvc-golang/server/routes"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

// construtor
func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

//instacia o server
func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("Server is runing at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
