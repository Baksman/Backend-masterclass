package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/baksman/backend_masterclass/db/sqlc"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,validCurrency"`
}

func (s *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := s.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

	if !s.validateAccount(ctx, arg.ToAccountID, req.Currency) {
		return
	}

	if !s.validateAccount(ctx, arg.FromAccountID, req.Currency) {
		return
	}

}

func (s *Server) validateAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := s.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return false
		}
	}

	if account.Currency != currency {
		err = fmt.Errorf("account currency mismatched %v isnt macatched for account id %v", account.Currency, account.ID)
		errorResponse(err)
		return false
	}
	return true
}
