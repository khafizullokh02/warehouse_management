package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Name string `json:"name" binding:"required,name"`
	SupCode string `json:"sup_code" binding:"required,sup_code"`
	BarCode string `json:"bar_code" binding:"required,bar_code"`
	Image string `json:"image" binding:"required,image"`
	Brand string `json:"brand" binding:"required,brand"`
	WholesalePrice float64 `json:"wholesale_price" binding:"required,wholesale_price"`
	RetailPrice float64 `json:"retail_price" binding:"required,retail_price"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	
}