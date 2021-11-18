package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jana-o/subscription/config"
	"github.com/jana-o/subscription/db/sqlc"
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

func (s *Server) setupRouter() {
	router := gin.Default()

	router.GET("/products", s.getProducts)
	router.GET("/products/:id", s.getProductByID)
	router.POST("/products/:id/buy", s.createSubscription)
	router.GET("/subscription/:id", s.getSubscriptionByID)
	router.PATCH("/subscription/:id", s.pauseSubscription)
	// router.PATCH("/subscription/:id/cancel", s.cancelSubscription)

	s.router = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
