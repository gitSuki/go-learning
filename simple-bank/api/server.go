package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/gitsuki/simplebank/db/sqlc"
	"github.com/gitsuki/simplebank/token"
	"github.com/gitsuki/simplebank/util"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccounts)
	router.PUT("/account/", server.updateAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server, nil
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
