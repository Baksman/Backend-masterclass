package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	db "github.com/baksman/backend_masterclass/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := s.store.ListAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)

}

func (s *Server) getAccountRequest(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := s.store.GetAccount(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	}

	ctx.JSON(http.StatusOK, account)
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := s.store.CreateAccount(ctx, arg)

	if err != nil {

		if pqError, ok := err.(*pq.Error); ok {
			var errorName = pqError.Code.Name()
			if errorName == "foreign_key_violation" || errorName == "unique_violation" {
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
			// log.Println(pqError.Code.Name()+"code name")
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
