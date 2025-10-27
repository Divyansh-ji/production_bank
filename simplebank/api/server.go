package api

import (
	db "github.com/Divyansh-ji/production_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
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
