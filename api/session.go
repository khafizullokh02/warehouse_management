package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/lib/pq"
)

type CreateSessionRequest struct {
	UserAgent string `json:"user_agent"`
	ClientIp  string `json:"client_ip"`
	IsBlocked bool   `json:"is_blocked"`
	UserID    int32  `json:"user_id"`
}

func (server *Server) createSession(ctx *gin.Context) {
	var req CreateSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSessionParams{
		UserAgent: req.UserAgent,
		ClientIp:  req.ClientIp,
		IsBlocked: req.IsBlocked,
		UserID:    req.UserID,
	}

	session, err := server.store.CreateSession(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, session)
}

type getSessionRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSession(ctx *gin.Context) {
	var req getSessionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	session, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, session)
}
