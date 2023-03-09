package api

import (
	"fmt"
	"net/http"

	db "github.com/Placebo900/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getTransferRequest struct {
	ID int64 `url:"id" binding:"required,min=1"`
}

func (s *Server) getTransfer(c *gin.Context) {
	arg := getTransferRequest{}
	if err := c.BindUri(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	transfer, err := s.store.GetTransfer(c, arg.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transfer)
}

type listTransferRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) listTransfer(c *gin.Context) {
	arg := listTransferRequest{}
	if err := c.BindQuery(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	transfers, err := s.store.ListTransfer(c, db.ListTransferParams{Limit: arg.PageSize, Offset: (arg.PageID - 1) * arg.PageSize})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, transfers)
}

type transferTxRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (s *Server) transferTx(c *gin.Context) {
	arg := transferTxRequest{}
	if err := c.BindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := checkCurrencies(c, s, arg.FromAccountID, arg.ToAccountID, arg.Amount, arg.Currency); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := s.store.TransferTx(c, db.TransferTxParams(arg))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func checkCurrencies(c *gin.Context, s *Server, from, to, amount int64, currency string) error {
	accFrom, err := s.store.GetAccount(c, from)
	if err != nil {
		return err
	}
	accTo, err := s.store.GetAccount(c, to)
	if err != nil {
		return err
	}

	if accFrom.Currency != currency {
		return fmt.Errorf("different currencies from id=%d: want: %s, have: %s", from, currency, accFrom.Currency)
	}
	if accTo.Currency != currency {
		return fmt.Errorf("different currencies from id=%d: want:%s, have:%s", to, currency, accTo.Currency)
	}
	if accFrom.Balance < amount {
		return fmt.Errorf("id=%d balance is lower than amount of transfer: amount:%d, balance:%d", accFrom.ID, amount, accFrom.Balance)
	}
	return nil
}
