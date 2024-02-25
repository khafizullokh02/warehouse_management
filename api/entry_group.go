package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/lib/pq"
)

type ActionType string
type PricingType string
type EntryGroupStatus string
type AgreementFormsStatus string
type CurrencyCode string

type createEntryGroupRequest struct {
	Quantity         int32            `json:"quantity" binding:"required"`
	ActionType       ActionType       `json:"action_type" binding:"required"`
	PricingType      PricingType      `json:"pricing_type" binding:"required"`
	Price            float64          `json:"price" binding:"required"`
	Currency         CurrencyCode     `json:"currency" binding:"required"`
	EntryGroupStatus EntryGroupStatus `json:"entry_group_status" binding:"required"`
}

func (server *Server) CreateEntryGroup(ctx *gin.Context) {
	var req createEntryGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEntryGroupParams{
		Quantity:         req.Quantity,
		ActionType:       db.ActionType(req.ActionType),
		PricingType:      db.PricingType(req.PricingType),
		Price:            req.Price,
		Currency:         db.CurrencyCode(req.Currency),
		EntryGroupStatus: db.EntryGroupStatus(req.EntryGroupStatus),
	}

	entryGroup, err := server.store.CreateEntryGroup(ctx, arg)
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

	ctx.JSON(http.StatusOK, entryGroup)
}

type listEntryGroupReq struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListEntryGroups(ctx *gin.Context) {
	var req listEntryGroupReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		return
	}

	arg := db.ListEntryGroupsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	dbResp, err := server.store.ListEntryGroups(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dbResp)
}

type getEntryGroupRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type EntryGroupWithItems struct {
	db.EntryGroup
	EntryItems []db.GetEntryItemsByEntryGroupIdRow `json:"entry_items"`
}

func (server *Server) GetEntryGroup(ctx *gin.Context) {
	var req getEntryGroupRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entryGroup, err := server.store.GetEntryGroup(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	entryItems, err := server.store.GetEntryItemsByEntryGroupId(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	entryGroupWithItems := EntryGroupWithItems{
		EntryGroup: entryGroup,
		EntryItems: entryItems,
	}

	ctx.JSON(http.StatusOK, entryGroupWithItems)
}
