package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "github.com/baksman/backend_masterclass/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	// v := validator.New()

	// Register the custom email validation function.
	// v.RegisterValidation("validCurrency", validateCurrency)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validCurrency", validateCurrency)
	}
	router.POST("/accounts", server.createAccount)
	router.POST("/users", server.createUser)
	router.GET("/accounts/:id", server.getAccountRequest)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// staert http server on the input address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
