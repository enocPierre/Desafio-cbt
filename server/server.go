package server

import (
	"github.com/estudo-go/Desafio-cbt/server/routes"
	"github.com/gin-gonic/gin"

	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

// construtor de server
func NewServer() Server {
	return Server{
		port:   "5001",
		server: gin.Default(),
	}
}

// metodo que vai roda nosso server
func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	//	router := routes.ConfigRoutes(s.server)

	log.Print("server is runing at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
