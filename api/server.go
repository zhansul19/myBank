package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhansul19/myBank/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(db *db.Store) *Server {
	server := &Server{
		store:  db,
		router: gin.Default(),
	}
	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccount)
	server.router.DELETE("/accounts/:id", server.deleteAccount)

	return server
} 

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
