package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/lib/pq"
)

type createEntryItemRequest struct {
	ProductID    int32  `json:"product_id" binding:"required"`
	EntryGroupID int32  `json:"entry_group_id" binding:"required"`
	SupCode      string `json:"sup_code" binding:"required"`
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

type getEntryItemsRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEntryItem(ctx *gin.Context) {
	var req getEntryItemsRequest
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

type listEntryItemsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listEntryItems(ctx *gin.Context) {
	var req listEntryItemsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEntryItemsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	entryItems, err := server.store.ListEntryItems(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entryItems)
}

type updateEntryItemsRequest struct {
	ProductID    int32 `json:"product_id" binding:"required"`
	EntryGroupID int32 `json:"entry_group_id" binding:"required"`
	SupCode      int32 `json:"sup_code" bindnig:"required"`
}

func (server *Server) updateEntryItem(ctx *gin.Context) {
	var req updateEntryItemsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var updateReq db.UpdateEntryItemParams
	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := json.Unmarshal(bytes, &updateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entryItems, err := server.store.UpdateEntryItem(ctx, updateReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entryItems)
}

type deleteEntryItemRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
}

func (server *Server) deleteEntryItem(ctx *gin.Context) {
	var req deleteEntryItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteEntryItem(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}