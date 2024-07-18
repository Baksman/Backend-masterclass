package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	db "github.com/baksman/backend_masterclass/db/sqlc"
	"github.com/baksman/backend_masterclass/util"
)

type createUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
}

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (s *Server) getUser(ctx *gin.Context) {
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

func (s *Server) getUserRequest(ctx *gin.Context) {
	var req getUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := s.store.GetUser(ctx, req.Username)

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

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	hashedPassowrd, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassowrd,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := s.store.CreateUser(ctx, arg)

	if err != nil {

		if pqError, ok := err.(*pq.Error); ok {
			var errorName = pqError.Code.Name()
			if errorName == "unique_violation" {
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
			// log.Println(pqError.Code.Name()+"code name")
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	createUser := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
	}

	ctx.JSON(http.StatusOK, createUser)

}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,min=3"`
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
