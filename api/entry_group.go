package api

import (
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
