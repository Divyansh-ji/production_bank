package api

import (
	"fmt"

	db "github.com/Divyansh-ji/production_bank/db/sqlc"
	"github.com/Divyansh-ji/production_bank/token"
	"github.com/Divyansh-ji/production_bank/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer create a new HTTP server and setup routing
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.GET("/accounts/all", server.listAllAccounts)
	router.DELETE("/account/:id", server.deleteAccount)
	router.POST("/transfer", server.createTransfer)
	server.router = router
	return server

}

// start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
