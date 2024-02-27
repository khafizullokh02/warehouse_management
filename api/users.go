package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required,name"`
	Role     string `json:"role" binding:"required,role"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

type userResponse struct {
	Name              string    `json:"name"`
	Role              string    `json:"role"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Name:              user.Name,
		Role:              user.Role,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:     req.Name,
		Role:     req.Role,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}
