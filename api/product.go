package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/lib/pq"
)

type createProductRequest struct {
	Name           string  `json:"name" binding:"required"`
	SupCode        string  `json:"sup_code" binding:"required"`
	BarCode        string  `json:"bar_code" binding:"required"`
	Image          string  `json:"image" binding:"required"`
	Brand          string  `json:"brand" binding:"required"`
	WholesalePrice float64 `json:"wholesale_price" binding:"required"`
	RetailPrice    float64 `json:"retail_price" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		Name:           req.Name,
		SupCode:        req.SupCode,
		BarCode:        req.BarCode,
		Image:          req.Image,
		Brand:          req.Brand,
		WholesalePrice: req.WholesalePrice,
		RetailPrice:    req.RetailPrice,
	}

	product, err := server.store.CreateProduct(ctx, arg)
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

	ctx.JSON(http.StatusOK, product)
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx, int32(req.ID))
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type listProductRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listProduct(ctx *gin.Context) {
	var req listProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		return
	}

	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	product, err := server.store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type updateProductRequest struct {
	Name           string  `json:"name" binding:"required,alphanum"`
	SupCode        string  `json:"sup_code" binding:"required,gte=5"`
	BarCode        string  `json:"bar_code" binding:"required,gte=5"`
	Image          string  `json:"image" binding:"gte=0"`
	Brand          string  `json:"brand" binding:"required,min=4"`
	WholesalePrice float64 `json:"wholesale_price" binding:"required,gt=0"`
	RetailPrice    float64 `json:"retail_price" binding:"required,gt=0"`
	Discount       float64 `json:"discount" binding:"required,gte=0"`
	ID             int64   `uri:"id" binding:"required,min=1"`
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var req updateProductRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	arg := db.UpdateProductParams{
		Name:           req.Name,
		SupCode:        req.SupCode,
		BarCode:        req.BarCode,
		Image:          req.Image,
		Brand:          req.Brand,
		WholesalePrice: req.WholesalePrice,
		RetailPrice:    req.RetailPrice,
		Discount:       req.Discount,
		ID:             int32(req.ID),
	}

	product, err := server.store.UpdateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type deleteProductRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
