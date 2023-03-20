package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	db "github.com/Placebo900/simple-bank/db/sqlc"
	"github.com/Placebo900/simple-bank/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (s *Server) createAccount(c *gin.Context) {
	req := createAccountRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	acc, err := s.store.CreateAccount(c,
		db.CreateAccountParams{Owner: authPayload.Username, Balance: 0, Currency: req.Currency})
	if err != nil {
		log.Println(err.(*pq.Error).Code.Name())
		switch err.(*pq.Error).Code.Name() {
		case "foreign_key_violation", "unique_violation":
			c.JSON(http.StatusForbidden, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, acc)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(c *gin.Context) {
	arg := getAccountRequest{}
	if err := c.ShouldBindUri(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	acc, err := s.store.GetAccount(c, arg.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	if acc.Owner != authPayload.Username {
		err := errors.New("account doesn't below to the authenticated user")
		c.JSON(http.StatusUnauthorized, err)
	}
	c.JSON(http.StatusOK, acc)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) listAccount(c *gin.Context) {
	req := listAccountRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	acc, err := s.store.ListAccount(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, acc)
}
