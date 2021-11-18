package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/jana-o/subscription/db/sqlc"
	"net/http"
	"time"
)

type createSubscriptionRequest struct {
	UserID    uuid.UUID       `json:"user_id"  binding:"required"`
	ProductID uuid.UUID       `json:"product_id"  binding:"required"`
	StartDate time.Time  	 `json:"start_date"  binding:"required"`
	EndDate   time.Time   	 `json:"end_date"`
	Tax       float64 		 `json:"tax"`
}

func (s *Server) createSubscription(ctx *gin.Context) {
	var req createSubscriptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// get user: id, tax
	// get product id, product.duration -> endDate=  startDate + duration,
	//startDate := time.Now()
	//endDate := startDate.AddDate(0,3,0)

	arg := db.CreateSubscriptionParams{
		UserID: 	req.UserID,
		ProductID: req.ProductID,
		//StartDate: startDate.Date(),
		//EndDate:   endDate.Date(),
		//Tax: 		req.Tax,
	}

	subscription, err := s.store.CreateSubscription(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, subscription)
}

type getSubscriptionByIDRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (s *Server) getSubscriptionByID(ctx *gin.Context) {
	var req getSubscriptionByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)
	subscription, err := s.store.GetSubscriptionByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, subscription)
}

type pauseSubscriptionRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

// pauseSubscription updates Subscription status and end date
func (s *Server) pauseSubscription(ctx *gin.Context) {
	var req pauseSubscriptionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)
	// if subscription status == active and startDate after today
	// set status = paused, paused_at= time.Now(), updated_at = time.Now()

	// unpause: if status == paused
	// assumption: pause/unpause and extend as often as you want
	// daysLeft on curr subscription:= endDate-startDate, newEndDate := endDate + daysLeft, set endDate=newEndDate,update = time.Now()
	subscription, err := s.store.PauseSubscription(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, subscription)
}

// cancelSubscription update subscription
func (s *Server) cancelSubscription(ctx *gin.Context) {
	var req pauseSubscriptionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := uuid.MustParse(req.ID)
	// if subscription.status == active and startDate after today
	// if subscription.status == paused return errors.New("cannot cancel paused subscription")
	// set status == canceled , updated = time.Now(),  deleted_at = time.Now()
	subscription, err := s.store.CancelSubscription(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, subscription)
}