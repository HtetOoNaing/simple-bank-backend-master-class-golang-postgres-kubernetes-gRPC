package api

import (
	db "github.com/HtetOoNaing/simple-bank-backend-master-class-golang-postgres-kubernetes-gRPC/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run()
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
