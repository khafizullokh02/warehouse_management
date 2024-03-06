package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/khafizullokh02/warehouse_management/token"
	"github.com/khafizullokh02/warehouse_management/util"
)

type Server struct {
	config util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users",  server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/user/:id", server.getUser)
	authRoutes.PUT("/users/update", server.updateUser)
	authRoutes.GET("/users/", server.listUsers)
	authRoutes.DELETE("/user/:id", server.deleteUser)

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/account/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.DELETE("/account/:id", server.deleteAccount)

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}