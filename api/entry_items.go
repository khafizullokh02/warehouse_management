package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/lib/pq"
)

type createEntryItemRequest struct {
	ProductID    int32  `json:"product_id"`
	EntryGroupID int32  `json:"entry_group_id"`
	SupCode      string `json:"sup_code"`
}

func (server *Server) createEntryItem(ctx *gin.Context) {
	var req createEntryItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEntryItemParams{
		ProductID:    req.ProductID,
		EntryGroupID: req.EntryGroupID,
		SupCode:      req.SupCode,
	}

	entryItem, err := server.store.CreateEntryItem(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entryItem)
}

type getEntryItemRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEntryItem(ctx *gin.Context) {
	var req getEntryItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entryItem, err := server.store.GetEntryItem(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entryItem)
}
