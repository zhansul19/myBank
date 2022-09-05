package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/zhansul19/myBank/config"
	db "github.com/zhansul19/myBank/db/sqlc"
	"github.com/zhansul19/myBank/token"
)

type Server struct {
	config     config.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config config.Config, db db.Store) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("couldn't create token: %w", err)
	}

	server := &Server{
		config:     config,
		store:      db,
		tokenMaker: token,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRooutes()
	return server, nil
}
func (server *Server) setupRooutes() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// router.GET("/user/:id", server.getUser)

	authRoutes:=router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
//	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	authRoutes.POST("/transfer", server.CreateTransfer)
	server.router = router
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
