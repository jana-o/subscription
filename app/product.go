package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/jana-o/subscription/db/sqlc"
	"net/http"
)

type createProductRequest struct {
	Name        string      `json:"name" binding:"required"`
	Duration    int32   	 `json:"duration" binding:"required"`
	Price       float64 	`json:"price" binding:"required"`
	Description string  	`json:"description" binding:"required"`
}

func (s *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
	Name:        req.Name,
	Duration:    req.Duration,
	Price:       req.Price,
	Description: req.Description,
}
	product, err := s.store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type getProductByIDRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (s *Server) getProductByID(ctx *gin.Context) {
	var req getProductByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)
	account, err := s.store.GetProductByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}


func (s *Server) getProducts(ctx *gin.Context) {

	accounts, err := s.store.GetProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
