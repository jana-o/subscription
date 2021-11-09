package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jana-o/subscription/config"
	"github.com/jana-o/subscription/db"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     config.Config
	store      db.Store
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config config.Config, store db.Store) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/products", server.getProducts)
	router.GET("/products/:id", server.getProductByID)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
